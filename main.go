package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"net"
)

func main() {
	router := gin.Default()
	router.GET("/hostname", handleHostname)
	router.GET("/network/interface", handleNetworkInterface)
	router.Run(":999")
}

func handleHostname(context *gin.Context) {
	hostname, err := os.Hostname()
	if err!=nil {
		returnError(context, err)
	}
	context.JSON(200, gin.H{
		"message": hostname,
	})
}

func handleNetworkInterface(context *gin.Context) {
	l, err := net.Interfaces()
	if err != nil {
		returnError(context, err)
	}
	context.JSON(200, gin.H{
		"message": l,
	})
}

func returnError(context *gin.Context, err error) {
	context.JSON(500, gin.H{
		"message": err.Error,
	})
}