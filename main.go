package main

import (
	"app/config"
	server "app/router"
	"fmt"
)

func main() {
	fmt.Println("Server started")
	config.DbConfiguration()
	router := server.Init()
	router.Run(":80")
}
