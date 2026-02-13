# Understanding Git Submodules in This Project

This document explains how Git submodules are used in this proof of concept and how they work.

## What are Git Submodules?

Git submodules allow you to keep a Git repository as a subdirectory of another Git repository. This lets you clone another repository into your project and keep your commits separate.

## How This Project Uses Submodules

In this POC, we have:

1. **Main Repository**: `poc-git-sub-module` - The Echo framework application
2. **Submodule**: `libs/greeting-lib` - A standalone greeting library

## Configuration Files

### .gitmodules

The `.gitmodules` file defines all submodules:

```ini
[submodule "libs/greeting-lib"]
	path = libs/greeting-lib
	url = ./libs/greeting-lib
```

This tells Git:
- The submodule name: `libs/greeting-lib`
- Where to find it in this repo: `libs/greeting-lib`
- Where to clone it from: `./libs/greeting-lib` (local path for demo purposes)

### go.mod

The Go module file uses a `replace` directive to use the local submodule:

```go
replace github.com/bakdaaswandi5818/greeting-lib => ./libs/greeting-lib
```

This tells Go to use the local directory instead of trying to fetch from a remote source.

## Working with Submodules

### Cloning a Repository with Submodules

When cloning, use `--recurse-submodules`:

```bash
git clone --recurse-submodules https://github.com/bakdaaswandi5818/poc-git-sub-module.git
```

Or initialize after cloning:

```bash
git clone https://github.com/bakdaaswandi5818/poc-git-sub-module.git
cd poc-git-sub-module
git submodule update --init --recursive
```

### Checking Submodule Status

```bash
# View status of all submodules
git submodule status

# View detailed information
git submodule foreach git status
```

### Making Changes to a Submodule

1. Navigate to the submodule directory:
```bash
cd libs/greeting-lib
```

2. Make your changes and commit them:
```bash
# Make changes to greeting.go
git add greeting.go
git commit -m "Update greeting functionality"
```

3. Go back to the main repository and commit the submodule update:
```bash
cd ../..
git add libs/greeting-lib
git commit -m "Update greeting-lib submodule"
```

### Updating Submodules

To update a submodule to the latest commit from its repository:

```bash
# Update specific submodule
cd libs/greeting-lib
git pull origin master
cd ../..
git add libs/greeting-lib
git commit -m "Update greeting-lib to latest version"

# Or update all submodules from main repo
git submodule update --remote --merge
```

### Viewing Submodule Changes

```bash
# See what commit the submodule is at
git submodule status

# See changes in the submodule
git diff libs/greeting-lib

# See the actual changes in the submodule files
cd libs/greeting-lib
git diff
```

## Important Concepts

### Detached HEAD State

When you initialize a submodule, Git checks out a specific commit, leaving the submodule in a "detached HEAD" state. This means you're not on any branch.

To fix this:
```bash
cd libs/greeting-lib
git checkout master  # or main, or whatever branch you want
```

### Submodule Commit Tracking

The parent repository doesn't track the submodule's files directly. Instead, it tracks the specific commit SHA of the submodule. When you update a submodule, you're changing which commit the parent points to.

### Two-Step Commit Process

1. First, commit changes inside the submodule
2. Then, commit the submodule reference update in the parent repository

## Common Workflows

### Scenario 1: Update Library and Use in Application

```bash
# 1. Make changes to the library
cd libs/greeting-lib
# ... edit files ...
git add .
git commit -m "Add new greeting feature"

# 2. Update parent to use new version
cd ../..
git add libs/greeting-lib
git commit -m "Update to latest greeting-lib"

# 3. Test and push
go test ./...
git push
```

### Scenario 2: Pull Latest Changes

```bash
# Pull parent repository changes
git pull

# Update submodules to match
git submodule update --init --recursive
```

### Scenario 3: Clone and Start Development

```bash
# Clone with submodules
git clone --recurse-submodules <url>
cd poc-git-sub-module

# Install dependencies
go mod download

# Start development
make dev
```

## Best Practices

1. **Always Use --recurse-submodules**: When cloning, always use this flag or run `git submodule update --init --recursive` after cloning

2. **Commit Submodules First**: When making changes, commit in the submodule before committing in the parent

3. **Document Submodule Dependencies**: Keep README files updated with submodule information

4. **Use Specific Commits**: Don't leave submodules on floating branches in production

5. **Update Regularly**: Keep submodules up to date but test thoroughly after updates

6. **Communicate Changes**: When updating submodules, clearly document what changed and why

## Troubleshooting

### Submodule Directory is Empty

```bash
git submodule update --init --recursive
```

### Submodule Shows Modified

```bash
cd libs/greeting-lib
git status
# If you didn't mean to change anything:
git reset --hard
cd ../..
```

### Can't Pull Latest Changes

```bash
# Reset submodule to tracked commit
git submodule update --force
```

### Removing a Submodule

```bash
# 1. Deinitialize the submodule
git submodule deinit -f libs/greeting-lib

# 2. Remove from Git
git rm -f libs/greeting-lib

# 3. Remove from .git directory
rm -rf .git/modules/libs/greeting-lib

# 4. Commit the change
git commit -m "Remove greeting-lib submodule"
```

## Real-World Use Cases

Git submodules are commonly used for:

1. **Shared Libraries**: Like in this POC, sharing common code across projects
2. **Third-Party Dependencies**: Including external projects that aren't in package managers
3. **Theme/Plugin Systems**: WordPress, Hugo, and other systems use submodules for themes
4. **Documentation**: Keeping documentation in a separate repository
5. **Microservices**: Managing related services in separate repositories

## Alternatives to Submodules

Consider these alternatives depending on your needs:

- **Go Modules**: For Go-specific dependencies (still recommended alongside submodules)
- **Git Subtree**: Merges the external project into your repository
- **Package Managers**: npm, pip, Maven, etc. for language-specific dependencies
- **Monorepos**: Keeping everything in one repository

## Summary

Git submodules in this project demonstrate:
- ✅ How to structure a modular Go application
- ✅ How to track specific versions of dependencies
- ✅ How to share code between projects
- ✅ How to maintain separation of concerns
- ✅ How to use local Go modules with replace directives

For more information, see:
- [Git Submodules Official Docs](https://git-scm.com/book/en/v2/Git-Tools-Submodules)
- [GitHub Submodules Guide](https://github.blog/2016-02-01-working-with-submodules/)
