package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	pacientTable = "pacient"
	doctorsTable = "doctor"
	recordsTable = "record"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	//fmt.Println(cfg.SSLMode)
	db, err := sqlx.Open("mysql", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return db, nil
}
