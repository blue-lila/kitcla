#!/bin/bash

# KitCLA Pre-Push Check Script
#
# This script runs both tests and code checks before pushing code.
# It ensures code quality and correctness by running:
#   1. Full test suite (./bin/run-tests.sh)
#   2. Code quality checks (./bin/run-code-checks.sh)
#   3. Shows commits that will be pushed
#
# Usage:
#   ./bin/pre-push-check.sh
#
# Exit codes:
#   0 - All checks passed
#   1 - Tests failed
#   2 - Code checks failed
#   3 - Uncommitted or unstaged changes detected

set -e

# Get the directory where this script is located
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"

# Change to project root
cd "$PROJECT_ROOT"

echo "🚀 Running pre-push checks..."
echo ""

# Check for uncommitted or unstaged changes
if ! git diff-index --quiet HEAD -- 2>/dev/null; then
    echo "❌ Error: You have uncommitted or unstaged changes!"
    echo ""
    echo "Please commit or stash your changes before pushing:"
    echo ""
    git status --short
    echo ""
    exit 3
fi

# Check for untracked files that might need to be committed
UNTRACKED_FILES=$(git ls-files --others --exclude-standard | wc -l | tr -d ' ')
if [ "$UNTRACKED_FILES" -gt 0 ]; then
    echo "⚠️  Warning: You have $UNTRACKED_FILES untracked file(s)"
    echo ""
    git ls-files --others --exclude-standard
    echo ""
    echo "Consider adding them with 'git add' if they should be committed."
    echo ""
fi

# Detect the default remote branch
REMOTE_BRANCH=$(git rev-parse --abbrev-ref --symbolic-full-name @{u} 2>/dev/null || echo "origin/$(git rev-parse --abbrev-ref HEAD)")
CURRENT_BRANCH=$(git rev-parse --abbrev-ref HEAD)

# Count commits to be pushed
COMMITS_TO_PUSH=$(git log $REMOTE_BRANCH..$CURRENT_BRANCH --oneline 2>/dev/null | wc -l | tr -d ' ' || echo "0")

echo "📊 Branch: $CURRENT_BRANCH"
echo "📡 Remote: $REMOTE_BRANCH"
echo "📦 Commits to push: $COMMITS_TO_PUSH"
echo ""

# Run tests
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "📝 Step 1/2: Running tests..."
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
if ! ./bin/run-tests.sh; then
    echo ""
    echo "❌ Tests failed! Please fix the failing tests before pushing."
    exit 1
fi

echo ""
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "🔍 Step 2/2: Running code checks..."
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
if ! ./bin/run-code-checks.sh; then
    echo ""
    echo "❌ Code checks failed! Please fix the issues before pushing."
    exit 2
fi

echo ""
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "✅ All pre-push checks passed!"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""

# Show summary of commits that will be pushed
if [ "$COMMITS_TO_PUSH" -gt 0 ]; then
    echo "📊 Push Summary:"
    echo "  • Branch: $CURRENT_BRANCH → $REMOTE_BRANCH"
    echo "  • Commits: $COMMITS_TO_PUSH"

    # Show commit messages (short format)
    echo "  • Commit messages:"
    git --no-pager log $REMOTE_BRANCH..$CURRENT_BRANCH --pretty=format:"    - %s" 2>/dev/null || echo "    Unable to display commits"
    echo ""

    echo "  • Files changed: $(git diff --stat $REMOTE_BRANCH..$CURRENT_BRANCH 2>/dev/null | tail -n 1 | sed 's/^ //' || echo 'N/A')"
    echo ""
fi

echo "You're ready to push! 🎉"