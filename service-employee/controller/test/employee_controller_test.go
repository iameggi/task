package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	"service-employee/controller"
)

func TestCreateEmployee(t *testing.T) {

	app := fiber.New()

	controller.SetUserURI("http://localhost:3001/user")

	requestBody := map[string]string{
		"name": "John Doe",
	}
	jsonBytes, _ := json.Marshal(requestBody)
	req := httptest.NewRequest(http.MethodPost, "/employee", bytes.NewReader(jsonBytes))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("access_token", "sample_access_token")

	res := httptest.NewRecorder()

	app.Post("/employee", controller.CreateEmployee)
	app.ServeHTTP(res, req)
	assert.Equal(t, http.StatusCreated, res.Code)

}
