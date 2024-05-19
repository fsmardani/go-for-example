package handlers

import (
	"errors"
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	jtoken "github.com/golang-jwt/jwt/v4"
	"github.com/fsmardani/go-for-example/database"
	"github.com/fsmardani/go-for-example/models"
	"github.com/fsmardani/go-for-example/config"

)

func FindByCredentials(email, password string) (*models.User, error) {
	user := models.User{}

	result := database.DB.Db.Where("email = ? AND password = ?", email, password).First(&user)
	if result.Error != nil {
		log.Error().Msg("user not found")
		return nil, errors.New("user not found")
	}
	return &user,result.Error
   }


func Login(c *fiber.Ctx) error {
	loginRequest := new(models.LoginRequest)
	if err := c.BodyParser(loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	user, err := FindByCredentials(loginRequest.Email, loginRequest.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	day := time.Hour * 24

	claims := jtoken.MapClaims{
		"ID":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(day * 1).Unix(),
	}

	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(models.LoginResponse{
		Token: t,
	})
}
