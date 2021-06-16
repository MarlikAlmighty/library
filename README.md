# Library

### A simple example of a server API, with a clean architecture.

***

[![Build Status](https://api.travis-ci.org/MarlikAlmighty/library.svg?branch=master&status=passed)](https://travis-ci.org/MarlikAlmighty/library) &nbsp;
[![Open Issues](https://img.shields.io/github/issues/google/fresnel)](https://github.com/MarlikAlmighty/library/issues)  &nbsp; 
[![License](https://img.shields.io/badge/License-MIT%201.0-orange.svg)](https://github.com/MarlikAlmighty/library/blob/master/LICENSE) &nbsp; 

***

I chose the library as a reference. We can manage books, bookcases.

***

### How to run
```sh
$ docker stack deploy -c stack.yml postgres 
$ docker stack ls

OR

$ docker-compose -f stack.yml up

$ export LIBRARY_HTTP_PORT=3000

$ export LIBRARY_DB=postgres://user:password@localhost:5432/library?sslmode=disable

$ migrate -database ${LIBRARY_DB} -path ./internal/repository/postgresql/migrations up

$ go run ./cmd/...
```

### Docker
```shell
$ docker build .
```

### Documentation: 
```sh
$ swagger serve ./swagger-api/swagger.yml
```
