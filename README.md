# KitCLA

A Go-based design system kit specialized for CRUD-like applications, following atomic design principles and built on a custom HTML generation framework called "goc".

> ⚠️ **Early Access**: This project is in active development and the API may change. Use in production at your own discretion.

## Overview

KitCLA provides a hierarchical component system with atoms, molecules, and organisms for building web user interfaces. The library generates HTML server-side with optional client-side Alpine.js enhancements and Tailwind CSS styling.

> 💡 **Component Book**: All components can be viewed in the interactive Component Book. Run `./bin/run-book.sh` to explore live examples and documentation at `http://localhost:7000`.

## Features

- **Atomic Design Architecture**: Organized components following atoms → molecules → organisms pattern
- **Type-Safe HTML Generation**: Custom `goc` framework for safe HTML creation
- **Alpine.js 3.8.1 Integration**: Client-side reactivity and interactive components
- **Tailwind CSS 2.2.19 Styling**: Utility-first CSS framework integration
- **Comprehensive Testing**: Extensive test coverage with HTML output validation
- **Accessibility Support**: Built-in ARIA support
- **Internationalization**: Multi-language support system (planned)

## Requirements

- Go 1.25+
- No external dependencies (self-contained)

## Installation

```bash
go get github.com/blue-lila/kitcla
```

## Quick Start

```go
package main

import (
    "github.com/blue-lila/kitcla"
    "github.com/blue-lila/kitcla/goc"
)

func main() {
    // Initialize the component library
    kit := kitcla.New()

    // Create a button component
    button := kit.Atoms.Buttons.Button.PrimaryLink("Click Me", "/my-page", nil)

    // Render HTML
    html := goc.RenderRoot(button)
    println(html)
}
```

## Component Architecture

KitCLA follows atomic design principles with three levels of components:

- **Atoms**: Basic building blocks (buttons, inputs, cells, icons, etc.)
- **Molecules**: Simple combinations of atoms (alerts, navbars, trees, etc.)
- **Organisms**: Complex assemblies (tables, forms, cards, modals, etc.)

For a complete list of all available components and their usage, see [docs/components.md](docs/components.md) or explore [kit.go](kit.go) directly.

### Alpine.js Variants

Some components have Alpine.js-enhanced variants (suffixed with `Alp`) for client-side interactivity:

```go
// Standard button with server-side link
kit.Atoms.Buttons.Button.PrimaryLink("Save", "/save", nil)

// Alpine.js button with client-side click handler
kit.Atoms.Buttons.ButtonAlp.PrimaryLink("Save", "handleSave()", nil)
```

### Mod Pattern

Most component methods accept a `Mod` struct as their last argument for optional customization:

```go
// Using default behavior (pass nil for mod)
button := kit.Atoms.Buttons.Button.PrimaryLink("Save", "/save", nil)

// Customizing with Mod
mod := &buttons.ButtonMod{
    Icon:     "floppy-disk",
    Disabled: false,
    Title:    "Save your changes",
}
button := kit.Atoms.Buttons.Button.PrimaryLink("Save", "/save", mod)
```

The Mod pattern allows you to override defaults while keeping convenience methods simple.

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

Div-specific shortcuts exist too: `Dcs`, `Ds`, `Da`, ...

Alpine.js specific:
- `Ti(condition, component)` - Template with x-if
- `Tf(condition, component)` - Template with x-for
- `Wa(attributes, components...)` - Wrapper with attributes

Utility methods:
- `W(css, components...)` - Wrapper div
- `ExpHtml(html)` - Experimental raw HTML (unsafe)
- `OrNil(element, isNil)` - Conditionally include element

## Composite Example

Here's a practical example showing how to combine atoms, organisms, and component glue to build a plant inventory table:

```go
package main

import (
    "github.com/blue-lila/kitcla"
    "github.com/blue-lila/kitcla/components/organisms/tables"
    "github.com/blue-lila/kitcla/dat"
    "github.com/blue-lila/kitcla/goc"
)

// Plant implements dat.Entity interface
type Plant struct {
    Id       string
    Name     string
    Species  string
    Location string
    Status   string
}

func (p Plant) GetId() string {
    return p.Id
}

func main() {
    kit := kitcla.New()

    // Sample data
    plants := []dat.Entity{
        Plant{Id: "1", Name: "Monstera Deliciosa", Species: "Swiss Cheese Plant", Location: "Living Room", Status: "Healthy"},
        Plant{Id: "2", Name: "Garden Rose", Species: "Rosa gallica", Location: "Balcony", Status: "Needs Water"},
        Plant{Id: "3", Name: "Snake Plant", Species: "Sansevieria", Location: "Bedroom", Status: "Healthy"},
    }

    // Define table columns
    columns := kit.Organisms.Tables.Table.Columns([]*tables.Column{
        {Key: "name", Label: "Plant Name", Kind: tables.ColumnKindData},
        {Key: "species", Label: "Species", Kind: tables.ColumnKindData},
        {Key: "location", Label: "Location", Kind: tables.ColumnKindData},
        {Key: "status", Label: "Status", Kind: tables.ColumnKindData},
    })

    // Build table with custom cell renderer
    tableMod := &tables.TableMod{
        Title:      "🌱 Plant Inventory",
        Columns:    columns,
        Items:      plants,
        BaseUrl:    "/plants",
        RowCell: func(item dat.Entity, column *tables.Column) goc.HTML {
            plant := item.(Plant)
            switch column.Key {
            case "name":
                return kit.Component.Dcs("flex items-center space-x-2",
                    kit.Component.Dcv("text-xl", "🌿"),
                    kit.Atoms.Links.Link.TextLink(plant.Name, "/plants/"+plant.Id),
                )
            case "species":
                return kit.Component.Dcv("text-gray-600 text-sm", plant.Species)
            case "location":
                return kit.Atoms.Cells.TextCell.TextCell(plant.Location)
            case "status":
                statusColor := "bg-green-100 text-green-800"
                if plant.Status == "Needs Water" {
                    statusColor = "bg-yellow-100 text-yellow-800"
                }
                return kit.Component.Dcv("inline-flex px-2 py-1 text-xs font-semibold rounded-full "+statusColor, plant.Status)
            default:
                return goc.HTML{}
            }
        },
    }

    // Wrap table in a page layout using component glue
    page := kit.Component.Dcs("min-h-screen bg-gray-50",
        kit.Component.Dcs("max-w-7xl mx-auto px-6 py-8",
            kit.Component.Ds(
                kit.Atoms.Headers.Header.H1("My Garden"),
                kit.Component.Dcv("text-gray-600 mb-8", "Track and manage your plant collection"),
                kit.Organisms.Tables.Table.H(tableMod),
            ),
        ),
    )

    // Render
    html := goc.RenderRoot(page)
    println(html)
}
```

This example demonstrates:
- **Atoms**: Button, Links, Cells for basic UI elements
- **Organisms**: Table for complex data display with custom cell rendering
- **Component Glue**: Using `kit.Component` methods (`Dcs`, `Ds`, `Dcv`) to structure and style the layout
- **Data Integration**: Implementing `dat.Entity` interface and passing structured data to components

## Documentation

- **[Component Reference](docs/components.md)** - Complete list of all components
- **[Advanced Usage](docs/advanced.md)** - API reference, frontend integration, and advanced patterns
- **[Contributing Guide](docs/contributing.md)** - Development setup, testing, and contribution guidelines

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.