# Extending the POC - Examples

This document provides examples of how to extend this proof of concept with additional features.

## Example 1: Adding a New Endpoint

### Add a Farewell Endpoint

1. **Add functionality to the greeting library:**

```go
// In libs/greeting-lib/greeting.go

// Farewell returns a farewell message
func (g *Greeter) Farewell(name string) string {
	if name == "" {
		name = g.defaultName
	}
	return fmt.Sprintf("Goodbye, %s! Thanks for trying the Git Submodule POC.", name)
}
```

2. **Add the endpoint to main.go:**

```go
e.GET("/farewell", func(c echo.Context) error {
	name := c.QueryParam("name")
	return c.JSON(http.StatusOK, map[string]string{
		"farewell": greeter.Farewell(name),
	})
})
```

3. **Test the new endpoint:**

```bash
curl "http://localhost:8080/farewell?name=Alice"
```

## Example 2: Adding Database Support

### Add SQLite Database

1. **Install dependencies:**

```bash
go get gorm.io/gorm
go get gorm.io/driver/sqlite
```

2. **Create a model:**

```go
// models/greeting.go
package models

import "gorm.io/gorm"

type GreetingLog struct {
	gorm.Model
	Name    string
	Message string
	Type    string // "greet" or "farewell"
}
```

3. **Initialize database in main.go:**

```go
import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// ... existing code ...
	
	// Initialize database
	db, err := gorm.Open(sqlite.Open("greetings.db"), &gorm.Config{})
	if err != nil {
		e.Logger.Fatal("Failed to connect to database")
	}
	
	// Auto migrate
	db.AutoMigrate(&models.GreetingLog{})
	
	// Modified greet endpoint with logging
	e.GET("/greet", func(c echo.Context) error {
		name := c.QueryParam("name")
		message := greeter.Greet(name)
		
		// Log to database
		db.Create(&models.GreetingLog{
			Name:    name,
			Message: message,
			Type:    "greet",
		})
		
		return c.JSON(http.StatusOK, map[string]string{
			"greeting": message,
		})
	})
	
	// ... rest of code ...
}
```

## Example 3: Adding Configuration File Support

### Use Environment Variables and Config File

1. **Install viper:**

```bash
go get github.com/spf13/viper
```

2. **Create config.yaml:**

```yaml
server:
  host: localhost
  port: 8080

greeting:
  default_name: Guest
  enable_logging: true
```

3. **Load configuration:**

```go
import "github.com/spf13/viper"

func loadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	
	// Set defaults
	viper.SetDefault("server.port", "8080")
	viper.SetDefault("greeting.default_name", "Guest")
	
	// Read environment variables
	viper.AutomaticEnv()
	
	return viper.ReadInConfig()
}

func main() {
	loadConfig()
	
	port := viper.GetString("server.port")
	defaultName := viper.GetString("greeting.default_name")
	
	// ... use in application ...
}
```

## Example 4: Adding More Submodules

### Add a Logger Submodule

1. **Create a new library in libs/logger-lib:**

```bash
mkdir -p libs/logger-lib
cd libs/logger-lib
git init
```

2. **Create the logger:**

```go
// libs/logger-lib/logger.go
package logger

import (
	"fmt"
	"time"
)

type Logger struct {
	prefix string
}

func NewLogger(prefix string) *Logger {
	return &Logger{prefix: prefix}
}

func (l *Logger) Info(message string) {
	fmt.Printf("[%s] [INFO] %s: %s\n", time.Now().Format(time.RFC3339), l.prefix, message)
}

func (l *Logger) Error(message string) {
	fmt.Printf("[%s] [ERROR] %s: %s\n", time.Now().Format(time.RFC3339), l.prefix, message)
}
```

3. **Add to .gitmodules:**

```ini
[submodule "libs/logger-lib"]
	path = libs/logger-lib
	url = ./libs/logger-lib
```

4. **Update go.mod:**

```go
replace (
	github.com/bakdaaswandi5818/greeting-lib => ./libs/greeting-lib
	github.com/bakdaaswandi5818/logger-lib => ./libs/logger-lib
)
```

5. **Use in application:**

