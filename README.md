# ibanApi
exposes one endpoint /iban/:iban where :iban must be replaced with the IBAN number one wants to validate. IBAN number to vlidate must be without whitespaces.
- Valid IBAN number gives 200 OK success status response code and empty body
- Invalid IBAN number gives 400 status reponse code and body contains an error message in json format
- Invalid endpoint gives 404 status response code

## Contents
- [What is checked]("what-is-checked)
- [Requirement](#requirement)
- [Run unit tests](#run-unit-tests)
- [Build Docker Image](#build-docker-image)
- [Run API Docker Image](#run-api-with-Docker)
- [API Examples](#api-examples)
- [Improvements](#improvements)

## What is checked
The validation of IBAN number is done with info found from [Wikipedia](https://en.wikipedia.org/wiki/International_Bank_Account_Number#Algorithms)
- Does not contain more than 34 character
- Have at least 4 character
- The first 2 characters are correct iban country code
- Have correct number of characters for the specified country
- The check digit are digits
> List of Iban Countries and length of the IBAN number was collected from [iban.com](https://www.iban.com/structure)

## Requirement
To build and run [Docker](https://docs.docker.com/get-docker/) is recommended

## Run unit tests
if git root for ibanApi is at /repo/ibanApi
```sh
$ docker run -it --rm -v /repo/ibanApi:/go/src golang:1.16-alpine
/go # apk add build-base
/go # cd src/iban/
/go/src/iban # go test -v
```

## Build Docker Image
Stand in git root directory for ibanApi
```sh
$ docker build -f DockerFile --tag iban-api iban
```

## Run API Docker Image
```sh
$ docker run --rm -p 80:8080 iban-api
```

## API Examples
A valid IBAN number
```sh
$ curl http://localhost/iban/AL47212110090000000235698741
```
> Returns 200 OK with empty body

An invalid IBAN number with invalid country code
```sh
$ curl http://localhost/iban/ZZ47212110090000000235698741
```
> Returns 400 Bad Request with body {"error":"iban does not have a valid country code"}

An invalid IBAN number with wrong length for Albania
```sh
$ curl http://localhost/iban/AL4721211009000000023569
```
> Returns 400 Bad Request with body {"error":"iban does not have a correct length"}

An invalid IBAN number with invalid characters
```sh
$ curl http://localhost/iban/AL472121100900000q0235698741
```
> Returns 400 Bad Request with body {"error":"iban does contain invalid characters"}

## Improvements
- Add digit check described on [Wikipedia](https://en.wikipedia.org/wiki/International_Bank_Account_Number#Algorithms)
- Add checks for BBAN format specified for the country code
- Move info on Iban Countries to a config file and read it at startup