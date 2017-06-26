package main

import (
	"database/sql"
	"log"

	_ "github.com/amsokol/go-ignite-client/sql/http"
)

func main() {
	db, err := sql.Open("ignite-sql-http",
		`{"servers" : ["http://localhost:8081/ignite", "http://localhost:8080/ignite"], "username" : "login", "password" : "password", "cache" : "Person"}`)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare(`insert into "Organization".Organization(_key, name) values(?, ?)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec("001", "Sample Org 001")
	if err != nil {
		log.Fatal(err)
	}
}
