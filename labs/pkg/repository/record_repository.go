package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/kirktriplefive/labsMed"
)

type Record interface {
	CreatePatientRecord(record labsMed.Record) (int, error)
	GetRecords() ([]labsMed.Record, error)
	GetRecordOfPatient(int) ([]labsMed.Record, error)
}

type RepositoryRecord struct {
	Record
}

func NewRecordRepository(db *sqlx.DB) *RepositoryRecord {
	return &RepositoryRecord{
		Record: NewRecordPostgres(db),
	}
}
