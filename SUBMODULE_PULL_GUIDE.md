# How to Pull and Update Git Submodules

This guide explains how to clone, pull, and update Git submodules in this project.

## Understanding Submodules

This project uses **two Git submodules**:
1. `libs/greeting-lib` - Greeting functionality library
2. `libs/logger-lib` - Structured logging library

Each submodule is a separate Git repository tracked at a specific commit by the parent repository.

---

## Initial Clone with Submodules

When cloning this repository for the first time, use one of these methods:

### Method 1: Clone with --recurse-submodules (Recommended)

```bash
git clone --recurse-submodules https://github.com/bakdaaswandi5818/poc-git-sub-module.git
cd poc-git-sub-module
```

This automatically initializes and updates all submodules during the clone.

### Method 2: Clone then Initialize Submodules

```bash
# Clone the repository
git clone https://github.com/bakdaaswandi5818/poc-git-sub-module.git
cd poc-git-sub-module

# Initialize and update all submodules
git submodule update --init --recursive
```

The `--init` flag initializes submodules, and `--recursive` handles nested submodules (if any).

---

## Checking Submodule Status

To see the current state of submodules:

```bash
# View basic status
git submodule status

# Example output:
# 522178c libs/greeting-lib (heads/master)
# 0bfaf62 libs/logger-lib (heads/master)
```

The hash before each path is the commit that the parent repository tracks.

To see if submodules have uncommitted changes:

```bash
git status
```

---

## Pulling Updates from Submodules

### Update All Submodules to Latest Commits

When someone pushes new changes to the submodule repositories, update them:

```bash
# Update all submodules to their latest commit on tracked branch
git submodule update --remote --merge

# Or update and rebase instead of merge
git submodule update --remote --rebase
```

The `--remote` flag fetches the latest changes from the submodule's remote repository.

### Update Specific Submodule

To update just one submodule:

```bash
# Update only greeting-lib
git submodule update --remote --merge libs/greeting-lib

# Update only logger-lib  
git submodule update --remote --merge libs/logger-lib
```

### Update Submodules to Commit Tracked by Parent

If you pull changes in the parent repository and the submodule references have changed:

```bash
# After git pull in parent repository
git submodule update --init --recursive
```

This ensures submodules match the exact commits tracked by the parent.

---

## Making Changes in Submodules

### Working Inside a Submodule

When you need to modify code in a submodule:

```bash
# Navigate to the submodule
cd libs/greeting-lib

# Check the branch (submodules start in "detached HEAD" state)
git branch

# Checkout a branch to make changes
git checkout master  # or main, depending on the repo

# Make your changes
# Edit files...

# Commit changes in the submodule
git add .
git commit -m "Update greeting functionality"

# Push changes to the submodule repository
git push origin master
```

### Updating Parent Repository

After committing changes in a submodule, update the parent repository:

```bash
# Go back to parent repository
cd ../..

# The parent will show the submodule as modified
git status
# Output: modified:   libs/greeting-lib (new commits)

# Stage the submodule update
git add libs/greeting-lib

# Commit the updated reference
git commit -m "Update greeting-lib to latest version"

# Push to parent repository
git push
```

---

## Common Scenarios

### Scenario 1: Fresh Clone and Setup

```bash
# Clone with submodules
git clone --recurse-submodules https://github.com/bakdaaswandi5818/poc-git-sub-module.git
cd poc-git-sub-module

# Verify submodules are initialized
git submodule status

# Install Go dependencies
go mod tidy

# Build the project
go build -o poc-server .

# Run tests
go test ./...
```

### Scenario 2: Someone Updated a Submodule

```bash
# In parent repository
git pull

# Output might show: Fetching submodule libs/greeting-lib

# Update submodules to new tracked commits
git submodule update --init --recursive

# Verify the update
git submodule status

# Rebuild if needed
go mod tidy
go build -o poc-server .
```

### Scenario 3: Update to Latest Submodule Changes

```bash
# Update all submodules to their latest remote commits
git submodule update --remote --merge

# Check what changed
git diff libs/greeting-lib
git diff libs/logger-lib

# Test the changes
go test ./...

# If everything works, commit the updates
git add libs/greeting-lib libs/logger-lib
git commit -m "Update submodules to latest versions"
git push
```

