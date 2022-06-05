package service

import (
	"github.com/kirktriplefive/labsMed"
	"github.com/kirktriplefive/labsMed/pkg/repository"
)

type PatientService struct {
	repo   repository.Patient
	repo_d repository.Doctor
	repo_r repository.Record
}

func NewOrderService(repo repository.Patient, repo_d repository.Doctor, repo_r repository.Record) *PatientService {
	return &PatientService{
		repo:   repo,
		repo_d: repo_d,
		repo_r: repo_r,
	}
}

func (s *PatientService) CreatePatient(patient labsMed.Patient) (int, error) {
	e, err := s.repo.CreatePatient(patient)
	return e, err
}

func (s *PatientService) CreateDoctor(doctor labsMed.Doctor) (int, error) {
	e, err := s.repo_d.CreateDoctor(doctor)
	return e, err
}

func (s *PatientService) CreatePatientRecord(record labsMed.Record) (int, error) {
	e, err := s.repo_r.CreatePatientRecord(record)
	return e, err
}

func (s *PatientService) GetDoctors() ([]labsMed.Doctor, error) {
	e, err := s.repo_d.GetDoctors()
	return e, err
}

func (s *PatientService) GetRecords() ([]labsMed.Record, error) {
	e, err := s.repo_r.GetRecords()
	return e, err
}

func (s *PatientService) GetRecordOfPatient(id int) ([]labsMed.Record, error) {
	e, err := s.repo_r.GetRecordOfPatient(id)
	return e, err
}