```go
import logger "github.com/bakdaaswandi5818/logger-lib"

func main() {
	log := logger.NewLogger("POC")
	log.Info("Server starting...")
	
	// ... rest of code ...
}
```

## Example 5: Adding Middleware

### Add Request ID Middleware

```go
func requestIDMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			requestID := c.Request().Header.Get("X-Request-ID")
			if requestID == "" {
				requestID = fmt.Sprintf("%d", time.Now().UnixNano())
			}
			c.Response().Header().Set("X-Request-ID", requestID)
			return next(c)
		}
	}
}

func main() {
	e := echo.New()
	
	// Add custom middleware
	e.Use(requestIDMiddleware())
	
	// ... rest of code ...
}
```

## Example 6: Adding Prometheus Metrics

### Add Metrics Endpoint

1. **Install Prometheus client:**

```bash
go get github.com/prometheus/client_golang/prometheus
go get github.com/prometheus/client_golang/prometheus/promhttp
```

2. **Add metrics:**

```go
import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	greetingsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "greetings_total",
			Help: "Total number of greetings",
		},
		[]string{"type"},
	)
)

func init() {
	prometheus.MustRegister(greetingsCounter)
}

func main() {
	e := echo.New()
	
	// ... existing middleware ...
	
	// Metrics endpoint
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
	
	// Modified greet endpoint with metrics
	e.GET("/greet", func(c echo.Context) error {
		name := c.QueryParam("name")
		greetingsCounter.WithLabelValues("basic").Inc()
		return c.JSON(http.StatusOK, map[string]string{
			"greeting": greeter.Greet(name),
		})
	})
	
	// ... rest of code ...
}
```

## Example 7: Adding WebSocket Support

### Add Real-time Greetings

```go
import (
	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

func greetingWebSocket(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		
		for {
			var name string
			err := websocket.Message.Receive(ws, &name)
			if err != nil {
				break
			}
			
			greeting := greeter.Greet(name)
			err = websocket.Message.Send(ws, greeting)
			if err != nil {
				break
			}
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

func main() {
	e := echo.New()
	
	// ... existing code ...
	
	e.GET("/ws/greet", greetingWebSocket)
	
	// ... rest of code ...
}
```

## Example 8: Adding Swagger Documentation

### Auto-generate API Documentation

1. **Install swag:**

```bash
go install github.com/swaggo/swag/cmd/swag@latest
go get github.com/swaggo/echo-swagger
```

2. **Add swagger comments:**

```go
// @title Git Submodule POC API
// @version 1.0
// @description This is a sample server demonstrating Git submodules with Echo.
// @host localhost:8080
// @BasePath /

// @Summary Get welcome message
// @Description Get welcome message and submodule info
// @Tags root
// @Produce json
// @Success 200 {object} map[string]string
// @Router / [get]
func rootHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Git Submodule POC with Golang Echo Framework",
		"info":    greeting.GetInfo(),
	})
}
```

3. **Generate docs and add endpoint:**

```bash
swag init

# In main.go
import echoSwagger "github.com/swaggo/echo-swagger"

e.GET("/swagger/*", echoSwagger.WrapHandler)
```

## Testing Your Extensions

After adding any extension:

1. **Update tests:**
```bash
# Add tests for new functionality in *_test.go files
```

2. **Run tests:**
```bash
make test
```

3. **Build and run:**
```bash
make build
make run
```

4. **Test manually:**
```bash
curl http://localhost:8080/your-new-endpoint
```

## Best Practices for Extensions

1. **Keep submodules focused**: Each submodule should have a single, clear purpose
2. **Add tests**: Always add tests for new functionality
3. **Update documentation**: Keep README and other docs updated
4. **Use interfaces**: Make your code testable and mockable
5. **Handle errors**: Always handle errors appropriately
6. **Add logging**: Log important events and errors
7. **Consider performance**: Profile and optimize as needed

## Resources

- [Echo Framework Guide](https://echo.labstack.com/guide/)
- [Go Best Practices](https://golang.org/doc/effective_go)
- [Git Submodules](https://git-scm.com/book/en/v2/Git-Tools-Submodules)
- [GORM Documentation](https://gorm.io/)
- [Prometheus Go Client](https://github.com/prometheus/client_golang)
