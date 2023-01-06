package main

import (
	"api-gorm/config"
	"api-gorm/routes"
)


type Product struct {
	ID    int             `form:"id" json:"id"`
	Name  string          `form:"name" json:"name"`
	Description  string 	`form:"code" json:"description"`
	Stock       int    		`json:"stock"`
}
type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func main() {
	config.Connect()
	routes.HandleRequests()
}
