# Library

### Sample server code generation via go-swagger. 

[![Build Status](https://api.travis-ci.org/MarlikAlmighty/library.svg?branch=master&status=passed)](https://travis-ci.org/MarlikAlmighty/library) &nbsp;
[![Open Issues](https://img.shields.io/github/issues/google/fresnel)](https://github.com/MarlikAlmighty/library/issues)  &nbsp; 
[![License](https://img.shields.io/badge/License-MIT%201.0-orange.svg)](https://github.com/MarlikAlmighty/library/blob/master/LICENSE) &nbsp; 

***

### How to run
```sh
$ docker stack deploy -c stack.yml postgres 

OR

$ docker-compose -f stack.yml up

$ export POSTGRESQL_URL=postgres://user:password@localhost:5432/usecase?sslmode=disable

$ migrate -database ${POSTGRESQL_URL} -path postgresql/migrations up

$ go run ./cmd/usecase-server/... --port 3000
```

### Docker
```shell
$ docker build .
```

### Documentation: 
```sh
$ docs serve ./docs/swagger.yml
```
