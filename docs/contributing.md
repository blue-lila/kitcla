# Contributing to KitCLA

Thank you for your interest in contributing to KitCLA! This document provides guidelines for contributing to the project.

## Development Setup

1. Clone the repository
2. Ensure you have Go 1.19+ installed
3. Run tests to verify your setup: `go test ./...`

## Testing

### Running Tests

```bash
# Run all tests
go test ./...

# Run specific component tests
go test ./tests/atoms/buttons/...
go test ./tests/molecules/...
go test ./tests/organisms/...

# Verbose output
go test -v ./tests/atoms/buttons/button_test.go
```

### Test Structure

Tests follow the component hierarchy:
- `tests/atoms/` - Atomic component tests
- `tests/molecules/` - Molecular component tests
- `tests/organisms/` - Organism component tests
- `tests/*/data/` - Test data files with expected HTML output

### Writing Tests

When creating new components, follow the existing test patterns:

```go
func TestComponentName(t *testing.T) {
    kit := kitcla.New()
    component := kit.Atoms.ComponentType.ComponentName

    h := component.MethodName(params...)
    html := goc.RenderRoot(h)

    sup.UpdateEqualHtmlFromDataFile(t, html, "./data/expected_output.html")
    sup.AssertEqualHtmlFromDataFile(t, html, "./data/expected_output.html")
}
```

### Test Data Files

- Store expected HTML output in `data/` subdirectories within test directories
- Use descriptive filenames that match the test case
- Keep HTML properly formatted and readable

## Contributing Guidelines

### 1. Follow Atomic Design Principles

Organize components according to atomic design:
- **Atoms**: Basic building blocks (buttons, inputs, cells)
- **Molecules**: Simple combinations of atoms
- **Organisms**: Complex component assemblies

### 2. Use Existing Component Patterns

Before creating new components:
- Study existing components in the same category
- Follow the same dependency injection patterns
- Use consistent naming conventions
- Maintain the same API style

### 3. Write Comprehensive Tests

- Every new component must have tests
- Test all public methods and variations
- Include edge cases and error conditions
- Ensure HTML output matches expectations

### 4. Update Documentation

When adding new components or features:
- Update `docs/components.md` with new component listings
- Add usage examples if the API is complex
- Update `docs/advanced.md` for advanced features
- Keep README.md focused on getting started

### 5. Follow Go Conventions

- Use proper Go naming conventions (PascalCase for exported types/functions)
- Write clear, self-documenting code
- Follow standard Go project structure
- Use meaningful variable and function names

### 6. Code Style

- Follow existing code formatting and style
- Use the same patterns for HTML generation
- Maintain consistent indentation and spacing
- Group related functionality logically

## Component Development

### Creating New Atoms

1. Create the component file in `components/atoms/category/`
2. Follow the existing struct pattern with Component embedding
3. Add the component to the Kit struct in `kit.go`
4. Wire dependencies in the `New()` function
5. Write tests in `tests/atoms/category/`

### Creating New Molecules

1. Create in `components/molecules/category/`
2. Define dependencies on atoms or other molecules
3. Add to Kit struct and wire dependencies
4. Write comprehensive tests

### Creating New Organisms

1. Create in `components/organisms/category/`
2. Define dependencies on atoms, molecules, or other organisms
3. Add to Kit struct and wire dependencies
4. Write tests covering complex interactions

## Code Review Process

1. Create a feature branch from main
2. Make your changes following these guidelines
3. Ensure all tests pass: `go test ./...`
4. Submit a pull request with:
   - Clear description of changes
   - Reference to any related issues
   - Screenshots for UI changes
   - Test coverage for new features

## Architecture Considerations

### Component Dependencies

- Components should declare their dependencies explicitly
- Use dependency injection through the Kit factory
- Avoid circular dependencies between components
- Keep atom components dependency-free when possible

### HTML Generation

- Use the `goc` framework for all HTML generation
- Leverage Component helper methods (`Cc`, `Ccv`, `Cs`, etc.)
- Ensure proper HTML escaping and safety
- Follow semantic HTML practices

### Styling Integration

- Use Tailwind CSS classes for styling
- Support custom CSS classes through parameters
- Maintain responsive design principles
- Consider accessibility in styling choices

## Getting Help

- Check existing issues in the repository
- Review the documentation in the `docs/` folder
- Look at existing component implementations for patterns
- Ask questions in pull requests or issues

## License

By contributing to KitCLA, you agree that your contributions will be licensed under the MIT License.