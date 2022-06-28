package main

import (
	"fiber-starter/database"
	_ "fiber-starter/docs"
	"fiber-starter/env"
	"fiber-starter/routers"
	"log"
)

// @title Fiber starter
// @version 1.0
// @description Fiber starter's API documentation
// @BasePath /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	if err := database.Connect(); err != nil {
		log.Panic("Can't connect to database: ", err.Error())
	}

	app := routers.New()
	log.Fatal(app.Listen(":" + env.PORT))
}
