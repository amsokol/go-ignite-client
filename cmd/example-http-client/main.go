package main

import (
	"log"

	"github.com/amsokol/go-ignite-client/http/v2"
	"net/url"
)

func main() {
	servers := []string{"http://localhost:8080/ignite"}
	quarantine := 10.0 // 10 mins

	c := v2.Open(servers, quarantine, "", "") // no login and password

	v, _, err := c.GetVersion()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server version is", v)

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
