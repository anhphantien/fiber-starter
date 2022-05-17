package main

import (
	"fiber-starter/database"
	_ "fiber-starter/docs"
	"fiber-starter/models"
	"fiber-starter/routes"
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
		log.Panic("Can't connect database:", err.Error())
	}

	database.DBConn.AutoMigrate(&models.Book{})

	app := routes.New()
	log.Fatal(app.Listen(":3000"))
}
