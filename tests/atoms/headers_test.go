package atoms

import (
	"testing"

	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
)

func TestHeaders(t *testing.T) {
	kit := kitcla.New()

	// Create a showcase of all header types
	title := kit.Component.Ccv("h1", "text-3xl font-bold text-gray-900 mb-8", "Header Components Overview")

	// Header hierarchy section
	headerSection := kit.Component.Dcs("space-y-6",
		kit.Component.Ccv("h2", "text-xl font-semibold text-gray-800 mb-4", "Header Hierarchy"),
		kit.Component.Dcs("space-y-4",
			kit.Component.Dcs("border-l-4 border-blue-500 pl-4",
				kit.Atoms.Headers.Header.H1("This is an H1 Header"),
			),
			kit.Component.Dcs("border-l-4 border-green-500 pl-4",
				kit.Atoms.Headers.Header.H2("This is an H2 Header"),
			),
			kit.Component.Dcs("border-l-4 border-yellow-500 pl-4",
				kit.Component.Ccv("h3", "text-lg font-semibold text-gray-800", "This is an H3 Header"),
			),
			kit.Component.Dcs("border-l-4 border-purple-500 pl-4",
				kit.Component.Ccv("h4", "text-base font-semibold text-gray-800", "This is an H4 Header"),
			),
		),
	)

	// Usage examples section
	examplesSection := kit.Component.Dcs("space-y-6",
		kit.Component.Ccv("h2", "text-xl font-semibold text-gray-800 mb-4", "Usage Examples"),
		kit.Component.Dcs("grid grid-cols-1 md:grid-cols-2 gap-6",
			kit.Component.Dcs("p-4 bg-gray-50 rounded-lg",
				kit.Atoms.Headers.Header.H2("Page Title"),
				kit.Component.Ccv("p", "text-gray-600 text-sm mt-2", "Used for main page titles"),
			),
			kit.Component.Dcs("p-4 bg-gray-50 rounded-lg",
				kit.Component.Ccv("h3", "text-lg font-semibold text-gray-800", "Section Header"),
				kit.Component.Ccv("p", "text-gray-600 text-sm mt-2", "Used for section divisions"),
			),
		),
	)

	// Main container
	container := kit.Component.Dcs("space-y-8 p-6", title, headerSection, examplesSection)

	// Render and save
	html := goc.RenderRoot(container)
	sup.AddIndexPage("headers", html)
}
