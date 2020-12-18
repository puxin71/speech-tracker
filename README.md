## Synopsis

This application is a simple REST API server that exposes endpoints to allow accessing and manipulating the "talk" database records. The operations that the endpoint will allow include:
* Fetching all talks and their attendants
* Fetching one talk and its attendees

## API Reference
The APIs should:
* Fetch all talks and their attendants in response to a valid GET request at `/talks`
* Fetch all attendees that are registered to a talk in response to a valid GET request at `/talks/[0-9]+/attendees`

## Code Example



## Motivation

This project shows that with proper layering we can easily integrate an existing Web Service with different database technologies, i.e:
* local file system
* Cloud object store
* MySQL
* Postgresql
* or other Cloud hosted databases

## Installation

For the short term, you have to fetch the source code from the [Github repo] (https://github.com/puxin71/speech-tracker) and run it as:
```
RESOURCES="./resources" go run server.go
```

Eventually we can launch the server and an embedded "database" using Docker
## Tests

Run all unit tests as:
```
go test -v ./...
```

Eventually this step is automated in a Makefile
