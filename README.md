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
- `version` - type `number`, optional, default value is `1`. Ignite HTTP REST API version you are going to use. `1` is for `1.x.x` server, `2` is for 2.x.x server
- `servers` - type `URL`s, mandatory, no default value. Ignite server list. Client automatically reconnects to another server if current server become unavailable
- `username` - type `string`, optional, no default value. User name to connect to servers. Client supports HTTP Basic Authentication
- `password` - type `string`, optional, no default value. Password to connect to servers. Client supports HTTP Basic Authentication
- `cache`- type `string`, mandatory. Ignite cache name as the default schema for SQL query. But I recommend provide table schema (cache name) in each SQL query explicitly
- `pageSize`- type `int`, optional, default value is `1000`. Pagination in Ignite is a mechanism to avoid fetching the whole data set from server nodes to the client. I.e., while you iterate through the Rows, the client will fetch data in chunks. The size of each chunk is defined by `pageSize` property.

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
| java.sql.Timestamp  | time.Time   |
| java.util.Date      | time.Time   |

### Apache Ignite HTTP REST API client:
```
go get -u github.com/amsokol/go-ignite-client/http
```
See [example](https://github.com/amsokol/go-ignite-client/tree/master/cmd/example-http-client)
