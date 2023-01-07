package main

import (
	"api-gorm/config"
	"api-gorm/routes"
)

func main() {
	config.Connect()
	routes.HandleRequests()
}
