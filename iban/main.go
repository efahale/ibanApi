package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	initApi()
}

func initApi() {
	router := gin.Default()
	router.GET("/iban/:iban", ibanEndpoint)
	router.Run()
}

func ibanEndpoint(c *gin.Context) {
	if validateIban(c.Param("iban")) {
		c.Status(http.StatusOK)
	} else {
		c.Status(http.StatusBadRequest)
	}
}