package tests

import (
	"testing"

	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
)

func TestAtomsIndex(t *testing.T) {
	kit := kitcla.New()

	// Title
	title := kit.Component.Ccv("h1", "text-3xl font-bold text-gray-900 mb-8", "Atoms Overview")

	// Description
	description := kit.Component.Dcs("prose max-w-none mb-8",
		kit.Component.Ccv("p", "text-lg text-gray-600", "Atoms are the foundational building blocks of the KitCLA design system. These are the smallest possible components that can't be broken down any further - buttons, inputs, icons, and other basic HTML elements with specific styling and behavior."),
	)

	// Component cards
	cards := []goc.HTML{
		createCard(kit, "Buttons", "/?page=atoms/buttons/index", "Interactive elements for user actions including primary, secondary, and Alpine.js variants."),
		createCard(kit, "Inputs", "/?page=atoms/inputs/index", "Form input elements including text, number, select, file upload, and specialized inputs."),
		createCard(kit, "Cells", "/?page=atoms/cells/index", "Data display components for tables and grids including text, pills, booleans, and time cells."),
		createCard(kit, "Icons", "/?page=atoms/icons/index", "SVG icon components with consistent sizing and styling options."),
		createCard(kit, "Headers", "/?page=atoms/headers/index", "Typography components for titles, headings, and section headers."),
		createCard(kit, "Links", "/?page=atoms/links/index", "Clickable text and navigation elements with hover states and styling."),
	}

	// Grid
	grid := kit.Component.Dcs("grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4", cards...)

	// Main container
	container := kit.Component.Dcs("space-y-8 p-6", title, description, grid)

	// Render and save
	html := goc.RenderRoot(container)
	sup.AddCustomPage(html, "docs/pages/atoms/index.html", "docs/pages/atoms/index.json")
}

func createCard(kit *kitcla.Kit, title, href, description string) goc.HTML {
	// Icon with colored background
	iconBg := kit.Component.Dcs("w-10 h-10 bg-blue-100 rounded-lg flex items-center justify-center mr-3",
		kit.Component.Cas("svg", goc.Attr{"class": "w-5 h-5 text-blue-600", "fill": "none", "stroke": "currentColor", "viewBox": "0 0 24 24"},
			kit.Component.Cas("path", goc.Attr{"stroke-linecap": "round", "stroke-linejoin": "round", "stroke-width": "2", "d": "M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10"}),
		),
	)

	// Title and subtitle
	titleSection := kit.Component.Dcs("",
		kit.Component.Ccv("h3", "font-semibold text-gray-900", title),
	)

	// Header with icon and title
	header := kit.Component.Dcs("flex items-center mb-3", iconBg, titleSection)

	// Description
	desc := kit.Component.Ccv("p", "text-sm text-gray-600", description)

	// Card container
	return kit.Component.Cas("a", goc.Attr{"href": href, "class": "block p-6 bg-white rounded-lg border border-gray-200 hover:border-blue-300 hover:shadow-md transition-all duration-200"},
		header,
		desc,
	)
}
