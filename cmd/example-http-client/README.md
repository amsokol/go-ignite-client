## Example how to use REST API client for Apache Ignite (GridGain)
### Instruction for `Apache Ignite 2.0.0`:
1. Download `Apache Ignite 2.0.0` from [official site](http://apache-mirror.rbc.ru/pub/apache//ignite/2.0.0/apache-ignite-fabric-2.0.0-bin.zip)

2. Extract distributive to any folder

3. Copy or move folder `ignite-rest-http` from `<path_with_ignite>/libs/optional/`  to `<path_with_ignite>/libs/`. It enables Ignite HTTP REST API.

4. `cd` to folder with current example files

5. Start Ignite server with configuration files from this example:
```shell
For Windows:
<path_with_ignite>\bin\ignite.bat .\example-http-client.xml

For Linux:
<path_with_ignite>/bin/ignite.sh ./example-http-client.xml
```
6. Run example:
```
go run main.go
```

### Source code:
```go
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
```
