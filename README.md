# ibanApi
exposes one endpoint /iban/:iban where :iban must be replaced with the IBAN number one wants to validate.
- Valid IBAN number gives 200 OK success status response code and empty body
- Invalid IBAN number gives 400 status reponse code and body contains an error message in json format
- Invalid endpoint gives 404 status response code


## Contents
- [What is checked]("what-is-checked)
- [Requirement](#requirement)
- [Build Docker Image](#build-docker-image)
- [Run API Docker Image](#run-api-with-Docker)
- [API Examples](#api-examples)

## What is checked
The validation of IBAN number is done with info found from [Wikipedia](https://en.wikipedia.org/wiki/International_Bank_Account_Number#Algorithms)
- Does not contain more than 34 character
- Have at least 4 character
- The first 2 characters are correct iban country code
- Have correct number of characters for the specified country
- The check digit are digits
> Only 2 Iban countries are valid AL, AD at this moment

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