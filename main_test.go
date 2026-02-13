package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	greeting "github.com/bakdaaswandi5818/greeting-lib"
)

func TestRootEndpoint(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Git Submodule POC with Golang Echo Framework",
			"info":    greeting.GetInfo(),
		})
	}

	if assert.NoError(t, handler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response map[string]string
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Git Submodule POC with Golang Echo Framework", response["message"])
		assert.Contains(t, response["info"], "Greeting Library")
	}
}

func TestGreetEndpoint(t *testing.T) {
	e := echo.New()
	greeter := greeting.NewGreeter("Guest")

	req := httptest.NewRequest(http.MethodGet, "/greet?name=Alice", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := func(c echo.Context) error {
		name := c.QueryParam("name")
		return c.JSON(http.StatusOK, map[string]string{
			"greeting": greeter.Greet(name),
		})
	}

	if assert.NoError(t, handler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response map[string]string
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Contains(t, response["greeting"], "Alice")
	}
}

func TestGreetEndpointWithoutName(t *testing.T) {
	e := echo.New()
	greeter := greeting.NewGreeter("Guest")

	req := httptest.NewRequest(http.MethodGet, "/greet", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := func(c echo.Context) error {
		name := c.QueryParam("name")
		return c.JSON(http.StatusOK, map[string]string{
			"greeting": greeter.Greet(name),
		})
	}

	if assert.NoError(t, handler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response map[string]string
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Contains(t, response["greeting"], "Guest")
	}
}

func TestHealthEndpoint(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status":    "healthy",
			"submodule": "greeting-lib loaded successfully",
		})
	}

	if assert.NoError(t, handler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response map[string]string
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "healthy", response["status"])
		assert.Equal(t, "greeting-lib loaded successfully", response["submodule"])
	}
}
