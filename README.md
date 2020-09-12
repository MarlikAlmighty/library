# Library

### Sample server code generation via go-swagger. 

[![Build Status](https://api.travis-ci.org/MarlikAlmighty/library.svg?branch=master&status=passed)](https://travis-ci.org/MarlikAlmighty/library) &nbsp; 
[![Contributing](https://img.shields.io/badge/contributions-welcome-brightgreen)](https://github.com/MarlikAlmighty/library/blob/master/CONTRIBUTING.md)  &nbsp; 
[![Open Issues](https://img.shields.io/github/issues/google/fresnel)](https://github.com/MarlikAlmighty/library/issues)  &nbsp; 
[![License](https://img.shields.io/badge/License-MIT%201.0-orange.svg)](https://github.com/MarlikAlmighty/library/blob/master/LICENSE) &nbsp; 

***

### How to run
```sh
$ docker stack deploy -c stack.yml postgres 

OR

$ docker-compose -f stack.yml up

$ export POSTGRESQL_URL=postgres://user:password@localhost:5432/library?sslmode=disable

$ migrate -database ${POSTGRESQL_URL} -path db/migrations up

$ go run ./cmd/library-server/... --port 3000
```

### Docker
```shell
$ docker build .
```

### Documentation: 
```sh
$ swagger serve ./swagger/swagger.yml
```
