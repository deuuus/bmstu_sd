package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/kirktriplefive/labsMed"
)

type RecordPostgres struct {
	db *sqlx.DB
}

func NewRecordPostgres(db *sqlx.DB) *RecordPostgres {
	return &RecordPostgres{db: db}
}

func (r *RecordPostgres) CreatePatientRecord(record labsMed.Record) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var recordId int
	createPatientQuery := fmt.Sprintf("INSERT INTO %s (date, pacient_id, doctor_id, diagnosis) VALUES ($1, $2, $3, $4) RETURNING r_id", recordsTable)
	row := tx.QueryRow(createPatientQuery, record.Date, record.PacientId, record.DoctorId, record.Diagnosis)
	if err := row.Scan(&recordId); err != nil {
		tx.Rollback()
		return 0, err
	}
	return recordId, tx.Commit()
}

func (r *RecordPostgres) GetRecords() ([]labsMed.Record, error) {
	var record []labsMed.Record
	recordQuery := fmt.Sprintf("SELECT r_id, date, doctor_id, pacient_id, diagnosis FROM %s",
		recordsTable)
	if err := r.db.Select(&record, recordQuery); err != nil {
		return record, err
	}

	return record, nil
}

func (r *RecordPostgres) GetRecordOfPatient(id int) ([]labsMed.Record, error) {
	var record []labsMed.Record
	recordQuery := fmt.Sprintf("SELECT td.r_id, td.date, td.doctor_id, td.pacient_id, td.diagnosis FROM %s td INNER JOIN %s pt ON pt.p_id=td.pacient_id WHERE pt.p_id = $1",
		recordsTable, pacientTable)
	if err := r.db.Select(&record, recordQuery, id); err != nil {
		return record, err
	}

	return record, nil
}
