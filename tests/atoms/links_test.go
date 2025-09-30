package atoms

import (
	"testing"

	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
)

func TestLinks(t *testing.T) {
	kit := kitcla.New()

	// Create a showcase of all link types with gardening theme
	title := kit.Component.Ccv("h1", "text-3xl font-bold text-gray-900 mb-8", "Garden Link Components")

	// Basic links section
	basicSection := kit.Component.Dcs("space-y-4",
		kit.Component.Ccv("h2", "text-xl font-semibold text-gray-800 mb-4", "Basic Garden Links"),
		kit.Component.Dcs("space-y-3",
			kit.Component.Dcs("flex items-center space-x-4",
				kit.Atoms.Links.Link.Link("Plant Library", "/plants"),
				kit.Component.Ccv("span", "text-sm text-gray-500", "Browse plant varieties"),
			),
			kit.Component.Dcs("flex items-center space-x-4",
				kit.Atoms.Links.Link.Link("Garden Tools", "/tools"),
				kit.Component.Ccv("span", "text-sm text-gray-500", "Essential gardening equipment"),
			),
			kit.Component.Dcs("flex items-center space-x-4",
				kit.Atoms.Links.Link.IconLink("trowel", "/tools"),
				kit.Component.Ccv("span", "text-sm text-gray-500", "Icon link to tools"),
			),
		),
	)

	// Link types section
	typesSection := kit.Component.Dcs("space-y-4",
		kit.Component.Ccv("h2", "text-xl font-semibold text-gray-800 mb-4", "Garden Link Types"),
		kit.Component.Dcs("grid grid-cols-1 md:grid-cols-2 gap-6",
			kit.Component.Dcs("p-4 bg-gray-50 rounded-lg",
				kit.Component.Ccv("h3", "text-sm font-medium text-gray-700 mb-3", "Standard Links"),
				kit.Atoms.Links.Link.Link("View Garden Plan", "/garden-plan"),
			),
			kit.Component.Dcs("p-4 bg-gray-50 rounded-lg",
				kit.Component.Ccv("h3", "text-sm font-medium text-gray-700 mb-3", "Submit Actions"),
				kit.Atoms.Links.Link.SubmitLink("Save Garden Changes"),
			),
		),
	)

	// Usage examples section
	usageSection := kit.Component.Dcs("space-y-4",
		kit.Component.Ccv("h2", "text-xl font-semibold text-gray-800 mb-4", "Garden Navigation Examples"),
		kit.Component.Dcs("space-y-6",
			kit.Component.Dcs("space-y-2",
				kit.Component.Ccv("h3", "text-lg font-medium text-gray-800", "Garden Navigation"),
				kit.Component.Dcs("flex space-x-6",
					kit.Atoms.Links.Link.Link("Dashboard", "/dashboard"),
					kit.Atoms.Links.Link.Link("My Plants", "/plants"),
					kit.Atoms.Links.Link.Link("Care Schedule", "/schedule"),
				),
			),
			kit.Component.Dcs("space-y-2",
				kit.Component.Ccv("h3", "text-lg font-medium text-gray-800", "In Garden Content"),
				kit.Component.Dcs("p text-gray-700",
					kit.Component.Ccv("span", "", "Visit our "),
					kit.Atoms.Links.Link.Link("planting guide", "/guide"),
					kit.Component.Ccv("span", "", " for seasonal tips."),
				),
			),
		),
	)

	// Main container
	container := kit.Component.Dcs("space-y-8 p-6", title, basicSection, typesSection, usageSection)

	// Render and save
	html := goc.RenderRoot(container)
	sup.AddIndexPage("links", html)
}
