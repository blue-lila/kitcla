#!/bin/bash

# Push to GitHub - Public Distribution Script
# This script pushes the current GitLab repository to a public GitHub repository

set -e  # Exit on any error

# Configuration
GITHUB_REMOTE="github"
GITHUB_REPO_URL=""  # Set this to your GitHub repository URL
DEFAULT_BRANCH="master"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Helper functions
log_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

log_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

log_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if we're in a git repository
if [ ! -d ".git" ]; then
    log_error "Not in a git repository"
    exit 1
fi

# Check if GitHub repository URL is configured
if [ -z "$GITHUB_REPO_URL" ]; then
    log_error "GITHUB_REPO_URL is not configured in the script"
    log_info "Please edit this script and set GITHUB_REPO_URL to your GitHub repository URL"
    log_info "Example: GITHUB_REPO_URL=\"https://github.com/username/kitcla.git\""
    exit 1
fi

# Check if GitHub remote exists
if ! git remote | grep -q "^${GITHUB_REMOTE}$"; then
    log_info "Adding GitHub remote..."
    git remote add $GITHUB_REMOTE $GITHUB_REPO_URL
    log_success "GitHub remote added"
else
    log_info "GitHub remote already exists, updating URL..."
    git remote set-url $GITHUB_REMOTE $GITHUB_REPO_URL
    log_success "GitHub remote URL updated"
fi

# Get current branch
CURRENT_BRANCH=$(git branch --show-current)
log_info "Current branch: $CURRENT_BRANCH"

# Check if working directory is clean
if [ -n "$(git status --porcelain)" ]; then
    log_warning "Working directory is not clean. Uncommitted changes:"
    git status --short
    echo
    read -p "Do you want to continue? [y/N]: " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        log_info "Aborted by user"
        exit 0
    fi
fi

# Show what will be pushed
log_info "Preparing to push to GitHub..."
log_info "Source: $(git remote get-url origin) ($CURRENT_BRANCH)"
log_info "Target: $GITHUB_REPO_URL ($CURRENT_BRANCH)"
echo

# Confirm before pushing
read -p "Continue with push to GitHub? [y/N]: " -n 1 -r
echo
if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    log_info "Aborted by user"
    exit 0
fi

# Push to GitHub
log_info "Pushing to GitHub..."
if git push $GITHUB_REMOTE $CURRENT_BRANCH; then
    log_success "Successfully pushed to GitHub!"
    log_info "Repository available at: ${GITHUB_REPO_URL%.git}"
else
    log_error "Failed to push to GitHub"
    exit 1
fi

# Push tags if any exist
if [ -n "$(git tag)" ]; then
    log_info "Pushing tags to GitHub..."
    if git push $GITHUB_REMOTE --tags; then
        log_success "Tags pushed successfully"
    else
        log_warning "Failed to push tags"
    fi
fi

log_success "GitHub distribution updated successfully!"