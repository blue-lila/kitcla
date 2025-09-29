# KitCLA Components

This document provides a comprehensive reference of all components available in KitCLA, organized by atomic design principles.

## Component Hierarchy

### Atoms (Basic Building Blocks)

**Buttons**
- `Button` - Standard HTML buttons
- `ButtonAlp` - Alpine.js enhanced buttons

**Inputs**
- `TextInput`, `TextInputAlp` - Text input fields
- `IntegerInput`, `DecimalInput` - Numeric inputs
- `BooleanInput`, `SwitchInput` - Boolean controls
- `SelectInput`, `SelectInputAlp` - Dropdown selections
- `FileInput` - File upload controls
- `DatetimeInput` - Date/time pickers
- `JsonInput` - JSON editors
- `RichTextInput`, `TextAreaInput` - Rich text controls

**Display Components**
- `TextCell`, `PillCell`, `BooleanCell` - Data display cells
- `Headers`, `Icons`, `Links` - Basic UI elements
- `Placeholders`, `Resources` - Content placeholders

**Layout Components**
- `Fields`, `Dropdowns` - Form layout elements
- `CardHeaders`, `CardBodies`, `CardWrappers` - Card components
- `Tabs`, `Paginations` - Navigation elements

### Molecules (Simple Combinations)

- `Alert` - Alert messages with icons
- `Message` - User messages
- `Navbar` - Navigation bars
- `Tree` - Hierarchical tree views
- `Stepper` - Step-by-step workflows
- `Popover` - Overlay content

### Organisms (Complex Assemblies)

**Data Display**
- `Table`, `KeyValueTable` - Data tables
- `IconCardGrid`, `KeyValueGrid` - Grid layouts

**Layout Systems**
- `Card` - Content cards
- `Modal` - Modal dialogs
- `Sidebar` - Side navigation
- `Menu`, `Navmenu` - Menu systems

**Forms**
- `Form` - Standard forms
- `DeleteForm` - Deletion confirmation forms

**Complex Components**
- `PeekDuoCard` - Advanced card layouts
- `KeyPoint` - Key information displays
- `Logs` - Log viewers
- `DualNavbar` - Advanced navigation

For detailed usage examples and API documentation, see the main [README.md](../README.md).