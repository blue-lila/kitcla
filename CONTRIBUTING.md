# Contributing to KitCLA

Thank you for your interest in contributing to KitCLA! This guide will help you get started with development, testing, and contributing to the project.

## Development Scripts

KitCLA includes several helper scripts in the `bin/` directory:

```bash
# Run tests
./bin/run-tests.sh                 # Run all tests
./bin/run-tests.sh -v              # Verbose output
./bin/run-tests.sh -c              # With coverage

# Run code quality checks
./bin/run-code-checks.sh           # Run linters and security checks

# Pre-push validation
./bin/pre-push-check.sh            # Run tests + code checks (used by git hook)

# Component book server
./bin/run-book.sh                  # Start component documentation server
./bin/run-book.sh 8080             # Start on custom port
```

## Testing

KitCLA includes a comprehensive testing system with 80+ test files covering all component levels. The testing framework validates HTML output, supports regression testing, and automatically generates documentation.

### Running Tests

```bash
# Using the test runner script (recommended)
./bin/run-tests.sh                 # Run all tests
./bin/run-tests.sh -v              # Run with verbose output
./bin/run-tests.sh -c              # Run with coverage report
./bin/run-tests.sh -h              # Show help and available options

# Test specific component categories
./bin/run-tests.sh ./tests/atoms/...       # Basic components
./bin/run-tests.sh ./tests/molecules/...   # Simple combinations
./bin/run-tests.sh ./tests/organisms/...   # Complex assemblies
./bin/run-tests.sh ./tests/widgets/...     # Complete application examples

# Using go test directly
go test ./tests/...                        # Run all tests
go test ./tests/atoms/buttons/...          # Test specific components
go test -v ./tests/atoms/inputs/text_input_test.go  # Verbose single test
```

### Test Structure

Tests are organized following the atomic design structure:

```
tests/
├── atoms/           # Button, input, cell tests
├── molecules/       # Alert, navbar, tree tests
├── organisms/       # Form, table, card tests
└── widgets/         # Complete application examples
```

Each test directory includes:
- **Test files**: `*_test.go` with component validation
- **Data files**: `data/*.html` with expected HTML output
- **Themes**: Gardening application examples for realistic testing

### Test Utilities

The `sup` package provides testing utilities:

```go
import "github.com/blue-lila/kitcla/sup"

// Update expected HTML output (for component changes)
sup.UpdateEqualHtmlFromDataFile(t, html, "./data/button_primary.html")

// Assert HTML matches expected output (regression testing)
sup.AssertEqualHtmlFromDataFile(t, html, "./data/button_primary.html")

// Generate documentation pages
sup.AddPage("PrimaryButton", html)
sup.AddIndexPage("buttons", html)  // Component overview
```

### Component Book Integration

Tests automatically generate documentation for the component book:
- **HTML examples**: Rendered component demonstrations
- **JSON metadata**: Component organization and navigation
- **Showcase pages**: Complete component variations and states

All test data uses consistent gardening themes (plant management, watering schedules, garden tools) making examples realistic and relatable.