package service

import (
	"github.com/kirktriplefive/labsMed"
	"github.com/kirktriplefive/labsMed/pkg/repository"
)

type Patient interface {
	CreatePatient(patient labsMed.Patient) (int, error)
	CreatePatientRecord(record labsMed.Record) (int, error)
	CreateDoctor(doctor labsMed.Doctor) (int, error)
	GetDoctors() ([]labsMed.Doctor, error)
	GetRecords() ([]labsMed.Record, error)
	GetRecordOfPatient(int) ([]labsMed.Record, error)
}

type ServicePatient struct {
	Patient
}

func NewServicePatient(repos *repository.RepositoryPatient, repos_d *repository.RepositoryDoctor, repos_r *repository.RepositoryRecord) *ServicePatient {
	return &ServicePatient{
		Patient: &PatientService{*repos, *repos_d, *repos_r},
	}
}
