#!/bin/bash

# KitCLA Book Server Runner
#
# This script runs the KitCLA component book server.
# The server displays components generated from tests in a web interface.
#
# Usage:
#   ./bin/run-book.sh [PORT]
#
# Examples:
#   ./bin/run-book.sh        # Run on default port 7000
#   ./bin/run-book.sh 8080   # Run on port 8080

set -e

# Get the directory where this script is located
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"

# Change to project root
cd "$PROJECT_ROOT"

# Set port from argument or use default
if [ $# -eq 1 ]; then
    export PORT="$1"
    echo "Starting KitCLA Book server on port $PORT..."
else
    echo "Starting KitCLA Book server on default port 7000..."
fi

# Run the book server (include all .go files in the package)
exec go run tools/book/*.go