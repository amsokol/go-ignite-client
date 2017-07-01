package main

import (
	"log"

	ignite "github.com/amsokol/go-ignite-client/http/v2"
	"net/url"
)

func main() {
	servers := []string{"http://localhost:8080/ignite"}

	c := ignite.NewClient(servers, "", "") // no login and password

	// get version
	v, _, err := c.GetVersion()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Version returned:")
	log.Println("version=", v)

	// show server log from line 10
	from := 10
	to := 15
	lg, _, err := c.GetLog("", &from, &to)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("")
	log.Println("Log returned:")
	log.Println("log=", lg)

	// decrement atomic long
	v64, nodeID, _, err := c.Decrement("Person", "sequence", nil, 10)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("")
	log.Println("Decrement returned:")
	log.Println("value=", v64)
	log.Println("affinityNodeId=", nodeID)

	// increment atomic long
	v64, nodeID, _, err = c.Increment("Person", "sequence", nil, 100)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("")
	log.Println("Increment returned:")
	log.Println("value=", v64)
	log.Println("affinityNodeId=", nodeID)

	/*
		// show metrics for Ignite cache
		m, nodeID, _, err := c.GetCacheMetrics("Person", "")
		log.Println("")
		log.Println("CacheMetrics returned:")
		log.Println("metrics=", m)
		log.Println("affinityNodeId=", nodeID)
	*/

	_, _, err = c.SQLFieldsQueryExecute("Person", 1000, `DELETE FROM Person`, url.Values{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("")
	log.Println("Deleted all rows from Person")

	_, _, err = c.SQLFieldsQueryExecute("Organization", 1000, `DELETE FROM Organization`, url.Values{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("")
	log.Println("Deleted all rows from Organization")

	_, _, err = c.SQLFieldsQueryExecute("Organization", 1000, `INSERT INTO Organization(_key, name) VALUES(1, 'Org 1')`, url.Values{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("")
	log.Println("Added one record to Organization")

	r, _, err := c.SQLFieldsQueryExecute("Organization", 1000, `SELECT _key, name FROM  Organization`, url.Values{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("")
	log.Println("Organizations:")
	log.Println(r)
}
