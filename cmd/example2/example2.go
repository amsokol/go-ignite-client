package main

import (
	"database/sql"
	"log"

	_ "github.com/amsokol/go-ignite-client/sql"
)

func main() {
	db, err := sql.Open("ignite-sql-http",
		`{"servers" : ["http://localhost:8080/ignite"], "username" : "login", "password" : "password", "cache" : "Person", "pageSize" : 10}`)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(`select _key,name from "Organization".Organization`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var (
		id   int64
		name string
	)

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
