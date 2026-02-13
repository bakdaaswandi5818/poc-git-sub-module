package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	// Import the greeting library from the submodule
	greeting "github.com/bakdaaswandi5818/greeting-lib"
)

func main() {
	// Create a new Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Create a greeter instance
	greeter := greeting.NewGreeter("Guest")

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Git Submodule POC with Golang Echo Framework",
			"info":    greeting.GetInfo(),
		})
	})

	e.GET("/greet", func(c echo.Context) error {
		name := c.QueryParam("name")
		return c.JSON(http.StatusOK, map[string]string{
			"greeting": greeter.Greet(name),
		})
	})

	e.GET("/greet/time-based", func(c echo.Context) error {
		name := c.QueryParam("name")
		return c.JSON(http.StatusOK, map[string]string{
			"greeting": greeter.GetTimeBasedGreeting(name),
		})
	})

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status":    "healthy",
			"submodule": "greeting-lib loaded successfully",
		})
	})

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
