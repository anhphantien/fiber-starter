package main

import (
	"fiber-starter/database"
	_ "fiber-starter/docs"
	"fiber-starter/env"
	"fiber-starter/routers"
	"log"
)

// @title       Fiber starter
// @version     1.0
// @description Fiber starter's API documentation

// @securityDefinitions.apikey Bearer
// @in   header
// @name Authorization
func main() {
	database.Connect()
	app := routers.New()
	log.Fatal(app.Listen(":" + env.PORT))
}
