package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/kirktriplefive/labsMed"
)

type Doctor interface {
	CreateDoctor(doctor labsMed.Doctor) (int, error)
	GetDoctors() ([]labsMed.Doctor, error)
}

type RepositoryDoctor struct {
	Doctor
}

func NewDoctorRepository(db *sqlx.DB) *RepositoryDoctor {
	return &RepositoryDoctor{
		Doctor: NewDoctorPostgres(db),
	}
}
