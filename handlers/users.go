package handlers

import (

	"github.com/fsmardani/go-for-example/database"
	"github.com/fsmardani/go-for-example/models"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"

)

func ListUsers(c *fiber.Ctx) error {
	users := []models.User{}
	database.DB.Db.Find(&users)
	log.Info().Msg("Hello from Zerolog global logger")

	return c.JSON(users)

}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		log.Error().Msg(err.Error())

		return c.JSON(err)
	}

	result := database.DB.Db.Create(&user)
	if result.Error != nil {
		log.Error().Msg(result.Error.Error())

		return c.JSON(result.Error)
	}

	return ListUsers(c)
}

func ShowUser(c *fiber.Ctx) error {
	user := models.User{}
	id := c.Params("id")

	result := database.DB.Db.Where("id = ?", id).First(&user)
	if result.Error != nil {
		log.Error().Msg(result.Error.Error())
		return c.JSON(result.Error)
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func EditUser(c *fiber.Ctx) error {
	user := models.User{}
	id := c.Params("id")

	result := database.DB.Db.Where("id = ?", id).First(&user)
	if result.Error != nil {
		log.Error().Msg(result.Error.Error())
		return c.JSON(result.Error)
	}

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	user := models.User{}
	id := c.Params("id")

	result := database.DB.Db.Where("id = ?", id).Delete(&user)
	if result.Error != nil {
		log.Error().Msg(result.Error.Error())
		return c.JSON(result.Error)
	}

	return ListUsers(c)
}
