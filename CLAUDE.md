# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

KitCLA is a Go-based design system kit specialized for CRUD-like applications, following atomic design principles. It provides a hierarchical component system with atoms, molecules, and organisms, built on top of a custom HTML generation framework called "goc". The library is hosted at `github.com/blue-lila/kitcla` and requires Go 1.19+.

**Current Status**: Early access - the project is in active development and the API may change.

## Development Commands

### Testing
```bash
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

### Testing Structure

- Tests are located in `tests/` directory, mirroring the component structure
- Each test file includes component instantiation via `kitcla.New()`
- Tests verify HTML output using `goc.RenderRoot()`
- Test data files are stored in `data/` subdirectories
- Support utilities in `sup/` for test assertions and HTML comparison

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