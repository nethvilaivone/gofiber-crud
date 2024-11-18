package main

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func getEVN(c *fiber.Ctx) error {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "defaut value...."
	}
	return c.JSON(fiber.Map{
		"SECRET": secret,
	})
}

func checkMideeleWare(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if claims["role"] != "admin" {
		return fiber.ErrUnauthorized
	}
	return c.Next()
}

func login(c *fiber.Ctx) error {

	user := new(User)
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if user.Email != userMember.Email {
		return c.Status(fiber.StatusBadRequest).SendString("email incorect")
	}
	if user.Password != userMember.Password {
		return c.Status(fiber.StatusBadRequest).SendString("password incorrect")
	}
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["role"] = "admin"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"login": "sucessful",
		"token": t})

}
