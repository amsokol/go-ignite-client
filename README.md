# go-ignite-client
## Apache Ignite (GridGain) go language native client and SQL driver

### Requirements:
- Apache Ignite v1.3+
- go v1.8+

### Roadmap:
1. Develop SQL driver (`ignite-sql-http`) for Apache Ignite HTTP REST API (In progress)
2. Develop SQL driver (`ignite-sql-native`) for Apache Ignite protocol (Not started)

### Issues:
- `ignite-sql-http` SQL driver does not support transactions (Ignite HTTP REST API does not support transactions)
- Fields with type Time and Date are not supported yet (will be fixed soon)
- Fields with type Binary are not supported yet (will be fixed soon)

### SQL driver (`ignite-sql-http`) for Apache Ignite HTTP REST API:
```
go get -u github.com/amsokol/go-ignite-client/sql/http
```
See [example](https://github.com/amsokol/go-ignite-client/tree/master/cmd/example-http-sql)

### Apache Ignite HTTP REST API client:
```
go get -u github.com/amsokol/go-ignite-client/http
```
See [example](https://github.com/amsokol/go-ignite-client/tree/master/cmd/example-http-client)
