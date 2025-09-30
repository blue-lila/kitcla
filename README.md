# KitCLA

A Go-based design system kit specialized for CRUD-like applications, following atomic design principles and built on a custom HTML generation framework called "goc".

> ⚠️ **Early Access**: This project is in active development and the API may change. Use in production at your own discretion.

## Overview

KitCLA provides a hierarchical component system with atoms, molecules, and organisms for building web user interfaces. The library generates HTML server-side with optional client-side Alpine.js enhancements and Tailwind CSS styling.

## Features

- **Atomic Design Architecture**: Organized components following atoms → molecules → organisms pattern
- **Type-Safe HTML Generation**: Custom `goc` framework for safe HTML creation
- **Alpine.js Integration**: Client-side reactivity and interactive components
- **Tailwind CSS Styling**: Utility-first CSS framework integration
- **Comprehensive Testing**: Extensive test coverage with HTML output validation
- **Accessibility Support**: Built-in ARIA support
- **Internationalization**: Multi-language support system (planned)

## Requirements

- Go 1.19+
- No external dependencies (self-contained)

## Installation

```bash
go get github.com/blue-lila/kitcla
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/blue-lila/kitcla"
    "github.com/blue-lila/kitcla/goc"
)

func main() {
    // Initialize the component library
    kit := kitcla.New()

    // Create a button component
    button := kit.Atoms.Buttons.Button.H("Click Me", "primary", "")

    // Render HTML
    html := goc.RenderRoot(button)
    fmt.Println(html)
}
```

## Component Architecture

KitCLA follows atomic design principles with three levels of components:

- **Atoms**: Basic building blocks (buttons, inputs, cells, icons, etc.)
- **Molecules**: Simple combinations of atoms (alerts, navbars, trees, etc.)
- **Organisms**: Complex assemblies (tables, forms, cards, modals, etc.)

For a complete list of all available components and their usage, see [docs/components.md](docs/components.md).

## HTML Generation with goc

The `goc` (Go Component) framework provides type-safe HTML generation:

```go
// Basic HTML element
element := goc.H("div", goc.Attr{"class": "container"}, "Hello World")

// Using component helpers
kit := kitcla.New()
comp := kit.Component

// Element with CSS class
div := comp.Cc("div", "bg-blue-500")

// Element with CSS and content
div := comp.Ccv("div", "text-white", "Content")

// Element with attributes and subcomponents
div := comp.Cas("div", goc.Attr{"id": "main"}, subComponent)
```

### Component Helper Methods

All components are built using the base `H()` method which accepts:
- `el` - HTML element name (string)
- `attrs...` - Variable arguments that can be:
  - `goc.Attr{"key": "value"}` - HTML attributes
  - `string` - Text content (automatically escaped)
  - `[]string` - Multiple text values
  - `[]goc.HTML` - Child components
  - `goc.HTML` - Single child component
  - `goc.UnsafeContent()` - Raw HTML (use with caution)

The base Component provides many convenience methods built on `H()`:

- `C(el)` - Element only
- `Cc(el, css)` - Element with CSS class
- `Ccv(el, css, values...)` - Element with CSS and content
- `Cv(el, values...)` - Element with values
- `Cs(el, components...)` - Element with subcomponents
- `Ca(el, attributes)` - Element with attributes
- `Cas(el, attributes, components...)` - Element with attributes and subcomponents
- `Ccs(el, css, components...)` - Element with CSS and subcomponents

Div-specific shortcuts:
- `Dcs(css, components...)` - Div with CSS and subcomponents
- `Ds(components...)` - Div with subcomponents
- `Da(attributes)` - Div with attributes
- `Dc(css)` - Div with CSS
- `Dv(values...)` - Div with values
- `Dcv(css, value)` - Div with CSS and value
- `Dav(attributes, value)` - Div with attributes and value
- `Das(attributes, components...)` - Div with attributes and subcomponents

Alpine.js specific:
- `Ti(condition, component)` - Template with x-if
- `Tf(condition, component)` - Template with x-for
- `Wa(attributes, components...)` - Wrapper with attributes

Utility methods:
- `W(css, components...)` - Wrapper div
- `ExpHtml(html)` - Expand raw HTML (unsafe)
- `OrNil(element, isNil)` - Conditionally include element

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

## Documentation

- **[Component Reference](docs/components.md)** - Complete list of all components
- **[Advanced Usage](docs/advanced.md)** - API reference, frontend integration, and advanced patterns
- **[Contributing Guide](docs/contributing.md)** - Development setup, testing, and contribution guidelines

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.