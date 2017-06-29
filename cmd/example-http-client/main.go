package main

import (
	"log"

	ignite "github.com/amsokol/go-ignite-client/http/v2"
	"net/url"
)

func main() {
	servers := []string{"http://localhost:8080/ignite"}

	c := ignite.Open(servers, "", "") // no login and password

	v, _, err := c.Version()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server version is", v)

	// show server log from line 10
	lg, _, err := c.Log("", 10, 15)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server log is", lg)

	_, _, err = c.SQLFieldsQueryExecute("Person", 1000, `DELETE FROM Person`, url.Values{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Deleted all rows from Person")
	log.Println("")

	_, _, err = c.SQLFieldsQueryExecute("Organization", 1000, `DELETE FROM Organization`, url.Values{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Deleted all rows from Organization")
	log.Println("")

	_, _, err = c.SQLFieldsQueryExecute("Organization", 1000, `INSERT INTO Organization(_key, name) VALUES(1, 'Org 1')`, url.Values{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Added one record to Organization")
	log.Println("")

	r, _, err := c.SQLFieldsQueryExecute("Organization", 1000, `SELECT _key, name FROM  Organization`, url.Values{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Organizations:")
	log.Println(r)
}
