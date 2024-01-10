package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"db_name"`
}

func DatabaseConnect(config DatabaseConfig) (*gorm.DB, error) {

	dsn := fmt.Sprintf("postgres://%s:%s@postgres/%s?sslmode=disable",
		config.User, config.Password, config.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return db, nil
}