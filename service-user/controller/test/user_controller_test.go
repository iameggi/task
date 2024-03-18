package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	app := fiber.New()

	requestBody := model.User{
		Email:    "test@example.com",
		Password: "password123",
	}
	jsonBytes, _ := json.Marshal(requestBody)
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(jsonBytes))
	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()

	
	app.Post("/register", Register)
	app.ServeHTTP(res, req)
	assert.Equal(t, http.StatusCreated, res.Code)

	var responseBody WebResponse
	err := json.Unmarshal(res.Body.Bytes(), &responseBody)
	assert.NoError(t, err)
	assert.Equal(t, "OK", responseBody.Status)
	assert.Equal(t, "test@example.com", responseBody.Data)

}


