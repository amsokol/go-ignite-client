## Example how to use REST API client for Apache Ignite (GridGain)
Instruction for `Apache Ignite 2.0.0`:
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
