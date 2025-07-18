# Development Commit Guide

A comprehensive guide for using the `dev-commit.sh` script to automate git commits with semantic versioning and consistent messaging.

## Overview

The `dev-commit.sh` script automates the development workflow by:
- Standardizing commit messages
- Automatically managing semantic versioning
- Creating git tags
- Pushing changes and tags to remote repository

## Prerequisites

- Git repository initialized
- Bash shell environment
- Git configured with user credentials

## Installation

The script is already included in the LupettoGo project root. Make sure it's executable:

```bash
chmod +x ./dev-commit.sh
```

## Usage

### Interactive Mode (Recommended)

Run the script without arguments for guided prompts:

```bash
./dev-commit.sh
```

The script will:
1. Auto-detect change type based on modified files
2. Prompt for commit description
3. Suggest version bump type
4. Show commit message preview
5. Ask for confirmation before committing

### Command Line Mode

Provide change type and description directly:

```bash
./dev-commit.sh <change_type> "<description>"
```

**Examples:**
```bash
./dev-commit.sh feat "add new CLI command for module generation"
./dev-commit.sh fix "resolve doctor command version checking"
./dev-commit.sh docs "update README with new installation steps"
```

### Help

Display usage information:

```bash
./dev-commit.sh --help
```

## Change Types

The script supports semantic commit types that determine version bumping:

| Change Type | Description | Version Bump | Example |
|-------------|-------------|--------------|---------|
| `feat` | New feature | Minor (0.1.0) | Adding new CLI command |
| `fix` | Bug fix | Patch (0.0.1) | Fixing existing functionality |
| `docs` | Documentation | Patch (0.0.1) | README updates |
| `chore` | Maintenance | Patch (0.0.1) | Dependencies, build scripts |
| `breaking` | Breaking change | Major (1.0.0) | API changes |

## Auto-Detection Logic

When no change type is specified, the script analyzes modified files:

- **Go files** (`*.go`, `go.mod`, `go.sum`) → `feat` or `fix`
- **Documentation** (`*.md`, `README`, `CHANGELOG`) → `docs`
- **Configuration** (`.gitignore`, `Makefile`, `Dockerfile`) → `chore`
- **Default** → `feat`

## Versioning

The script follows [Semantic Versioning](https://semver.org/) (SemVer):

- **Major** (X.0.0): Breaking changes
- **Minor** (0.X.0): New features (backward compatible)
- **Patch** (0.0.X): Bug fixes and small changes

### Version Bump Examples

Starting from `v1.2.3`:
- `feat` → `v1.3.0`
- `fix` → `v1.2.4`
- `breaking` → `v2.0.0`

## Commit Message Format

The script generates standardized commit messages:

```
<type>: <description>

🚀 Version: <version>

🤖 Generated with automated commit script
```

**Example:**
```
feat: add PostgreSQL support to doctor command

🚀 Version: v1.2.0

🤖 Generated with automated commit script
```

## Workflow Steps

1. **Stage Changes**: Script automatically stages all changes if nothing is staged
2. **Detect Type**: Analyzes files or uses provided type
3. **Get Description**: Prompts for or uses provided description
4. **Version Bump**: Calculates new version based on change type
5. **Preview**: Shows commit message for review
6. **Commit**: Creates commit with formatted message
7. **Tag**: Creates git tag with new version
8. **Push**: Optionally pushes commit and tags to remote

## Interactive Prompts

### Change Type Detection
```
Auto-detected change type: feat
```

### Description Input
```
Enter commit description: add new module generation command
```

### Version Bump Confirmation
```
Version bump type [minor]: 
```
Press Enter to accept or type alternative (`major`, `minor`, `patch`)

### Commit Preview
```
Commit message preview:
------------------------
feat: add new module generation command

🚀 Version: v1.2.0

🤖 Generated with automated commit script
------------------------
```

### Confirmation
```
Proceed with commit and tag? (y/N): y
```

### Push Option
```
Push to remote with tags? (y/N): y
```

## Best Practices

### 1. Stage Specific Changes
```bash
git add specific-file.go
./dev-commit.sh fix "resolve specific issue"
```

### 2. Use Descriptive Messages
```bash
# Good
./dev-commit.sh feat "add PostgreSQL support to doctor command"

# Avoid
./dev-commit.sh feat "updates"
```

### 3. Review Changes Before Committing
```bash
git diff --cached  # Review staged changes
./dev-commit.sh    # Then commit
```

### 4. Use Appropriate Change Types
- Use `feat` for new functionality
- Use `fix` for bug fixes
- Use `docs` for documentation only
- Use `breaking` sparingly for major changes

## Common Scenarios

### Bug Fix
```bash
# Fix issue in doctor command
git add cmd/doctor.go
./dev-commit.sh fix "resolve Go version validation logic"
```

### New Feature
```bash
# Add new CLI command
git add cmd/generate.go internal/generator/module.go
./dev-commit.sh feat "add module generation command"
```

### Documentation Update
```bash
# Update README
git add README.md
./dev-commit.sh docs "add installation instructions"
```

### Multiple File Changes
```bash
# Let script auto-detect
git add .
./dev-commit.sh
```

## Troubleshooting

### Script Not Executable
```bash
chmod +x ./dev-commit.sh
```

### Not in Git Repository
```bash
cd /path/to/your/git/repo
./dev-commit.sh
```

### No Staged Changes
The script will automatically stage all changes. To stage specific files:
```bash
git add file1.go file2.go
./dev-commit.sh
```

### Push Fails
Ensure you have proper git remote configuration:
```bash
git remote -v
git push origin main  # Manual push
git push origin --tags  # Manual tag push
```

## Integration with Development Workflow

### Daily Development
1. Make changes to code
2. Test changes
3. Run `./dev-commit.sh`
4. Continue development

### Release Preparation
1. Complete feature/fix
2. Update documentation
3. Run tests
4. Use `./dev-commit.sh` with appropriate type
5. Tag is automatically created and pushed

### Team Collaboration
- Consistent commit messages across team
- Automatic versioning prevents conflicts
- Clear change history with semantic types

## Advanced Usage

### Custom Version Bump
```bash
./dev-commit.sh feat "add new feature"
# When prompted: Version bump type [minor]: patch
```

### Skip Push
Choose "N" when prompted about pushing to work offline or push manually later.

### Manual Tag Creation
The script handles tagging automatically, but you can also create tags manually:
```bash
git tag -a v1.2.3 -m "Release v1.2.3"
```

## Script Output

The script provides colored output for better visibility:
- 🟢 **Green**: Success messages
- 🟡 **Yellow**: Warnings
- 🔴 **Red**: Errors
- 🔵 **Blue**: Information

## Contributing

When contributing to LupettoGo:
1. Use the `dev-commit.sh` script for all commits
2. Follow the change type conventions
3. Write clear, descriptive commit messages
4. Test changes before committing

---

*This guide covers the complete usage of the `dev-commit.sh` script for LupettoGo development.*