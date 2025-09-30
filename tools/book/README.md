# KitCLA Book

A lightweight book for browsing and viewing KitCLA components generated from tests.

## Usage

### Using the Runner Script (Recommended)

The easiest way to start the book server is using the provided script:

```bash
# Start on default port 7000
./bin/run-book.sh

# Start on a specific port
./bin/run-book.sh 8080
```

### Direct Go Command

Alternatively, you can run the server directly:

```bash
# From the project root
go run tools/book/main.go

# Use a custom port with environment variable
PORT=8080 go run tools/book/main.go
```

The server will start on port 7000 by default. Open your browser to:
```
http://localhost:7000  # or your custom port
```

## Features

- Browse components organized by atomic design hierarchy (atoms, molecules, organisms)
- View live component previews
- View HTML source code
- Hierarchical navigation with collapsible categories
- Responsive design using Tailwind CSS
- Alpine.js for interactive features

## Architecture

The book reads component data from the `./docs/pages` directory, which contains:
- `.json` files with component metadata
- `.html` files with rendered component HTML

The server provides:
- `/` - Main book interface
- `/static/` - Static assets (Tailwind CSS, Alpine.js)
- `/api/components` - JSON API for component metadata
- `/api/component/{path}` - HTML content for specific components

## Dependencies

Uses local assets from the `./assets/libs/` directory:
- `tailwind.min.css` - Tailwind CSS framework
- `alpine.3.8.1.min.js` - Alpine.js for reactivity