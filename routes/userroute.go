package routes

import (
	"github.com/gofiber/fiber/v2"
	"hosp.com/handlers"
)

func Users(app *fiber.App) {
	userGroup := app.Group("/user")

	userGroup.Get("/:id", handlers.Getuser)

	userGroup.Post("", handlers.Createuser)
	userGroup.Put("/:id", handlers.Updateuser)
	userGroup.Delete("/:id", handlers.Deleteuser)
}
