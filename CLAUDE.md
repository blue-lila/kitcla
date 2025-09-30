# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

KitCLA is a Go-based design system kit specialized for CRUD-like applications, following atomic design principles. It provides a hierarchical component system with atoms, molecules, and organisms, built on top of a custom HTML generation framework called "goc". The library is hosted at `github.com/blue-lila/kitcla` and requires Go 1.19+.

**Current Status**: Early access - the project is in active development and the API may change.

## Development Commands

### Testing
```bash
# Using the test runner script (recommended)
./bin/run-tests.sh                                # Run all tests
./bin/run-tests.sh -v                             # Run all tests with verbose output
./bin/run-tests.sh -c                             # Run all tests with coverage
./bin/run-tests.sh ./tests/atoms/...              # Run tests for specific component group
./bin/run-tests.sh -v ./tests/atoms/buttons/...   # Run specific tests verbosely

# Using go test directly
go test ./...                                     # Run all tests
go test ./tests/atoms/...                         # Run tests for specific component group
go test -v ./tests/atoms/buttons/button_test.go   # Run specific test file with verbose output
go test ./tests/molecules/...                     # Run molecule component tests
go test ./tests/organisms/...                     # Run organism component tests
```

### Building
```bash
go build                         # Build the project
go run tools/generator/main.go   # Run the component generator tool
```

## Architecture

### Core Components

1. **goc (Go Component)** - Custom HTML generation framework located in `goc/goc.go`
   - Provides HTML element creation with `H()` function
   - Supports attributes, unsafe content injection, and self-closing tags
   - Handles proper HTML escaping and attribute rendering
   - Foundation for all UI components with self-closing tag support

2. **Component Base** - Located in `components/component.go`
   - Base struct for all UI components with debug capabilities
   - Provides helper methods like `H()`, `Cc()`, `Ccv()`, `Cs()`, etc.
   - Wrapper methods for common div operations (`Dcs()`, `Ds()`, `Da()`, etc.)
   - Supports Alpine.js directives through `Ti()`, `Tf()`, and `Wa()` methods
   - Debug mode support for development

3. **Kit Structure** - Main entry point in `kit.go`
   - Organizes components into Atoms, Molecules, and Organisms
   - Factory pattern with `New()` function that wires all dependencies
   - Components are organized hierarchically following atomic design
   - Extensive dependency injection system for component relationships

### Component Hierarchy

#### **Atoms** (Basic building blocks):
- **Buttons**: Button, ButtonAlp (Alpine.js version)
- **Inputs**: TextInput, TextInputAlp, IntegerInput, DecimalInput, BooleanInput, SwitchInput, SelectInput, FileInput, DatetimeInput, JsonInput, HiddenInput, RichTextInput, TextAreaInput, AdvancedSelectInput, GridIdInput, RadioInputAlp
- **Cells**: TextCell, PillCell, IntegerCell, DecimalCell, BooleanCell, TimeCell, JsonCell, LongTextCell, RichTextCell, RelationCell
- **Display**: Headers, Icons, Links, Shows (TextShow, RichTextShow), Placeholders, Resources
- **Layout**: Fields, Dropdowns, CardHeaders, CardBodies, CardWrappers, Tabs, Paginations, ButtonsGroups

#### **Molecules** (Simple combinations):
- **Alerts**: Alert components with icons
- **Messages**: Message components
- **Navigation**: Navbars, Trees, Steppers
- **Interactive**: Popovers

#### **Organisms** (Complex combinations):
- **Data Display**: Tables (Table, KeyValueTable), Grids (IconCardGrid, KeyValueGrid)
- **Layout**: Cards, Modals, Sidebars, Menus, Navmenus
- **Forms**: Forms (Form, DeleteForm)
- **Complex**: DuoCards (PeekDuoCard), KeyPoint, Logs
- **Navigation**: Navbars (Navbar, DualNavbar)

### Data Types and Interfaces

- **dat/entity.go**: Defines `Entity` interface with `GetId()` method
- **dat/form.go**: Form data structures
- **dat/table.go**: Table data structures
- **dat/message.go**: Message data structures

### Frontend Integration

- **Alpine.js 3.8.1**: Client-side reactivity and component behavior
- **Tailwind CSS**: Utility-first CSS framework for styling
- **Custom JavaScript**: Common utilities and debug tools
- **Component Files**: Modular JavaScript components

### Testing System

KitCLA includes a comprehensive testing system designed for component validation, regression testing, and documentation generation. The testing architecture supports 80+ test files organized across atomic design levels.

#### Test Structure and Organization

**Directory Layout:**
```
tests/
├── atoms/           # Basic component tests (buttons, inputs, cells, etc.)
├── molecules/       # Simple combination tests (alerts, messages, popovers)
├── organisms/       # Complex component tests (cards, forms, tables)
├── widgets/         # Complete application widgets (gardening examples)
├── components/      # Legacy component tests
└── aria/           # Accessibility tests
```

