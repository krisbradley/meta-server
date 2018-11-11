package main

import (
	"net"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/hostname", handleHostname)
	router.GET("/network", handleNetwork)
	router.GET("/network/interface", handleNetworkInterface)
	router.Run(":999")
}

func handleHostname(context *gin.Context) {
	hostname, err := os.Hostname()
	if err != nil {
		returnError(context, err)
	}
	returnOk(context, hostname)
}

func handleNetwork(context *gin.Context) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		returnError(context, err)
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				returnOk(context, ipnet.IP.String())
			}
		}
	}
}

func handleNetworkInterface(context *gin.Context) {
	interfaces, err := net.Interfaces()
	if err != nil {
		returnError(context, err)
	}
	returnOk(context, interfaces)
}

func returnOk(context *gin.Context, message interface{}) {
	context.JSON(200, message)
}

func returnError(context *gin.Context, err error) {
	context.JSON(500, err.Error)
}
