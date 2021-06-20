# Library

### A simple example of a server API, with a clean architecture.

***

[![Build Status](https://api.travis-ci.com/MarlikAlmighty/library.svg?branch=master&status=passed)](https://travis-ci.org/MarlikAlmighty/library) &nbsp;
[![Open Issues](https://img.shields.io/github/issues/google/fresnel)](https://github.com/MarlikAlmighty/library/issues)  &nbsp; 
[![License](https://img.shields.io/badge/License-MIT%201.0-orange.svg)](https://github.com/MarlikAlmighty/library/blob/master/LICENSE) &nbsp; 

***

I chose the library as a reference. We can manage books, bookcases.

***

### How to run
```sh
$ docker stack deploy -c postgresql.yml postgresql

$ export PREFIX=LIBRARY

$ export LIBRARY_HTTP_PORT=3000

$ export LIBRARY_DB=postgres://user:password@localhost:5432/library?sslmode=disable

$ export LIBRARY_MIGRATE=true

$ export LIBRARY_PATH_TO_MIGRATE="./migrations"

$ go run ./cmd/...
```

### Docker
```sh
$ docker build .
```

### Documentation: 
```sh
$ swagger serve ./swagger-api/swagger.yml
```

### How to generate project:
```sh
$ swagger generate server --spec ./swagger-api/swagger.yml \ 
--target ./internal/gen -C ./swagger-templates/default-server.yml \
--template-dir ./swagger-templates --name library
```