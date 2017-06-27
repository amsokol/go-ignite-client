package main

import (
	"database/sql"
	"log"

	"fmt"
	_ "github.com/amsokol/go-ignite-client/sql/http"
)

func main() {
	db, err := sql.Open("ignite-sql-http", `{
		"version": 2,
		"servers" : [
			"http://localhost:8080/ignite"
		],
		"cache" : "Person"
	}`)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Ping for test
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Clear Person table
	_, err = db.Exec(`DELETE FROM "Person".Person`)
	if err != nil {
		log.Fatal(err)
	}

	// Clear Organization table
	_, err = db.Exec(`DELETE FROM "Organization".Organization`)
	if err != nil {
		log.Fatal(err)
	}

	// Add 10 Organizations
	stmt, err := db.Prepare(`INSERT INTO "Organization".Organization(_key, name) VALUES(?, ?)`)
	if err != nil {
		log.Fatal(err)
	}
	for i := 1; i <= 10; i++ {
		_, err = stmt.Exec(i, // _key
			fmt.Sprintf("Organization #%d", i)) // name
		if err != nil {
			log.Fatal(err)
		}
	}

	// Add 30 Persons
	stmt, err = db.Prepare(`INSERT INTO "Person".Person(_key, orgId, firstName, lastName, resume, salary)
		VALUES(?, ?, ?, ?, ?, ?)`)
	if err != nil {
		log.Fatal(err)
	}
	for i := 1; i <= 30; i++ {
		orgID := i % 10
		if orgID == 0 {
			orgID = 1
		}
		_, err = stmt.Exec(i, // _key
			orgID, // orgId
			fmt.Sprintf("FirstName%d", i),        // firstName
			fmt.Sprintf("LastName%d", i),         // lastName
			fmt.Sprintf("This is resume #%d", i), // resume
			float64(i)+100.0+float64(i)/30)       // salary
		if err != nil {
			log.Fatal(err)
		}
	}

	// Show all Organizations
	rows, err := db.Query(`SELECT _key, name FROM "Organization".Organization`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var (
		orgID   int64
		orgName string
	)
	log.Println("Organizations:")
	for rows.Next() {
		err := rows.Scan(&orgID, &orgName)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(orgID, orgName)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	// Show all Persons
	rows, err = db.Query(`SELECT _key, orgId, firstName, lastName, resume, salary FROM "Person".Person ORDER BY _key`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var (
		prsID        int64
		prsOrgID     int64
		prsFirstName string
		prsLastName  string
		prsResume    string
		prsSalary    float64
	)
	log.Println("")
	log.Println("Persons:")
	for rows.Next() {
		err := rows.Scan(&prsID, &prsOrgID, &prsFirstName, &prsLastName, &prsResume, &prsSalary)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(prsID, prsOrgID, prsFirstName, prsLastName, prsResume, prsSalary)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	// Show Person with Organization name
	rows, err = db.Query(`SELECT o.name, p.firstName, p.lastName, p.salary FROM "Person".Person as p
		INNER JOIN "Organization".Organization o
		ON p.orgId = o._key
		ORDER BY name`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	log.Println("")
	log.Println("Persons by Organization:")
	for rows.Next() {
		err := rows.Scan(&orgName, &prsFirstName, &prsLastName, &prsSalary)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(orgName, prsFirstName, prsLastName, prsSalary)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	// Show Person with Organization with filter
	rows, err = db.Query(`SELECT o.name, p.firstName, p.lastName, p.salary FROM "Person".Person as p
		INNER JOIN "Organization".Organization o
		ON p.orgId = o._key
		WHERE p.salary > ?
		ORDER BY name`, 115)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	log.Println("")
	log.Println("Persons who has salary more than 115 by Organization:")
	for rows.Next() {
		err := rows.Scan(&orgName, &prsFirstName, &prsLastName, &prsSalary)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(orgName, prsFirstName, prsLastName, prsSalary)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
