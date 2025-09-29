# Advanced Usage

This document covers advanced features and patterns for working with KitCLA.

## Frontend Integration

### Alpine.js Components

```go
// Alpine.js enhanced button
buttonAlp := kit.Atoms.Buttons.ButtonAlp.H("Interactive Button")

// Alpine.js directives
comp.Ti("show", content)  // x-if directive
comp.Tf("item in items", template)  // x-for directive
```

### Styling with Tailwind CSS

Components use Tailwind CSS classes:

```go
// Apply Tailwind classes
button := comp.Cc("button", "bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded")
```

### Assets Structure

- `assets/libs/` - Alpine.js 3.8.1, Tailwind CSS
- `assets/css/` - Custom stylesheets
- `assets/js/` - JavaScript modules
- `assets/components/` - Frontend component files

## Development

### Debug Mode

Enable debug mode for development:

```go
kit := kitcla.New()
kit.Component.Data.Debug = true
```

### ARIA Accessibility

Built-in accessibility support:

```go
import "github.com/blue-lila/kitcla/aria"

// Use ARIA attributes in components
button := comp.Ca("button", goc.Attr{
    "aria-label": "Close dialog",
    "aria-pressed": "false",
})
```

### Internationalization

Internationalization support is planned. The `i18n/en.json` file contains initial structure for future multi-language support.

## API Reference

### Core Types

```go
// Kit - Main component factory
type Kit struct {
    Component *components.Component
    Atoms     struct { /* ... */ }
    Molecules struct { /* ... */ }
    Organisms struct { /* ... */ }
}

// Component - Base component with HTML generation methods
type Component struct {
    Data struct {
        Debug bool
    }
}

// HTML - Represents an HTML element
type HTML struct {
    El    string
    Attrs []interface{}
}
```

### Key Methods

```go
// Initialize kit
kit := kitcla.New()

// HTML generation
goc.H(element, attributes...)           // Create HTML element
goc.RenderRoot(html)                   // Render to string

// Component shortcuts
comp.H(el, attrs...)                   // Base HTML creation
comp.Cc(el, css)                       // Element with CSS
comp.Ccv(el, css, values...)           // Element with CSS and content
comp.Cs(el, components...)             // Element with subcomponents
comp.Dcs(css, components...)           // Div with CSS and subcomponents
```

## Advanced Patterns

### Component Composition

```go
kit := kitcla.New()

// Build complex components from atoms
header := kit.Atoms.Headers.Header.H("Page Title", "h1", "text-2xl font-bold")
button := kit.Atoms.Buttons.Button.H("Action", "primary", "/action")
card := kit.Organisms.Cards.Card.H([]goc.HTML{header, button}, "card-container")
```

### Conditional Rendering

```go
comp := kit.Component

// Using OrNil for conditional elements
elements := []goc.HTML{}
elements = append(elements, comp.OrNil(errorMessage, hasError)...)
elements = append(elements, comp.OrNil(successMessage, isSuccess)...)
```

### Dynamic Content

```go
// Using Alpine.js for dynamic behavior
dynamicList := comp.Tf("item in items",
    comp.Ccv("li", "list-item", "{{ item.name }}"))

conditionalContent := comp.Ti("showDetails",
    comp.Ccv("div", "details", "Additional information"))
```

### Custom Styling

```go
// Combining Tailwind with custom CSS
styledButton := comp.Cc("button",
    "bg-gradient-to-r from-blue-500 to-purple-600 hover:from-blue-600 hover:to-purple-700 " +
    "text-white font-semibold py-2 px-4 rounded-lg shadow-md hover:shadow-lg " +
    "transform hover:scale-105 transition-all duration-200")
```

For basic usage and getting started, see the main [README.md](../README.md).