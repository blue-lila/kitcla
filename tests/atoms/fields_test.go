package atoms

import (
	"testing"

	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
)

func TestFields(t *testing.T) {
	kit := kitcla.New()

	// Create a showcase of field components with gardening theme
	title := kit.Component.Ccv("h1", "text-3xl font-bold text-gray-900 mb-8", "Garden Field Components")

	// Create input components using proper API
	usernameInput := kit.Component.Ca("input", goc.Attr{
		"type":        "text",
		"class":       "w-full px-3 py-2 border border-gray-300 rounded-md",
		"placeholder": "Enter gardener username",
	})

	emailInput := kit.Component.Ca("input", goc.Attr{
		"type":        "email",
		"class":       "w-full px-3 py-2 border border-gray-300 rounded-md",
		"placeholder": "gardener@example.com",
	})

	descriptionInput := kit.Component.Ccv("textarea", "w-full px-3 py-2 border border-gray-300 rounded-md", "This is my vegetable garden with tomatoes, peppers, and herbs.")

	passwordInput := kit.Component.Ca("input", goc.Attr{
		"type":  "password",
		"class": "w-full px-3 py-2 border border-gray-300 rounded-md",
	})

	// Basic field section
	basicSection := kit.Component.Dcs("space-y-4",
		kit.Component.Ccv("h2", "text-xl font-semibold text-gray-800 mb-4", "Basic Garden Fields"),
		kit.Component.Dcs("space-y-6",
			kit.Atoms.Fields.Field.Field("Gardener Name", usernameInput),
			kit.Atoms.Fields.Field.Field("Garden Email", emailInput),
			kit.Atoms.Fields.Field.Field("Garden Description", descriptionInput),
		),
	)

	// Field variations section
	variationsSection := kit.Component.Dcs("space-y-4",
		kit.Component.Ccv("h2", "text-xl font-semibold text-gray-800 mb-4", "Field Variations"),
		kit.Component.Dcs("space-y-6",
			kit.Atoms.Fields.Field.Field("Garden Password", passwordInput),
			kit.Atoms.Fields.Field.HiddenField("Garden ID", kit.Component.Ca("input", goc.Attr{"type": "hidden", "value": "garden_123"})),
		),
	)

	// Main container
	container := kit.Component.Dcs("space-y-8 p-6", title, basicSection, variationsSection)

	// Render and save
	html := goc.RenderRoot(container)
	sup.AddIndexPage("fields", html)
}
