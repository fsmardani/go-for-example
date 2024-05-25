package handlers

import (
	"errors"
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	jtoken "github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"

	"github.com/fsmardani/go-for-example/database"
	"github.com/fsmardani/go-for-example/models"
	"github.com/fsmardani/go-for-example/config"

)

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func FindByCredentials(email, password string) (*models.User, error) {
	user := models.User{}
	
	result := database.DB.Db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		log.Error().Msg("user not found")
		return nil, errors.New("user not found")
	}
	passMatch := CheckPasswordHash(password, user.Password)
	if !passMatch{
		log.Error().Msg("password not match!")
		return nil, errors.New("password not match")
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
