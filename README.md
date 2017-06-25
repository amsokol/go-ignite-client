# go-ignite-client
## Apache Ignite (GridGain) go language native client and SQL driver
```
go get -u github.com/amsokol/go-ignite-client
```

Roadmap:
1. Develop SQL driver ("ignite-sql-http") based on Apache Ignite HTTP REST API (In progress)
2. Develop SQL driver ("ignite-sql-native") based on native Apache Ignite protocol (Not started)

Example #1 (INSERT) how to use:
```go
package main

import (
	"database/sql"
	"log"

	_ "github.com/amsokol/go-ignite-client/sql"
)

func main() {
	db, err := sql.Open("ignite-sql-http",
		`{"servers" : ["http://localhost:8080/ignite"], "cache" : "Person"}`)
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

	_, err = stmt.Exec("111", "Sample Org")
	if err != nil {
		log.Fatal(err)
	}
}
```

Example #2 (SELECT) how to use:
```go
package main

import (
	"database/sql"
	"log"

	_ "github.com/amsokol/go-ignite-client/sql"
)

func main() {
	db, err := sql.Open("ignite-sql-http",
		`{"servers" : ["http://localhost:8080/ignite"], "cache" : "Person", "pageSize" : 10}`)
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
```
