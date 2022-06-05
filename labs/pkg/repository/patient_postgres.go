package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/kirktriplefive/labsMed"
)

type PatientPostgres struct {
	db *sqlx.DB
}

func NewPatientPostgres(db *sqlx.DB) *PatientPostgres {
	return &PatientPostgres{db: db}
}

func (r *PatientPostgres) CreatePatient(patient labsMed.Patient) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var patientId int
	createPatientQuery := fmt.Sprintf("INSERT INTO %s (name, second_name) VALUES ($1, $2) RETURNING p_id", pacientTable)
	row := tx.QueryRow(createPatientQuery, patient.Name, patient.SecondName)
	if err := row.Scan(&patientId); err != nil {
		tx.Rollback()
		return 0, err
	}
	return patientId, tx.Commit()

}

func (r *PatientPostgres) CreatePatientRecord(record labsMed.Record) (int, error) {
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

func (r *PatientPostgres) GetRecordOfPatient(id int) ([]labsMed.Record, error) {
	var record []labsMed.Record
	recordQuery := fmt.Sprintf("SELECT td.r_id, td.date, td.doctor_id, td.pacient_id, td.diagnosis FROM %s td INNER JOIN %s pt ON pt.p_id=td.pacient_id WHERE pt.p_id = $1",
		recordsTable, pacientTable)
	if err := r.db.Select(&record, recordQuery, id); err != nil {
		return record, err
	}

	return record, nil
}
