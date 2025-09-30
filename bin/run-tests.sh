#!/bin/bash

# KitCLA Test Runner
#
# This script runs the KitCLA test suite with various options.
# Tests validate components, check HTML output, and generate documentation.
#
# Usage:
#   ./bin/run-tests.sh [OPTIONS] [PATH]
#
# Options:
#   -v, --verbose    Run tests with verbose output
#   -c, --coverage   Run tests with coverage report
#   -h, --help       Show this help message
#
# Examples:
#   ./bin/run-tests.sh                    # Run all tests
#   ./bin/run-tests.sh -v                 # Run all tests verbosely
#   ./bin/run-tests.sh ./tests/atoms/...  # Run only atom tests
#   ./bin/run-tests.sh -c                 # Run with coverage
#   ./bin/run-tests.sh -v ./tests/atoms/buttons/...  # Verbose button tests

set -e

# Get the directory where this script is located
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"

# Change to project root
cd "$PROJECT_ROOT"

# Parse arguments
VERBOSE=""
COVERAGE=""
TEST_PATH="./tests/..."

show_help() {
    echo "KitCLA Test Runner"
    echo ""
    echo "Usage: ./bin/run-tests.sh [OPTIONS] [PATH]"
    echo ""
    echo "Options:"
    echo "  -v, --verbose    Run tests with verbose output"
    echo "  -c, --coverage   Run tests with coverage report"
    echo "  -h, --help       Show this help message"
    echo ""
    echo "Examples:"
    echo "  ./bin/run-tests.sh                    # Run all tests"
    echo "  ./bin/run-tests.sh -v                 # Run all tests verbosely"
    echo "  ./bin/run-tests.sh ./tests/atoms/...  # Run only atom tests"
    echo "  ./bin/run-tests.sh -c                 # Run with coverage"
    echo "  ./bin/run-tests.sh -v ./tests/atoms/buttons/...  # Verbose button tests"
    echo ""
    echo "Test Categories:"
    echo "  ./tests/...                    All tests (80+ test files)"
    echo "  ./tests/atoms/...              Basic components"
    echo "  ./tests/molecules/...          Simple combinations"
    echo "  ./tests/organisms/...          Complex assemblies"
    echo "  ./tests/widgets/...            Complete application examples"
    exit 0
}

while [[ $# -gt 0 ]]; do
    case $1 in
        -v|--verbose)
            VERBOSE="-v"
            shift
            ;;
        -c|--coverage)
            COVERAGE="-cover"
            shift
            ;;
        -h|--help)
            show_help
            ;;
        *)
            TEST_PATH="$1"
            shift
            ;;
    esac
done

# Build the go test command
CMD="go test $VERBOSE $COVERAGE $TEST_PATH"

echo "Running KitCLA tests..."
echo "Command: $CMD"
echo ""

# Run the tests
exec $CMD