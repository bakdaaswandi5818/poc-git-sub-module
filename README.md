# Git Submodule POC - Full Stack with Golang Echo & React

This project demonstrates how to use Git submodules in both backend (Go/Echo) and frontend (React) applications, showcasing a complete modular architecture.

## Overview

This proof of concept shows:
- How to structure a **full-stack project** with **multiple Git submodules**
- **Backend**: Go Echo framework with 2 library submodules
- **Frontend**: React application with 2 library submodules
- Best practices for managing dependencies with Go modules and Git submodules
- Seamless integration between frontend and backend

## Project Structure

```
.
├── main.go                     # Main Echo application (Backend)
├── go.mod                      # Go module definition
├── libs/                       # Backend submodules
│   ├── greeting-lib/           # Git submodule: Greeting library
│   │   ├── greeting.go
│   │   ├── go.mod
│   │   └── README.md
│   └── logger-lib/             # Git submodule: Logger library
│       ├── logger.go
│       ├── go.mod
│       └── README.md
├── frontend/                   # Frontend application
│   ├── components-lib/         # Git submodule: React UI components
│   │   ├── src/index.js
│   │   ├── package.json
│   │   └── README.md
│   ├── utils-lib/              # Git submodule: Utility functions
│   │   ├── src/index.js
│   │   ├── package.json
│   │   └── README.md
│   ├── src/                    # Main React app
│   │   ├── App.jsx
│   │   ├── main.jsx
│   │   └── index.css
│   ├── package.json
│   └── vite.config.js
└── README.md                   # This file
```

## Features

### Backend Submodules (Go)

This project demonstrates **two backend Git submodules**:

#### 1. Greeting Library (`libs/greeting-lib`)
A standalone Go module that provides:
- Basic greeting functionality
- Time-based greetings (morning, afternoon, evening)
- Customizable default name
- Version information

#### 2. Logger Library (`libs/logger-lib`)
A standalone logging module that provides:
- Structured logging with timestamps
- Multiple log levels (Info, Warn, Error, Debug)
- Customizable prefix for different components
- Lightweight and simple to use

### Frontend Submodules (React)

This project demonstrates **two frontend Git submodules**:

#### 3. Components Library (`frontend/components-lib`)
A reusable React UI components library with:
- Button (4 variants)
- Card (with header/footer)
- Alert (4 types)
- Input (with validation)
- Spinner (3 sizes)

#### 4. Utils Library (`frontend/utils-lib`)
A comprehensive utility functions library with:
- Date formatting (formatDate, getTimeAgo)
- String utilities (capitalize, truncate, slugify)
- Validation (isEmail, isURL, isPhoneNumber)
- Number formatting (formatCurrency, formatNumber)
- Array utilities (groupBy, sortBy, unique, chunk)
- Storage wrapper, debounce, throttle, and more

### Backend API (Echo Framework)
The Go backend provides the following endpoints:

- `GET /` - Welcome message with both backend submodules' info
- `GET /greet?name=John` - Get a personalized greeting
- `GET /greet/time-based?name=John` - Get a time-based greeting
- `GET /health` - Health check endpoint showing both submodules loaded

### Frontend Application (React + Vite)
The React frontend demonstrates:
- Integration with both frontend submodules
- Real-time communication with backend API
- Component library usage
- Utility functions usage
- Responsive design
- Modern development setup with Vite

## Setup and Installation

### Prerequisites
- Go 1.24 or higher
- Node.js 16+ and npm
- Git

### Clone with Submodules

```bash
# Clone the repository with ALL submodules (RECOMMENDED)
git clone --recurse-submodules https://github.com/bakdaaswandi5818/poc-git-sub-module.git

# Or if you already cloned without submodules
git clone https://github.com/bakdaaswandi5818/poc-git-sub-module.git
cd poc-git-sub-module
git submodule update --init --recursive
```

> **📘 For detailed instructions on pulling and updating submodules, see [SUBMODULE_PULL_GUIDE.md](SUBMODULE_PULL_GUIDE.md)**

### Install Dependencies

**Backend:**
```bash
# Install Go dependencies
go mod download
go mod tidy
```

**Frontend:**
```bash
# Install Node.js dependencies
cd frontend
npm install
cd ..
```

### Running the Full Stack Application

**Option 1: Run Both Services (Recommended for Full Experience)**

Terminal 1 - Backend:
```bash
# Build and run the Go server
go build -o poc-server .
./poc-server
```

Terminal 2 - Frontend:
```bash
# Start the React development server
cd frontend
npm run dev
```

- Backend API: `http://localhost:8080`
- Frontend App: `http://localhost:3000`

**Option 2: Backend Only**

```bash
# Run the backend server
go run main.go
```

The server will start on `http://localhost:8080`

**Option 3: Frontend Only (with Backend Running)**

```bash
cd frontend
npm run dev
```

The frontend will be available at `http://localhost:3000` and will proxy API requests to `http://localhost:8080`

### Build for Production

**Backend:**
```bash
go build -o poc-server .
```

**Frontend:**
```bash
cd frontend
npm run build
# Production files will be in frontend/dist/
```

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