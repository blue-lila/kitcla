#!/bin/bash

set -e

echo "🔍 Running automated code checks..."

# Check Go version requirement (must be 1.25 or higher)
echo "🔍 Checking Go version..."
GO_VERSION=$(go version | grep -oE 'go[0-9]+\.[0-9]+' | sed 's/go//')
MAJOR_VERSION=$(echo $GO_VERSION | cut -d. -f1)
MINOR_VERSION=$(echo $GO_VERSION | cut -d. -f2)

if [ "$MAJOR_VERSION" -lt 1 ] || ([ "$MAJOR_VERSION" -eq 1 ] && [ "$MINOR_VERSION" -lt 25 ]); then
    echo "❌ Go version $GO_VERSION is too old. Minimum required version is 1.25"
    echo "Please upgrade Go to version 1.25 or higher"
    exit 1
fi
echo "✅ Go version $GO_VERSION meets minimum requirement (1.25+)"

# Check if golangci-lint is installed or needs updating
if ! command -v golangci-lint &> /dev/null; then
    echo "📦 Installing golangci-lint ..."
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.4.0
fi

# Check if govulncheck is installed
if ! command -v govulncheck &> /dev/null; then
    echo "📦 Installing govulncheck..."
    go install golang.org/x/vuln/cmd/govulncheck@latest
fi

echo "✅ Running go mod tidy..."
go mod tidy

echo "✅ Running go vet..."
go vet .

echo "✅ Running golangci-lint..."
golangci-lint run

echo "✅ Running govulncheck for security vulnerabilities..."
govulncheck ./...

echo "🎉 All code checks passed!"