### Scenario 4: Submodule Shows as Modified but You Didn't Change It

This happens when the submodule is on a different commit than tracked:

```bash
# Reset submodule to tracked commit
git submodule update --force

# Or if you want to keep local changes
cd libs/greeting-lib
git checkout master
git pull
cd ../..
git add libs/greeting-lib
git commit -m "Update greeting-lib reference"
```

---

## Advanced Operations

### Update and Initialize New Submodules

If a new submodule is added to the project:

```bash
git pull
git submodule update --init --recursive
```

### Execute Command in All Submodules

```bash
# Run git status in all submodules
git submodule foreach 'git status'

# Run git pull in all submodules
git submodule foreach 'git pull origin master'

# Check which branch each submodule is on
git submodule foreach 'git branch'
```

### Synchronize Submodule URLs

If submodule URLs change in `.gitmodules`:

```bash
git submodule sync --recursive
git submodule update --init --recursive
```

---

## Troubleshooting

### Submodule Directory is Empty

```bash
git submodule update --init --recursive
```

### Submodule is in Detached HEAD State

```bash
cd libs/greeting-lib
git checkout master  # or the appropriate branch
cd ../..
```

### Can't Pull Latest Submodule Changes

```bash
cd libs/greeting-lib
git fetch origin
git merge origin/master  # or rebase
cd ../..
git add libs/greeting-lib
git commit -m "Update greeting-lib"
```

### Submodule Changes Won't Stage

Make sure you're in the parent directory, not inside the submodule:

```bash
cd /home/runner/work/poc-git-sub-module/poc-git-sub-module
git add libs/greeting-lib
git status
```

---

## Best Practices

1. **Always Use `--recurse-submodules` When Cloning**
   ```bash
   git clone --recurse-submodules <url>
   ```

2. **Update Submodules After Pulling Parent**
   ```bash
   git pull
   git submodule update --init --recursive
   ```

3. **Check Submodule Status Regularly**
   ```bash
   git submodule status
   git status
   ```

4. **Work on Branches in Submodules**
   - Don't work in detached HEAD state
   - Always checkout a branch before making changes

5. **Commit Submodule Changes First**
   - Commit and push submodule changes before updating parent
   - Then update the parent's submodule reference

6. **Test After Submodule Updates**
   ```bash
   go mod tidy
   go test ./...
   go build -o poc-server .
   ```

---

## Quick Reference

| Task | Command |
|------|---------|
| Clone with submodules | `git clone --recurse-submodules <url>` |
| Initialize submodules | `git submodule update --init --recursive` |
| Update to tracked commits | `git submodule update --init --recursive` |
| Update to latest remote | `git submodule update --remote --merge` |
| Check submodule status | `git submodule status` |
| Run command in all submodules | `git submodule foreach '<command>'` |
| Update specific submodule | `git submodule update --remote libs/greeting-lib` |

---

## Workflow Summary

**For Regular Development:**
1. `git pull` (parent repository)
2. `git submodule update --init --recursive` (sync submodules)
3. Make changes
4. Test: `go test ./...`
5. Commit and push

**For Updating Submodules:**
1. `cd libs/greeting-lib`
2. `git checkout master`
3. Make changes
4. `git commit -m "message"`
5. `git push origin master`
6. `cd ../..`
7. `git add libs/greeting-lib`
8. `git commit -m "Update greeting-lib"`
9. `git push`

---

## Additional Resources

- [Git Submodules Official Documentation](https://git-scm.com/book/en/v2/Git-Tools-Submodules)
- [GitHub Guide on Submodules](https://github.blog/2016-02-01-working-with-submodules/)
- [Pro Git Book - Submodules](https://git-scm.com/book/en/v2/Git-Tools-Submodules)

For more information, see also:
- [GIT_SUBMODULES_GUIDE.md](GIT_SUBMODULES_GUIDE.md) - In-depth submodule concepts
- [TROUBLESHOOTING.md](TROUBLESHOOTING.md) - Common issues and solutions
- [README.md](README.md) - Project overview
