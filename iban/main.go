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
	if err := validateIban(c.Param("iban")); err == nil {
		c.Status(http.StatusOK)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
	}
}