# ibanApi
exposes one endpoint /iban/:iban where :iban must be replaced with the IBAN number one wants to validate. If the IBAN number is valid a 200 OK success status response code is received. For an invalid endpoint 404 status response code is received.
 
## Contents
- [Requirement](#requirement)
- [Build Docker Image](#build-docker-image)
- [Run API Docker Image](#run-api-with-Docker)
- [API Examples](#api-examples)

## Requirement
To build and run [Docker](https://docs.docker.com/get-docker/) is recommended

## Build Docker Image
Stand in git root directory for ibanApi
```sh
$ docker build -f DockerFile  --tag iban-api iban
```

## Run API Docker Image
```sh
$ docker run --rm -p 80:8080 iban-api
```

## API Examples
```sh
$ curl http://localhost/iban/AL47212110090000000235698741
```