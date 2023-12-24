package main

import (
	"crud_l0_ms/repository"
	"fmt"
)

func main() {
	

	db, err := repository.NewPostgresDB(repository.Config {
		Host: "localhost",
        Port: "5432",
        Username: "postgres",
        Password: "170305",
        DBname: "wb_tech_l0",
        SSLmode: "disable",
	})
	if err!= nil {
        fmt.Println("Failed to initialize database")
	}

	if db.Ping() != nil {
	    fmt.Println("Connected to database")
	}

	// repos := repository.NewRepository(db)
	

}