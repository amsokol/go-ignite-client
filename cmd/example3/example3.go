package main

import (
	"database/sql"
	"log"

	_ "github.com/amsokol/go-ignite-client/sql/http"
)

func main() {
	db, err := sql.Open("ignite-sql-http", `{
		"version" : 1.0,
		"servers" : ["http://localhost:8081/ignite", "http://localhost:8080/ignite"],
		"username" : "login",
		"password" : "password",
		"cache" : "Person",
		"pageSize" : 10,
		"quarantine" : 5}`)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
