# Architecture Diagram

## System Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                     Client Applications                       │
│  (curl, Browser, Mobile App, Postman, etc.)                  │
└────────────────────────┬────────────────────────────────────┘
                         │ HTTP/REST
                         ▼
┌─────────────────────────────────────────────────────────────┐
│                    Echo Web Server                            │
│                     (Port: 8080)                              │
│  ┌──────────────────────────────────────────────────────┐   │
│  │              Middleware Stack                         │   │
│  │  • Logger                                             │   │
│  │  • Recover                                            │   │
│  │  • CORS                                               │   │
│  └──────────────────────────────────────────────────────┘   │
│                                                               │
│  ┌──────────────────────────────────────────────────────┐   │
│  │                Route Handlers                         │   │
│  │                                                       │   │
│  │  GET /                  → Welcome Handler            │   │
│  │  GET /greet             → Greet Handler              │   │
│  │  GET /greet/time-based  → Time-based Greet Handler   │   │
│  │  GET /health            → Health Check Handler       │   │
│  └──────────────────────────────────────────────────────┘   │
└────────────────────────┬────────────────────────────────────┘
                         │ Import
                         ▼
┌─────────────────────────────────────────────────────────────┐
│              Greeting Library (Git Submodule)                 │
│                    libs/greeting-lib/                         │
│  ┌──────────────────────────────────────────────────────┐   │
│  │               Greeter Struct                          │   │
│  │  • defaultName: string                                │   │
│  └──────────────────────────────────────────────────────┘   │
│                                                               │
│  ┌──────────────────────────────────────────────────────┐   │
│  │              Public Methods                           │   │
│  │  • NewGreeter(name) → Greeter                         │   │
│  │  • Greet(name) → string                               │   │
│  │  • GetTimeBasedGreeting(name) → string                │   │
│  │  • GetInfo() → string                                 │   │
│  └──────────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────┘
```

## Git Repository Structure

```
┌─────────────────────────────────────────────────────────────┐
│        Main Repository (poc-git-sub-module)                   │
│  ┌────────────────────────────────────────────────────┐     │
│  │  Application Files                                  │     │
│  │  • main.go                                          │     │
│  │  • main_test.go                                     │     │
│  │  • go.mod (with replace directive)                 │     │
│  │  • go.sum                                           │     │
│  │  • Makefile                                         │     │
│  │  • Dockerfile                                       │     │
│  │  • docker-compose.yml                               │     │
│  │  • Documentation files                              │     │
│  └────────────────────────────────────────────────────┘     │
│                                                               │
│  ┌────────────────────────────────────────────────────┐     │
│  │  .gitmodules (submodule configuration)              │     │
│  │  [submodule "libs/greeting-lib"]                    │     │
│  │      path = libs/greeting-lib                       │     │
│  │      url = ./libs/greeting-lib                      │     │
│  └────────────────────────────────────────────────────┘     │
│                          │                                    │
│                          │ Git Submodule Reference           │
│                          ▼                                    │
│  ┌────────────────────────────────────────────────────┐     │
│  │    Submodule Repository (greeting-lib)              │     │
│  │  ┌──────────────────────────────────────────────┐  │     │
│  │  │  • greeting.go                                │  │     │
│  │  │  • greeting_test.go                           │  │     │
│  │  │  • go.mod                                     │  │     │
│  │  │  • README.md                                  │  │     │
│  │  │  • .git/ (separate git repository)            │  │     │
│  │  └──────────────────────────────────────────────┘  │     │
│  └────────────────────────────────────────────────────┘     │
└─────────────────────────────────────────────────────────────┘
```

## Request Flow Diagram

```
Client Request → Echo Router → Middleware Stack → Handler
                                                      │
                                                      ├→ Create Greeter
                                                      │   (from submodule)
                                                      │
                                                      ├→ Call Method
                                                      │   (Greet, GetTimeBasedGreeting, etc.)
                                                      │
                                                      └→ Return JSON Response
```

## Example: Greet Endpoint Flow

```
1. Client:
   curl "http://localhost:8080/greet?name=Alice"
   
2. Echo Server:
   - Logger middleware logs request
   - CORS middleware sets headers
   - Router matches GET /greet
   
3. Greet Handler:
   - Extract query param "name=Alice"
   - Call greeter.Greet("Alice")
   
