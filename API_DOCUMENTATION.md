# API Documentation

## Base URL
```
http://localhost:8080
```

## Endpoints

### 1. Welcome / Root Endpoint

Get welcome message and submodule information.

**Endpoint:** `GET /`

**Request:**
```bash
curl http://localhost:8080/
```

**Response:**
```json
{
  "message": "Git Submodule POC with Golang Echo Framework",
  "info": "Greeting Library v1.0.0 - Loaded from Git Submodule"
}
```

**Status Codes:**
- `200 OK`: Success

---

### 2. Greet Endpoint

Get a personalized greeting message.

**Endpoint:** `GET /greet`

**Query Parameters:**
- `name` (optional): The name to greet. If not provided, uses default name "Guest"

**Request:**
```bash
# With name
curl "http://localhost:8080/greet?name=Alice"

# Without name (uses default)
curl http://localhost:8080/greet
```

**Response:**
```json
{
  "greeting": "Hello, Alice! Welcome to the Git Submodule POC."
}
```

**Status Codes:**
- `200 OK`: Success

---

### 3. Time-Based Greet Endpoint

Get a greeting message based on the current time of day.

**Endpoint:** `GET /greet/time-based`

**Query Parameters:**
- `name` (optional): The name to greet. If not provided, uses default name "Guest"

**Request:**
```bash
# With name
curl "http://localhost:8080/greet/time-based?name=Bob"

# Without name (uses default)
curl http://localhost:8080/greet/time-based
```

**Response:**
```json
{
  "greeting": "Good afternoon, Bob! This greeting comes from a Git Submodule."
}
```

The greeting prefix varies by time:
- **Before 12:00**: "Good morning"
- **12:00 - 17:59**: "Good afternoon"
- **18:00 and later**: "Good evening"

**Status Codes:**
- `200 OK`: Success

---

### 4. Health Check Endpoint

Check the health status of the application and verify submodule is loaded.

**Endpoint:** `GET /health`

**Request:**
```bash
curl http://localhost:8080/health
```

**Response:**
```json
{
  "status": "healthy",
  "submodule": "greeting-lib loaded successfully"
}
```

**Status Codes:**
- `200 OK`: Service is healthy

---

## Request Examples

### Using curl

```bash
# Basic greeting
curl http://localhost:8080/greet?name=John

# Time-based greeting
curl http://localhost:8080/greet/time-based?name=Jane

# Health check
curl http://localhost:8080/health
```

### Using HTTPie

```bash
# Basic greeting
http GET localhost:8080/greet name==Alice

# Time-based greeting
http GET localhost:8080/greet/time-based name==Bob

# Health check
http GET localhost:8080/health
```

### Using JavaScript (Fetch API)

```javascript
// Basic greeting
fetch('http://localhost:8080/greet?name=Alice')
  .then(response => response.json())
  .then(data => console.log(data.greeting));

// Time-based greeting
fetch('http://localhost:8080/greet/time-based?name=Bob')
  .then(response => response.json())
  .then(data => console.log(data.greeting));

// Health check
fetch('http://localhost:8080/health')
  .then(response => response.json())
  .then(data => console.log(data));
```

### Using Python (requests)

```python
import requests

# Basic greeting
response = requests.get('http://localhost:8080/greet', params={'name': 'Alice'})
print(response.json()['greeting'])

# Time-based greeting
response = requests.get('http://localhost:8080/greet/time-based', params={'name': 'Bob'})
print(response.json()['greeting'])

# Health check
response = requests.get('http://localhost:8080/health')
print(response.json())
```

## Response Format

All endpoints return JSON responses with appropriate HTTP status codes.

### Success Response
All successful responses return `200 OK` with a JSON body.

### Error Responses (Future Enhancement)
Currently, the application uses Echo's default error handling. In a production application, you might want to add:

```json
{
  "error": "Error message",
  "code": "ERROR_CODE",
  "details": {}
}
```

## CORS

The application has CORS enabled via Echo's CORS middleware, allowing cross-origin requests from any domain.

## Rate Limiting

Currently, there is no rate limiting. For production use, consider adding rate limiting middleware.

## Authentication

This POC does not include authentication. For production use, consider adding:
- API keys
- JWT tokens
- OAuth2

## Middleware

The application uses the following Echo middleware:
- **Logger**: Logs all HTTP requests
- **Recover**: Recovers from panics
- **CORS**: Enables cross-origin resource sharing

## Testing

You can use the included test suite to verify the API:

```bash
go test -v ./...
```

Or test manually with the provided examples above.
