package main

import (
	"net"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/network", handleNetwork)
	router.GET("/network/interface", handleNetworkInterface)
	router.GET("/network/interface/:interfaceName", handleSingleNetworkInterface)
	router.Run(":999")
}

func handleNetwork(context *gin.Context) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		returnError(context, err)
	}
	ip := []string{}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = append(ip, ipnet.IP.String())
			}
		}
	}
	hostname, err := os.Hostname()
	if err != nil {
		returnError(context, err)
	}
	returnOk(context, gin.H{
		"hostname": hostname,
		"ip": ip,
	})
}

func handleNetworkInterface(context *gin.Context) {
	interfaces, err := net.Interfaces()
	if err != nil {
		returnError(context, err)
	}
	returnOk(context, interfaces)
}


func handleSingleNetworkInterface(context *gin.Context) {
	singleInterface, err := net.InterfaceByName(context.Param("interfaceName"))
	if err != nil {
		returnError(context, err)
	}
	returnOk(context, singleInterface)
}

func returnOk(context *gin.Context, message interface{}) {
	context.JSON(200, message)
}

func returnError(context *gin.Context, err error) {
	context.JSON(500, err.Error)
}
