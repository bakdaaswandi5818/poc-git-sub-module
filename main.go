package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	// Import libraries from submodules
	greeting "github.com/bakdaaswandi5818/greeting-lib"
	logger "github.com/bakdaaswandi5818/logger-lib"
)

func main() {
	// Create a custom logger from the logger submodule
	appLogger := logger.NewLogger("POC-Server")
	appLogger.Info("Starting Git Submodule POC application")

	// Create a new Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Create a greeter instance from the greeting submodule
	greeter := greeting.NewGreeter("Guest")
	appLogger.Info("Initialized greeting service")

	// Routes
	e.GET("/", func(c echo.Context) error {
		appLogger.Debug("Root endpoint accessed")
		return c.JSON(http.StatusOK, map[string]string{
			"message":       "Git Submodule POC with Golang Echo Framework",
			"greeting_info": greeting.GetInfo(),
			"logger_info":   logger.GetVersion(),
		})
	})

	e.GET("/greet", func(c echo.Context) error {
		name := c.QueryParam("name")
		appLogger.Info("Greet endpoint accessed for: " + name)
		return c.JSON(http.StatusOK, map[string]string{
			"greeting": greeter.Greet(name),
		})
	})

	e.GET("/greet/time-based", func(c echo.Context) error {
		name := c.QueryParam("name")
		appLogger.Info("Time-based greet endpoint accessed for: " + name)
		return c.JSON(http.StatusOK, map[string]string{
			"greeting": greeter.GetTimeBasedGreeting(name),
		})
	})

	e.GET("/health", func(c echo.Context) error {
		appLogger.Debug("Health check endpoint accessed")
		return c.JSON(http.StatusOK, map[string]string{
			"status":            "healthy",
			"greeting_submodule": "greeting-lib loaded successfully",
			"logger_submodule":   "logger-lib loaded successfully",
		})
	})

	// Start server
	appLogger.Info("Server starting on :8080")
	e.Logger.Fatal(e.Start(":8080"))
}
