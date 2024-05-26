package handlers

import (
	"github.com/fsmardani/go-for-example/database"
	"github.com/fsmardani/go-for-example/models"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/google/uuid"

	// "github.com/rs/zerolog"
)


func ListUsers(c *fiber.Ctx) error {
	users := []models.User{}
	database.DB.Db.Get(&users, `SELECT * FROM users`)
	log.Info().Str("service", "APP").Msg("users list")

	return c.Status(fiber.StatusAccepted).JSON(users)

}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		log.Error().Str("service", "APP").Msg(err.Error())

		return c.JSON(err)
	}
	user.Password, _ = HashPassword(user.Password)
	user.ID = uuid.New()
	_, err := database.DB.Db.NamedExec(`INSERT INTO users (id, name, age, phone_number, email, password) VALUES (:id, :name, :age, :phone_number, :email, :password)`, user)
	if err != nil {
		log.Error().Str("service", "APP")

		return c.Status(fiber.StatusAccepted).JSON(err)
	}

	return ListUsers(c)
}

func ShowUser(c *fiber.Ctx) error {
	user := models.User{}
	id := c.Params("id")

	_ = database.DB.Db.Get(&user, "SELECT * FROM users WHERE id = ? Limit 1", id)
	// if result.Error != nil {
	// 	log.Error().Str("service", "APP").Msg(result.Error())
	// 	return c.JSON(result.Error)
	// }

	return c.Status(fiber.StatusAccepted).Status(fiber.StatusOK).JSON(user)
}

func EditUser(c *fiber.Ctx) error {
	user := models.User{}
	id := c.Params("id")

	result := database.DB.Db.Get(&user, `SELECT * FROM users WHERE id = ? LIMIT 1`, id)
	if result.Error != nil {
		log.Error().Str("service", "APP").Msg(result.Error())
		return c.JSON(result.Error)
	}

	return c.Status(fiber.StatusAccepted).JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	// user := models.User{}
	id := c.Params("id")

	_, err := database.DB.Db.NamedExec(`DELETE FROM users WHERE id = ?`, id)
	if err != nil {
		log.Error().Str("service", "APP")
		return c.JSON(err)
	}

	return ListUsers(c)
}
