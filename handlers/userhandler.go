package handlers

import (
	"context"

	database "hosp.com/configs"
	"hosp.com/helper"
	"hosp.com/models"

	"github.com/gofiber/fiber/v2"

	"go.mongodb.org/mongo-driver/bson"
)

func Getuser(c *fiber.Ctx) error {
	coll := database.GetCollection("user")
	id := c.Params("id")
	var user models.User
	err := coll.FindOne(context.Background(), bson.M{"id": id}).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	}
	return c.JSON(user)
}
func Createuser(c *fiber.Ctx) error {
	coll := database.GetCollection("user")
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	user.ID = helper.Genid()
	res, err := coll.InsertOne(context.Background(), user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(res)
}
func Updateuser(c *fiber.Ctx) error {
	coll := database.GetCollection("user")
	id := c.Params("id")
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	update := bson.M{
		"$set": user,
	}
	result, err := coll.UpdateOne(context.Background(), bson.M{"id": id}, update)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update user"})
	}
	return c.JSON(result)
}
func Deleteuser(c *fiber.Ctx) error {
	coll := database.GetCollection("user")
	id := c.Params("id")
	result, err := coll.DeleteOne(context.Background(), bson.M{"id": id})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to Delete user"})
	}
	return c.JSON(result)
}
