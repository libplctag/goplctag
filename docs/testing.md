# Testing

> NB: In progress

Todo

- check func signatures with libplctag.h
- add configuration file to avoid hard coded values
- add config file to gitignore
- calculate coverage
- create tests for everything

## Running tests

- configure settings for plc

config.go

```
package goplctag

const protocol = "protocol=ab_eip&"
const gateway = "gateway=192.168.1.14&path=1,1&"
const plcType = "plc=controllogix&"
const TIMEOUT = 5000
```

```
go test --cover
```
