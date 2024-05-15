package main

import (
	"github.com/gofiber/fiber/v2"
	database "hosp.com/configs"
	"hosp.com/routes"
)

func main() {
	app := fiber.New()
	database.DBconnect()
	defer database.DBdisconnect()
	routes.Users(app)
	app.Listen(":3000")
}