4. Greeting Library (Submodule):
   - Receive name "Alice"
   - Generate message: "Hello, Alice! Welcome to the Git Submodule POC."
   - Return string
   
5. Handler:
   - Wrap in JSON: {"greeting": "Hello, Alice! ..."}
   - Send response with 200 OK
   
6. Client:
   - Receive JSON response
```

## Development Workflow

```
┌──────────────────────────────────────────────────────────────┐
│                  Developer Workflow                           │
└──────────────────────────────────────────────────────────────┘

1. Clone Repository
   ↓
   git clone --recurse-submodules <url>
   
2. Make Changes
   ↓
   ┌─────────────────────┐          ┌──────────────────────┐
   │  Change Submodule   │    OR    │  Change Main App     │
   │  libs/greeting-lib/ │          │  main.go             │
   └─────────────────────┘          └──────────────────────┘
   
3. Test Changes
   ↓
   make test
   
4. Build Application
   ↓
   make build
   
5. Run Application
   ↓
   make run
   
6. Commit Changes
   ↓
   ┌─────────────────────┐          ┌──────────────────────┐
   │  If Submodule:      │          │  If Main App:        │
   │  cd libs/greeting-lib│          │  git add .          │
   │  git add .          │          │  git commit -m "..." │
   │  git commit -m "..." │          │  git push           │
   │  cd ../..           │          └──────────────────────┘
   │  git add libs/...   │
   │  git commit -m "..." │
   │  git push           │
   └─────────────────────┘
```

## Deployment Options

```
┌──────────────────────────────────────────────────────────────┐
│                    Deployment Options                         │
└──────────────────────────────────────────────────────────────┘

Option 1: Direct Binary
   ├─ make build
   └─ ./poc-server

Option 2: Go Run
   └─ go run main.go

Option 3: Docker
   ├─ docker build -t poc-server .
   └─ docker run -p 8080:8080 poc-server

Option 4: Docker Compose
   └─ docker-compose up

Option 5: Cloud Deployment
   ├─ Build binary
   ├─ Create container
   └─ Deploy to:
      ├─ AWS (ECS, EKS, Lambda)
      ├─ Google Cloud (Cloud Run, GKE)
      ├─ Azure (Container Instances, AKS)
      └─ Heroku, DigitalOcean, etc.
```

## Technology Stack

```
┌─────────────────────────────────────────────────────────────┐
│                    Technology Stack                          │
└─────────────────────────────────────────────────────────────┘

Language:
   └─ Go 1.24

Web Framework:
   └─ Echo v4.15.0

Middleware:
   ├─ Logger
   ├─ Recover
   └─ CORS

Dependency Management:
   ├─ Go Modules
   └─ Git Submodules

Build Tools:
   ├─ Go toolchain
   ├─ Make
   └─ Docker

Testing:
   ├─ testing package
   └─ testify/assert

Version Control:
   └─ Git (with submodules)

Documentation:
   └─ Markdown
```

## Security & Quality

```
┌─────────────────────────────────────────────────────────────┐
│                  Security & Quality Checks                   │
└─────────────────────────────────────────────────────────────┘

Code Quality:
   ├─ go fmt (formatting)
   ├─ go vet (static analysis)
   └─ Code review ✓

Security:
   ├─ CodeQL analysis ✓
   └─ No vulnerabilities found ✓

Testing:
   ├─ Unit tests (8/8 passing) ✓
   ├─ Main app tests (4/4) ✓
   └─ Library tests (4/4) ✓

Build:
   ├─ Successful build ✓
   └─ Server runs correctly ✓
```

## Key Benefits

```
┌─────────────────────────────────────────────────────────────┐
│                      Key Benefits                            │
└─────────────────────────────────────────────────────────────┘

Modularity
   └─ Separate concerns with submodules

Reusability
   └─ Greeting library can be used in other projects

Version Control
   └─ Track specific library versions

Maintainability
   └─ Clear separation between app and library

Scalability
   └─ Easy to add more submodules/endpoints

Documentation
   └─ Comprehensive guides for all aspects

Developer Experience
   └─ Simple commands (make build, make run, etc.)
```

---

For more details, see the individual documentation files:
- README.md - Project overview
- QUICKSTART.md - Getting started
- API_DOCUMENTATION.md - API reference
- GIT_SUBMODULES_GUIDE.md - Submodule details
- EXTENDING.md - Extension examples
