package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	router := gin.Default()
	router.GET("/hostname", handleHostname)
	router.Run(":999")
}

func handleHostname(context *gin.Context) {
	hostname, err := os.Hostname()
	if err==nil {
		context.JSON(200, gin.H{
			"message": hostname,
		})
	} else {
		context.JSON(500, gin.H{
			"message": err.Error,
		})
	}
}