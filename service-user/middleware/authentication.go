package middleware

import (
	"fmt"
	"service-user/config"
	"service-user/helpers"
	"service-user/model"

	"github.com/gofiber/fiber/v2"
)

func Authentication(c *fiber.Ctx) error {

	db := config.GetPostgresDB()
	access_token := c.Get("access_token")

	if len(access_token) == 0 {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid token: Access token missing")
	}

	checkToken, err := helpers.VerifyToken(access_token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid token: Failed to verify token")
	}

	fmt.Println(checkToken, "CEKKKK", checkToken["email"])

	var user model.User
	err = db.QueryRow("SELECT email, password FROM users WHERE email = $1", checkToken["email"]).Scan(&user.Email, &user.Password)
	if err != nil {
		fmt.Println(err, "Error fetching user from database")
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid token: User not found")
	}

	c.Locals("user", user)

	return c.Next()
}
