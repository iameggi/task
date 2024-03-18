package controller

import (
	"errors"
	"net/http"
	"service-user/config"
	"service-user/helpers"
	"service-user/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type WebResponse struct {
	Code   int
	Status string
	Data   interface{}
}

func Register(c *fiber.Ctx) error {
	var requestBody model.User

	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid request body")
	}

	requestBody.ID = uuid.New().String()

	db := config.GetPostgresDB()
	hashedPassword := helpers.HashPassword([]byte(requestBody.Password))
	_, err := db.Exec("INSERT INTO users (id, email, password) VALUES ($1, $2, $3)",
		requestBody.ID,
		requestBody.Email,
		hashedPassword)

	if err != nil {
		panic(err)
	}

	return c.JSON(WebResponse{
		Code:   201,
		Status: "OK",
		Data:   requestBody.Email,
	})
}

func Login(c *fiber.Ctx) error {
	var requestBody model.User
	var result model.User

	c.BodyParser(&requestBody)

	db := config.GetPostgresDB()

	row := db.QueryRow("SELECT email, password FROM users WHERE email = $1", requestBody.Email)
	if err := row.Scan(&result.Email, &result.Password); err != nil {
		return c.Status(http.StatusUnauthorized).JSON(WebResponse{
			Code:   401,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}

	checkPassword := helpers.ComparePassword([]byte(result.Password), []byte(requestBody.Password))
	if !checkPassword {
		return c.Status(http.StatusUnauthorized).JSON(WebResponse{
			Code:   401,
			Status: "BAD_REQUEST",
			Data:   errors.New("invalid password").Error(),
		})
	}

	access_token := helpers.SignToken(requestBody.Email)

	return c.JSON(struct {
		Code        int         `json:"code"`
		Status      string      `json:"status"`
		AccessToken string      `json:"access_token"`
		Data        interface{} `json:"data"`
	}{
		Code:        200,
		Status:      "OK",
		AccessToken: access_token,
		Data:        result,
	})
}

func Auth(c *fiber.Ctx) error {
	return c.JSON("OK")
}
