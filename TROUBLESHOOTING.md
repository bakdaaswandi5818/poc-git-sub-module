# Troubleshooting Guide

This guide helps you resolve common issues when working with this Git Submodule POC.

## Table of Contents

- [Git Submodule Issues](#git-submodule-issues)
- [Build Issues](#build-issues)
- [Runtime Issues](#runtime-issues)
- [Test Issues](#test-issues)
- [Docker Issues](#docker-issues)
- [General Go Issues](#general-go-issues)

---

## Git Submodule Issues

### Issue: Submodule directory is empty after cloning

**Symptom:**
```bash
$ ls libs/greeting-lib/
# Empty directory
```

**Solution:**
```bash
git submodule update --init --recursive
```

**Prevention:**
Clone with `--recurse-submodules` flag:
```bash
git clone --recurse-submodules https://github.com/bakdaaswandi5818/poc-git-sub-module.git
```

---

### Issue: Submodule shows as modified but no files changed

**Symptom:**
```bash
$ git status
modified:   libs/greeting-lib (modified content)
```

**Solution:**
The submodule might be on a different commit. Reset it:
```bash
cd libs/greeting-lib
git status
# If you didn't intend changes:
git reset --hard HEAD
cd ../..
```

Or update to the commit tracked by parent:
```bash
git submodule update --force
```

---

### Issue: Can't pull latest changes from submodule

**Symptom:**
```bash
$ cd libs/greeting-lib
$ git pull
fatal: no remote repository specified
```

**Solution:**
The submodule might not have a remote set. Check .gitmodules:
```bash
cat .gitmodules
```

If it's a local path, you need to be in the submodule directory and commit there directly.

---

## Build Issues

### Issue: `go build` fails with module not found

**Symptom:**
```bash
$ go build
main.go:9:2: no required module provides package github.com/bakdaaswandi5818/greeting-lib
```

**Solution 1:** Run `go mod tidy`
```bash
go mod tidy
```

**Solution 2:** Ensure the replace directive exists in go.mod
```bash
grep "replace" go.mod
# Should show: replace github.com/bakdaaswandi5818/greeting-lib => ./libs/greeting-lib
```

**Solution 3:** Ensure submodule is initialized
```bash
git submodule update --init --recursive
```

---

### Issue: Build succeeds but binary is missing

**Symptom:**
```bash
$ make build
Building the application...
go build -o poc-server .
$ ls poc-server
ls: cannot access 'poc-server': No such file or directory
```

**Solution:**
Check if .gitignore is excluding it (which is intentional). The binary should be there:
```bash
ls -la | grep poc-server
```

If truly missing, try building directly:
```bash
go build -v -o poc-server .
```

---

### Issue: Build is very slow

**Symptom:**
First build takes a long time.

**Solution:**
This is normal for the first build as Go downloads dependencies. Subsequent builds are cached:
```bash
# First build - slow
go build -o poc-server .

# Subsequent builds - fast (uses cache)
go build -o poc-server .
```

To clean cache and rebuild:
```bash
go clean -cache
go build -o poc-server .
```

---

## Runtime Issues

### Issue: Server won't start - port already in use

**Symptom:**
```bash
$ ./poc-server
panic: listen tcp :8080: bind: address already in use
```

**Solution 1:** Stop the process using port 8080
```bash
# Find the process
lsof -i :8080
# Or on Linux
netstat -tlnp | grep 8080

# Kill it (replace PID with actual process ID)
kill <PID>
```

**Solution 2:** Change the port in main.go
```go
// Change from:
e.Logger.Fatal(e.Start(":8080"))

// To:
e.Logger.Fatal(e.Start(":3000"))
```

---

### Issue: Endpoints return 404

**Symptom:**
```bash
$ curl http://localhost:8080/greet
{"message":"Not Found"}
```

**Solution:**
Ensure the server is running and you're using the correct endpoint:
```bash
# Check if server is running
curl http://localhost:8080/health

# Available endpoints:
curl http://localhost:8080/
curl http://localhost:8080/greet?name=Alice
curl http://localhost:8080/greet/time-based?name=Bob
curl http://localhost:8080/health
```

---

### Issue: Server starts but crashes immediately

**Symptom:**
Server starts but exits with panic.

**Solution:**
Check the error message. Common causes:
1. Port already in use (see above)
2. Missing dependencies - run `go mod tidy`
3. Submodule not initialized - run `git submodule update --init`

Check logs:
```bash
./poc-server 2>&1 | tee server.log
```

---

## Test Issues

### Issue: Tests fail with "package not found"

**Symptom:**
```bash
$ go test ./...
main_test.go:9:2: no required module provides package github.com/stretchr/testify
```

**Solution:**
Install test dependencies:
```bash
go mod tidy
go mod download
```

---

### Issue: Tests pass individually but fail together

**Symptom:**
```bash
$ go test -v -run TestGreetEndpoint  # Passes
$ go test -v ./...                   # Fails
```

**Solution:**
Might be a caching issue:
```bash
go clean -testcache
go test -v ./...
```

---

### Issue: Submodule tests not running

**Symptom:**
Only seeing main app tests, not greeting-lib tests.

**Solution:**
Run tests from root to test all modules:
```bash
go test -v ./...
```

Or test submodule specifically:
```bash
cd libs/greeting-lib
go test -v .
```

---

## Docker Issues

### Issue: Docker build fails at submodule copy

**Symptom:**
```
COPY libs/greeting-lib ./libs/greeting-lib
COPY failed: file not found
```

**Solution:**
Ensure submodule is initialized before building:
```bash
git submodule update --init --recursive
docker build -t poc-server .
```

---

### Issue: Docker container exits immediately

**Symptom:**
```bash
$ docker run poc-server
# Container exits immediately
```

**Solution:**
Check logs:
```bash
docker run poc-server
# Or
docker logs <container-id>
```

Run interactively to debug:
```bash
docker run -it poc-server /bin/sh
```

---

### Issue: Can't access server in Docker

**Symptom:**
Server runs in Docker but can't access from host.

**Solution:**
Ensure you're mapping ports:
```bash
docker run -p 8080:8080 poc-server
```

Test from host:
```bash
curl http://localhost:8080/health
```

---

## General Go Issues

### Issue: `go mod tidy` changes go.mod unexpectedly

**Symptom:**
go.mod keeps changing after `go mod tidy`.

**Solution:**
This is normal. Go manages dependencies automatically. If changes are correct:
```bash
git add go.mod go.sum
git commit -m "Update dependencies"
```

---

### Issue: "go: inconsistent vendoring" error

**Symptom:**
```bash
go: inconsistent vendoring in /path/to/project
```

**Solution:**
If using vendoring:
```bash
go mod vendor
```

If not using vendoring (this project doesn't):
```bash
rm -rf vendor/
go mod tidy
```

---

### Issue: Import path errors

**Symptom:**
```go
import "github.com/bakdaaswandi5818/greeting-lib" 
// Error: package not found
```

**Solution:**
Verify the replace directive in go.mod:
```bash
cat go.mod | grep replace
```

Should show:
```
replace github.com/bakdaaswandi5818/greeting-lib => ./libs/greeting-lib
```

---

## Getting Help

If you encounter an issue not covered here:

1. **Check Documentation:**
   - README.md
   - QUICKSTART.md
   - GIT_SUBMODULES_GUIDE.md

2. **Verify Environment:**
   ```bash
   go version          # Should be 1.24+
   git --version       # Should be 2.0+
   docker --version    # If using Docker
   ```

3. **Clean and Rebuild:**
   ```bash
   make clean
   go clean -cache -testcache -modcache
   git submodule update --init --recursive
   go mod tidy
   make build
   make test
   ```

4. **Check for Common Issues:**
   - Submodule not initialized
   - Wrong directory
   - Port conflicts
   - Firewall blocking
   - Antivirus interference

5. **Enable Verbose Logging:**
   ```bash
   go build -v -o poc-server .
   go test -v ./...
   ```

6. **Still Stuck?**
   - Check git status: `git status`
   - Check submodule status: `git submodule status`
   - Verify Go modules: `go list -m all`
   - Check for file conflicts: `git diff`

---

## Quick Fixes Checklist

When something breaks, try these in order:

- [ ] `git submodule update --init --recursive`
- [ ] `go mod tidy`
- [ ] `make clean && make build`
- [ ] `go clean -testcache && make test`
- [ ] Check if port 8080 is free
- [ ] Verify Go version (1.24+)
- [ ] Read error message carefully
- [ ] Check documentation

---

## Common Error Messages

| Error | Solution |
|-------|----------|
| `no required module provides package` | Run `go mod tidy` |
| `address already in use` | Change port or kill process |
| `submodule directory is empty` | Run `git submodule update --init` |
| `file not found` | Ensure you're in project root |
| `permission denied` | Check file permissions or use sudo |
| `GOPROXY error` | Check internet connection |

---

## Prevention Tips

1. **Always clone with submodules:**
   ```bash
   git clone --recurse-submodules <url>
   ```

2. **Run tests before committing:**
   ```bash
   make test
   ```

3. **Keep dependencies updated:**
   ```bash
   go get -u ./...
   go mod tidy
   ```

4. **Use the Makefile:**
   ```bash
   make help  # See all available commands
   ```

5. **Read error messages carefully** - they usually tell you what's wrong

6. **Keep documentation handy** - refer to README.md and other docs

---

For more information, see:
- [Git Submodules Guide](GIT_SUBMODULES_GUIDE.md)
- [Quick Start Guide](QUICKSTART.md)
- [README](README.md)