**Data Storage Pattern:**
- Each test directory contains a `data/` subdirectory for expected HTML output
- Test data files use naming convention: `{component}_{variant}_{index}.html`
- Example: `button_primary_link_1.html` stores expected HTML for primary link button

#### Test Types and Patterns

**1. Component Unit Tests**
- Test individual component methods (H(), PrimaryLink(), etc.)
- Verify HTML output against stored expected results
- Use gardening-themed test data for realistic examples

**2. Showcase Tests**
- Generate comprehensive component overviews
- Create documentation pages via `sup.AddIndexPage()`
- Demonstrate multiple variations and states

**3. Widget Integration Tests**
- Test complete application scenarios (garden overview, plant management)
- Combine multiple components into realistic user interfaces
- Stored in `tests/widgets/gardening/` for thematic consistency

#### Testing Utilities (sup/ package)

**Core Functions:**
- `UpdateEqualHtmlFromDataFile(t, html, filepath)`: Updates expected HTML output files
- `AssertEqualHtmlFromDataFile(t, html, filepath)`: Validates HTML matches expected output
- `AddIndexPage(name, html)`: Generates component overview documentation
- `AddPage(name, html)`: Creates individual component example pages

**Test Data Management:**
```go
// Standard test pattern
h := component.PrimaryLink("Plant Seeds", "/plant", nil)
html := goc.RenderRoot(h)

// Update expected output (run when changing components)
sup.UpdateEqualHtmlFromDataFile(t, html, "./data/primary_link_1.html")

// Assert output matches expected (regression testing)
sup.AssertEqualHtmlFromDataFile(t, html, "./data/primary_link_1.html")

// Generate documentation page
sup.AddPage("PrimaryLink", html)
```

#### Running Tests

**Command Examples:**
```bash
go test ./tests/...                              # Run all tests (80+ files)
go test ./tests/atoms/...                        # Test all atomic components
go test ./tests/atoms/buttons/...                # Test button components only
go test -v ./tests/atoms/inputs/text_input_test.go  # Verbose single test
go test ./tests/widgets/gardening/...            # Test complete widget examples
```

**Test Categories by Count:**
- Atomic components: ~50 test files
- Molecular components: ~10 test files
- Organism components: ~15 test files
- Widget examples: ~5 test files

#### Documentation Generation

Tests automatically generate documentation files in `docs/pages/` structure:
- **HTML files**: Rendered component examples for component book
- **JSON metadata**: Component information for navigation and organization
- **Index pages**: Overview pages showing multiple component variations

#### Test Data Themes

All test data uses consistent gardening application themes:
- **Buttons**: "Plant Seeds", "Water Garden", "Harvest Crops"
- **Inputs**: Plant varieties, watering schedules, garden locations
- **Forms**: Plant management, care tracking, harvest logging
- **Tables**: Plant inventories, care schedules, growth tracking

This thematic consistency makes tests more readable and provides realistic usage examples.

### Code Generation

- **tools/generator/main.go**: Automated test file generation
- Generates test files based on JSON specifications
- Creates component tests with proper package structure
- Supports Pascal case naming conventions

### Internationalization

- **i18n/en.json**: Initial structure for planned multi-language support
- Future support for component-based translation keys
- Framework prepared for additional languages

### Development Utilities

- **aria/aria.go**: ARIA accessibility support
- **goc_utils/utils.go**: Utility functions for goc framework
- **Debug Support**: Built-in debug mode for component development
- **Development Interceptor**: Assets for development workflow

## Key Patterns

1. **Component Initialization**: Always use `kitcla.New()` to get a properly initialized Kit
2. **HTML Generation**: Use the base `Component.H()` method or its shortcuts (`Cc`, `Ccv`, `Cs`, etc.)
3. **Alpine.js Integration**: Use `Ti()`, `Tf()`, and `Wa()` for Alpine.js directives
4. **Testing**: Import the kit, create components, render with `goc.RenderRoot()`, and compare output
5. **Styling**: Components use Tailwind CSS classes via the `css` parameter in helper methods
6. **Component Dependencies**: Components are injected with their dependencies through the Kit factory

## Directory Structure

- `components/`: Component implementations organized by atomic design levels
  - `atoms/`: Basic UI elements (buttons, inputs, cells, etc.)
  - `molecules/`: Simple component combinations
  - `organisms/`: Complex component assemblies
- `tests/`: Test files mirroring component structure with data subdirectories
- `goc/`: HTML generation framework
- `goc_utils/`: Utilities for the goc framework
- `dat/`: Data structures and interfaces
- `assets/`: Frontend assets
  - `libs/`: Third-party libraries (Alpine.js, Tailwind CSS)
  - `css/`: Custom stylesheets
  - `js/`: JavaScript modules
  - `components/`: Frontend component files
- `docs/`: Component documentation with HTML/JSON examples
- `tools/generator/`: Code generation utilities
- `sup/`: Support utilities for testing and development
- `aria/`: ARIA accessibility support
- `i18n/`: Internationalization files
- Always use ./bin/run-book when trying to run the book
- Always run tests before saying a task is completed