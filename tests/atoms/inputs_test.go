package atoms

import (
	"testing"

	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
)

func TestInputs(t *testing.T) {
	kit := kitcla.New()

	// Create a showcase of all input types
	title := kit.Component.Ccv("h1", "text-3xl font-bold text-gray-900 mb-8", "Input Components Overview")

	// Basic inputs section
	basicSection := kit.Component.Dcs("space-y-4",
		kit.Component.Ccv("h2", "text-xl font-semibold text-gray-800 mb-4", "Basic Inputs"),
		kit.Component.Dcs("grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6",
			kit.Component.Dcs("space-y-2",
				kit.Component.Ccv("label", "text-sm font-medium text-gray-700", "Text Input"),
				kit.Atoms.Inputs.TextInput.TextInput("Sample text", "text-input-1"),
			),
			kit.Component.Dcs("space-y-2",
				kit.Component.Ccv("label", "text-sm font-medium text-gray-700", "File Input"),
				kit.Atoms.Inputs.FileInput.FileInput("file-input-1"),
			),
			kit.Component.Dcs("space-y-2",
				kit.Component.Ccv("label", "text-sm font-medium text-gray-700", "Hidden Input"),
				kit.Atoms.Inputs.HiddenInput.HiddenInput("hidden-value", "hidden-input-1"),
			),
		),
	)

	// Usage note
	noteSection := kit.Component.Dcs("bg-blue-50 border border-blue-200 rounded-lg p-4",
		kit.Component.Ccv("h3", "text-lg font-semibold text-blue-800 mb-2", "Usage Note"),
		kit.Component.Ccv("p", "text-blue-700", "Input components provide the foundation for form interactions. Each input type has specific parameters and validation rules."),
	)

	// Main container
	container := kit.Component.Dcs("space-y-8 p-6", title, basicSection, noteSection)

	// Render and save
	html := goc.RenderRoot(container)
	sup.AddIndexPage("inputs", html)
}
