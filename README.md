# Git Submodule POC with Golang Echo Framework

This project demonstrates how to use Git submodules in a Go application built with the Echo web framework.

## Overview

This proof of concept shows:
- How to structure a Go project with Git submodules
- Using a separate Go module as a Git submodule
- Building a REST API with Echo framework that uses the submodule library
- Best practices for managing dependencies with Go modules and Git submodules

## Project Structure

```
.
├── main.go                  # Main Echo application
├── go.mod                   # Go module definition with submodule replacement
├── go.sum                   # Go dependencies checksum
├── .gitmodules              # Git submodules configuration
├── libs/
│   └── greeting-lib/        # Git submodule: Greeting library
│       ├── greeting.go      # Library implementation
│       ├── go.mod           # Library module definition
│       └── README.md        # Library documentation
└── README.md                # This file
```

## Features

### Greeting Library (Submodule)
The `greeting-lib` is a standalone Go module that provides:
- Basic greeting functionality
- Time-based greetings (morning, afternoon, evening)
- Customizable default name
- Version information

### Echo REST API
The main application provides the following endpoints:

- `GET /` - Welcome message with submodule info
- `GET /greet?name=John` - Get a personalized greeting
- `GET /greet/time-based?name=John` - Get a time-based greeting
- `GET /health` - Health check endpoint

## Setup and Installation

### Prerequisites
- Go 1.24 or higher
- Git

### Clone with Submodules

```bash
# Clone the repository with submodules
git clone --recurse-submodules https://github.com/bakdaaswandi5818/poc-git-sub-module.git

# Or if you already cloned without submodules
git clone https://github.com/bakdaaswandi5818/poc-git-sub-module.git
cd poc-git-sub-module
git submodule update --init --recursive
```

### Install Dependencies

```bash
# Install Go dependencies
go mod download
go mod tidy
```

### Build the Application

```bash
# Build the server
go build -o poc-server .
```

### Run the Application

```bash
# Run the server
./poc-server

# Or run directly with Go
go run main.go
```

The server will start on `http://localhost:8080`

## API Usage Examples

### Welcome Endpoint
```bash
curl http://localhost:8080/
```

Response:
```json
{
  "message": "Git Submodule POC with Golang Echo Framework",
  "info": "Greeting Library v1.0.0 - Loaded from Git Submodule"
}
```

### Basic Greeting
```bash
curl http://localhost:8080/greet?name=Alice
```

Response:
```json
{
  "greeting": "Hello, Alice! Welcome to the Git Submodule POC."
}
```

### Time-Based Greeting
```bash
curl http://localhost:8080/greet/time-based?name=Bob
```

Response:
```json
{
  "greeting": "Good afternoon, Bob! This greeting comes from a Git Submodule."
}
```

### Health Check
```bash
curl http://localhost:8080/health
```

Response:
```json
{
  "status": "healthy",
  "submodule": "greeting-lib loaded successfully"
}
```

## Git Submodule Management

### Working with Submodules

```bash
# Initialize submodules
git submodule update --init --recursive

# Update submodule to latest commit
cd libs/greeting-lib
git pull origin master
cd ../..
git add libs/greeting-lib
git commit -m "Update greeting-lib submodule"

# View submodule status
git submodule status

# Remove a submodule (if needed)
git submodule deinit libs/greeting-lib
git rm libs/greeting-lib
rm -rf .git/modules/libs/greeting-lib
```

### How It Works

1. **Submodule Configuration**: The `.gitmodules` file defines the submodule path and URL
2. **Go Module Replacement**: The `go.mod` file uses a `replace` directive to use the local submodule
3. **Import Path**: The main application imports the library using its module name
4. **Version Control**: Git tracks the specific commit of the submodule

## Key Concepts

### Git Submodules
- Git submodules allow you to keep a Git repository as a subdirectory of another Git repository
- The parent repository tracks the specific commit of the submodule
- Changes to the submodule must be committed separately from the parent repository

### Go Module Replacement
The `go.mod` file uses a `replace` directive to use the local submodule:

```go
replace github.com/bakdaaswandi5818/greeting-lib => ./libs/greeting-lib
```

This tells Go to use the local directory instead of fetching from a remote repository.

## Development Workflow

1. Make changes to the submodule in `libs/greeting-lib/`
2. Commit changes in the submodule:
   ```bash
   cd libs/greeting-lib
   git add .
   git commit -m "Update greeting functionality"
   cd ../..
   ```
3. Update the parent repository to track the new submodule commit:
   ```bash
   git add libs/greeting-lib
   git commit -m "Update greeting-lib submodule reference"
   ```
4. Run `go mod tidy` to update dependencies
5. Build and test the application

## Benefits of Using Git Submodules

1. **Separation of Concerns**: Keep libraries and main code in separate repositories
2. **Version Control**: Track specific versions of dependencies
3. **Reusability**: Share libraries across multiple projects
4. **Independent Development**: Submodules can be developed and versioned independently

## Troubleshooting

### Submodule not initialized
```bash
git submodule update --init --recursive
```

### Submodule shows modified changes
```bash
cd libs/greeting-lib
git status
# If unintended changes, reset to tracked commit
git reset --hard
cd ../..
```

### Go module errors
```bash
go mod tidy
go mod verify
```

## License

This is a proof of concept project for demonstration purposes.

## Contributing

This is a demonstration project. Feel free to fork and experiment with different configurations.

## References

- [Git Submodules Documentation](https://git-scm.com/book/en/v2/Git-Tools-Submodules)
- [Echo Framework Documentation](https://echo.labstack.com/)
- [Go Modules Reference](https://go.dev/ref/mod)