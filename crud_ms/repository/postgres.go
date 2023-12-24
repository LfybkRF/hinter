package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host string
	Port string
	Username string
	Password string
	DBname string
	SSLmode string
}

func NewPostgresDB(config Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.Username, config.Password, config.DBname, config.SSLmode))
	if err != nil {
        return nil, err
    }

	err = db.Ping()
	
	if err != nil {
        return nil, err
    }

	return db, nil
}