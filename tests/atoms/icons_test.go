package atoms

import (
	"testing"

	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
)

func TestIcons(t *testing.T) {
	kit := kitcla.New()

	// Create a showcase of all icon usage
	title := kit.Component.Ccv("h1", "text-3xl font-bold text-gray-900 mb-8", "Icon Components Overview")

	// Icon examples section
	iconSection := kit.Component.Dcs("space-y-6",
		kit.Component.Ccv("h2", "text-xl font-semibold text-gray-800 mb-4", "Common Icons"),
		kit.Component.Dcs("grid grid-cols-1 md:grid-cols-3 gap-6",
			kit.Component.Dcs("flex flex-col items-center p-6 border rounded-lg bg-gray-50",
				kit.Atoms.Icons.Icon.Icon("tree"),
				kit.Component.Ccv("span", "text-sm text-gray-600 mt-3 font-medium", "Tree Icon"),
				kit.Component.Ccv("span", "text-xs text-gray-500", "Available icon example"),
			),
			kit.Component.Dcs("flex flex-col items-center p-6 border rounded-lg bg-gray-100",
				kit.Component.Ccv("div", "w-6 h-6 bg-gray-300 rounded border-2 border-dashed border-gray-400"),
				kit.Component.Ccv("span", "text-sm text-gray-600 mt-3 font-medium", "Custom Icon"),
				kit.Component.Ccv("span", "text-xs text-gray-500", "Placeholder for custom icons"),
			),
			kit.Component.Dcs("flex flex-col items-center p-6 border rounded-lg bg-gray-100",
				kit.Component.Ccv("div", "w-6 h-6 bg-gray-300 rounded border-2 border-dashed border-gray-400"),
				kit.Component.Ccv("span", "text-sm text-gray-600 mt-3 font-medium", "More Icons"),
				kit.Component.Ccv("span", "text-xs text-gray-500", "Additional icons can be added"),
			),
		),
	)

	// Usage note
	noteSection := kit.Component.Dcs("bg-blue-50 border border-blue-200 rounded-lg p-4",
		kit.Component.Ccv("h3", "text-lg font-semibold text-blue-800 mb-2", "Usage Note"),
		kit.Component.Ccv("p", "text-blue-700", "Icons are loaded from the icon library. Available icons depend on your configured icon set. The examples above show common icon names that are typically available."),
	)

	// Main container
	container := kit.Component.Dcs("space-y-8 p-6", title, iconSection, noteSection)

	// Render and save
	html := goc.RenderRoot(container)
	sup.AddIndexPage("icons", html)
}
