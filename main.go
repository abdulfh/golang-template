package main

import (
	"fmt"
	"golangbe/routers"
	"golangbe/utils/helper"
)

func main() {
	listenAddress := helper.GetEnv("SERVER_PORT", ":40001")
	fmt.Println("Starting listen address: ", listenAddress)
	routers.Server(listenAddress)
}
