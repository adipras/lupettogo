#!/bin/bash

# LupettoGo Development Commit Script
# Automates git commits with standardized messages and version tagging

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Function to get current version from git tags
get_current_version() {
    local version=$(git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0")
    echo $version
}

# Function to bump version
bump_version() {
    local current_version=$1
    local bump_type=$2
    
    # Remove 'v' prefix if present
    version_number=${current_version#v}
    
    # Split version into major.minor.patch
    IFS='.' read -r major minor patch <<< "$version_number"
    
    case $bump_type in
        "major")
            major=$((major + 1))
            minor=0
            patch=0
            ;;
        "minor")
            minor=$((minor + 1))
            patch=0
            ;;
        "patch")
            patch=$((patch + 1))
            ;;
        *)
            print_error "Invalid bump type: $bump_type"
            exit 1
            ;;
    esac
    
    echo "v$major.$minor.$patch"
}

# Function to determine change type based on files
detect_change_type() {
    local changed_files=$(git diff --cached --name-only)
    
    if echo "$changed_files" | grep -q -E "(\.go|main\.go|go\.mod|go\.sum)$"; then
        if echo "$changed_files" | grep -q -E "(cmd/|internal/|templates/)"; then
            echo "feat"
        else
            echo "fix"
        fi
    elif echo "$changed_files" | grep -q -E "(\.md|README|CHANGELOG)$"; then
        echo "docs"
    elif echo "$changed_files" | grep -q -E "(\.gitignore|\.github/|Makefile|Dockerfile)$"; then
        echo "chore"
    else
        echo "feat"
    fi
}

# Function to generate commit message
generate_commit_message() {
    local change_type=$1
    local description=$2
    local version=$3
    
    echo "${change_type}: ${description}

🚀 Version: ${version}

🤖 Generated with automated commit script"
}

# Main function
main() {
    print_status "Starting automated commit process..."
    
    # Check if we're in a git repository
    if ! git rev-parse --git-dir > /dev/null 2>&1; then
        print_error "Not in a git repository!"
        exit 1
    fi
    
    # Check if there are staged changes
    if ! git diff --cached --quiet; then
        print_status "Found staged changes"
    else
        print_warning "No staged changes found. Staging all changes..."
        git add .
    fi
    
    # Get current version
    current_version=$(get_current_version)
    print_status "Current version: $current_version"
    
    # Determine change type
    if [ -n "$1" ]; then
        change_type=$1
    else
        change_type=$(detect_change_type)
        print_status "Auto-detected change type: $change_type"
    fi
    
    # Get description
    if [ -n "$2" ]; then
        description=$2
    else
        echo -n "Enter commit description: "
        read description
    fi
    
    # Determine version bump type
    case $change_type in
        "feat")
            bump_type="minor"
            ;;
        "fix")
            bump_type="patch"
            ;;
        "docs"|"chore"|"style"|"refactor")
            bump_type="patch"
            ;;
        "breaking")
            bump_type="major"
            ;;
        *)
            bump_type="patch"
            ;;
    esac
    
    # Ask for version bump confirmation
    echo -n "Version bump type [$bump_type]: "
    read user_bump_type
    if [ -n "$user_bump_type" ]; then
        bump_type=$user_bump_type
    fi
    
    # Calculate new version
    new_version=$(bump_version $current_version $bump_type)
    print_status "New version: $new_version"
    
    # Generate commit message
    commit_message=$(generate_commit_message $change_type "$description" $new_version)
    
    # Show commit message preview
    echo -e "\n${BLUE}Commit message preview:${NC}"
    echo "------------------------"
    echo "$commit_message"
    echo "------------------------"
    
    # Confirm commit
    echo -n "Proceed with commit and tag? (y/N): "
    read confirm
    if [ "$confirm" != "y" ] && [ "$confirm" != "Y" ]; then
        print_warning "Commit cancelled"
        exit 0
    fi
    
    # Create commit
    print_status "Creating commit..."
    git commit -m "$commit_message"
    
    # Create tag
    print_status "Creating tag: $new_version"
    git tag -a "$new_version" -m "Release $new_version"
    
    # Ask about pushing
    echo -n "Push to remote with tags? (y/N): "
    read push_confirm
    if [ "$push_confirm" = "y" ] || [ "$push_confirm" = "Y" ]; then
        print_status "Pushing to remote..."
        git push origin main
        git push origin --tags
        print_status "Successfully pushed commit and tags!"
    else
        print_warning "Skipped pushing. Run 'git push origin main && git push origin --tags' manually."
    fi
    
    print_status "Commit process completed!"
    echo -e "${GREEN}✅ Committed with version $new_version${NC}"
}

# Help function
show_help() {
    echo "Usage: $0 [change_type] [description]"
    echo ""
    echo "change_type options:"
    echo "  feat     - New feature (minor version bump)"
    echo "  fix      - Bug fix (patch version bump)"
    echo "  docs     - Documentation changes (patch version bump)"
    echo "  chore    - Maintenance tasks (patch version bump)"
    echo "  breaking - Breaking changes (major version bump)"
    echo ""
    echo "Examples:"
    echo "  $0 feat 'add new CLI command'"
    echo "  $0 fix 'resolve doctor command issues'"
    echo "  $0"
    echo ""
    echo "If no arguments provided, script will prompt for input."
}

# Check for help flag
if [ "$1" = "--help" ] || [ "$1" = "-h" ]; then
    show_help
    exit 0
fi

# Run main function
main "$@"