# go-ignite-client
## Apache Ignite (GridGain) go language native client and SQL driver

### Requirements:
- Apache Ignite v1.3+
- go v1.8+

### Road map:
1. Develop SQL driver (`ignite-sql-http`) for Apache Ignite HTTP REST API (In progress)
2. Develop SQL driver (`ignite-sql-native`) for Apache Ignite protocol (Not started)

### Issues:
- `ignite-sql-http` SQL driver does not support transactions (Ignite HTTP REST API does not support transactions)
- Fields with type Time and Date are not supported yet (will be fixed soon)
- Fields with type Binary are not supported yet (will be fixed soon)
- Only 4 methods (are needed for SQL driver) of REST API are implemented now

### SQL driver (`ignite-sql-http`) for Apache Ignite HTTP REST API:
```
go get -u github.com/amsokol/go-ignite-client/sql/http
```
See [example](https://github.com/amsokol/go-ignite-client/tree/master/cmd/example-http-sql)
#### Data base configuration string json format:
Example:
```json
{
    "version": 2,
    "servers" : [
        "http://server1:8080/ignite",
        "http://server2:8080/ignite",
        "http://server3:8080/ignite"
    ],
    "username" : "myLogin",
    "password": "myPassword",
    "cache" : "Person",
    "pageSize": 1000
}
```
Specification:
| Parameter  | Type    | Mandatory | Default | Description
|------------|---------|-----------|---------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| version    | number  | no        | `1`     | Ignite HTTP REST API version you are going to use. `1` is for `1.x.x` server, `2` is for 2.x.x server                                                                                                                                               |
| servers    | URLs    | yes       |         | Ignite server list. Client automatically reconnects to another server if current server become unavailable                                                                                                                                          |
| username   | string  | no        |         | User name to connect to servers. Client supports HTTP Basic Authentication                                                                                                                                                                          |
| password   | string  | no        |         | Password to connect to servers. Client supports HTTP Basic Authentication                                                                                                                                                                           |
| cache      | string  | yes       |         | Ignite cache name as the default schema for SQL query. But I recommend provide table schema (cache name) in SQL query explicitly                                                                                                                    |
| pageSize   | int     | no        | `1000`  | Pagination in Ignite is a mechanism to avoid fetching the whole data set from server nodes to the client. I.e., while you iterate through the Rows, the client will fetch data in chunks. The size of each chunk is defined by `pageSize` property. |

#### Database fields mapping (Ignite -> golang and visa verse):
| Ignite (Java) type  | golang type |
|---------------------|-------------|
| java.lang.Byte      | int8        |
| java.lang.Short     | int16       |
| java.lang.Integer   | int32       |
| java.lang.Long      | int64       |
| java.lang.Double    | float64     |
| java.lang.Float     | float32     |
| java.lang.Boolean   | bool        |
| java.lang.Character | []rune      |
| java.lang.String    | string      |

### Apache Ignite HTTP REST API client:
```
go get -u github.com/amsokol/go-ignite-client/http
```
See [example](https://github.com/amsokol/go-ignite-client/tree/master/cmd/example-http-client)
