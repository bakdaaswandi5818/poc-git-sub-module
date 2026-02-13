# Quick Start Guide

This guide will help you get started with the Git Submodule POC project.

## Prerequisites

- Go 1.24 or higher
- Git
- curl (for testing)
- Docker (optional, for containerized deployment)

## Installation

### 1. Clone the Repository

```bash
git clone --recurse-submodules https://github.com/bakdaaswandi5818/poc-git-sub-module.git
cd poc-git-sub-module
```

If you forgot to use `--recurse-submodules`, run:

```bash
git submodule update --init --recursive
```

### 2. Install Dependencies

```bash
go mod download
go mod tidy
```

Or using make:

```bash
make install-deps
```

## Running the Application

### Option 1: Using Make (Recommended)

```bash
# Build and run
make run

# Or run in development mode (without building)
make dev
```

### Option 2: Using Go Commands

```bash
# Build
go build -o poc-server .

# Run
./poc-server
```

### Option 3: Using Docker

```bash
# Build and run with Docker Compose
docker-compose up --build

# Or build Docker image manually
docker build -t poc-git-submodule .
docker run -p 8080:8080 poc-git-submodule
```

## Testing the Application

Once the server is running on `http://localhost:8080`, test the endpoints:

### Welcome Endpoint
```bash
curl http://localhost:8080/
```

Expected output:
```json
{
  "message": "Git Submodule POC with Golang Echo Framework",
  "info": "Greeting Library v1.0.0 - Loaded from Git Submodule"
}
```

### Greet Endpoint
```bash
curl "http://localhost:8080/greet?name=Alice"
```

Expected output:
```json
{
  "greeting": "Hello, Alice! Welcome to the Git Submodule POC."
}
```

### Time-Based Greet Endpoint
```bash
curl "http://localhost:8080/greet/time-based?name=Bob"
```

Expected output (varies by time of day):
```json
{
  "greeting": "Good afternoon, Bob! This greeting comes from a Git Submodule."
}
```

### Health Check
```bash
curl http://localhost:8080/health
```

Expected output:
```json
{
  "status": "healthy",
  "submodule": "greeting-lib loaded successfully"
}
```

## Running Tests

```bash
# Run all tests
go test -v ./...

# Or using make
make test
```

## Project Structure

```
.
├── main.go                  # Main Echo application
├── main_test.go             # Application tests
├── go.mod                   # Go module definition
├── go.sum                   # Go dependencies checksum
├── .gitmodules              # Git submodules configuration
├── .gitignore               # Git ignore rules
├── Makefile                 # Build automation
├── Dockerfile               # Docker configuration
├── docker-compose.yml       # Docker Compose configuration
├── README.md                # Main documentation
├── QUICKSTART.md            # This file
└── libs/
    └── greeting-lib/        # Git submodule: Greeting library
        ├── greeting.go      # Library implementation
        ├── greeting_test.go # Library tests
        ├── go.mod           # Library module definition
        └── README.md        # Library documentation
```

## Common Commands

```bash
# Build the application
make build

# Run the application
make run

# Run in development mode
make dev

# Clean build artifacts
make clean

# Install dependencies
make install-deps

# Initialize submodules
make init-submodules

# Update submodules
make update-submodules

# Run tests
make test

# Format code
make fmt

# Show all available commands
make help
```

## Understanding Git Submodules

This project uses `libs/greeting-lib` as a Git submodule to demonstrate:

1. **Modularity**: Separating reusable libraries from the main application
2. **Version Control**: Tracking specific versions of dependencies
3. **Reusability**: Sharing code across multiple projects

### Key Concepts

- The `.gitmodules` file defines the submodule configuration
- The parent repository tracks a specific commit of the submodule
- Changes to the submodule must be committed separately
- Go's `replace` directive allows using local modules

## Troubleshooting

### Submodule Not Found

```bash
git submodule update --init --recursive
```

### Dependency Errors

```bash
go mod tidy
go mod download
```

### Port Already in Use

Change the port in `main.go` from `:8080` to another port:

```go
e.Logger.Fatal(e.Start(":3000"))
```

### Tests Failing

Make sure dependencies are up to date:

```bash
go mod tidy
go test -v ./...
```

## Next Steps

1. Explore the code in `main.go` to understand the Echo framework setup
2. Check out `libs/greeting-lib/greeting.go` to see the submodule implementation
3. Try modifying the endpoints or adding new functionality
4. Experiment with updating the submodule and see how the parent repository tracks changes

## Additional Resources

- [Full Documentation](README.md)
- [Echo Framework Docs](https://echo.labstack.com/)
- [Git Submodules Documentation](https://git-scm.com/book/en/v2/Git-Tools-Submodules)
- [Go Modules Reference](https://go.dev/ref/mod)
