package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/kirktriplefive/labsMed"
)

type Patient interface {
	CreatePatient(patient labsMed.Patient) (int, error)
}

type RepositoryPatient struct {
	Patient
}

func NewRepository(db *sqlx.DB) *RepositoryPatient {
	return &RepositoryPatient{
		Patient: NewPatientPostgres(db),
	}
}
