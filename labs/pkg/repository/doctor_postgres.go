package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/kirktriplefive/labsMed"
)

type DoctorPostgres struct {
	db *sqlx.DB
}

func NewDoctorPostgres(db *sqlx.DB) *DoctorPostgres {
	return &DoctorPostgres{db: db}
}

func (r *DoctorPostgres) CreateDoctor(doctor labsMed.Doctor) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var doctorId int
	createPatientQuery := fmt.Sprintf("INSERT INTO %s (name, second_name, middle_name, specialization) VALUES ($1, $2, $3, $4) RETURNING d_id", doctorsTable)
	row := tx.QueryRow(createPatientQuery, doctor.Name, doctor.SecondName, doctor.MiddleName, doctor.Specialization)
	if err := row.Scan(&doctorId); err != nil {
		tx.Rollback()
		return 0, err
	}
	return doctorId, tx.Commit()
}

func (r *DoctorPostgres) GetDoctors() ([]labsMed.Doctor, error) {
	var doctor []labsMed.Doctor
	doctorQuery := fmt.Sprintf("SELECT d_id, name, second_name, middle_name, specialization FROM %s",
		doctorsTable)
	if err := r.db.Select(&doctor, doctorQuery); err != nil {
		return doctor, err
	}

	return doctor, nil
}
