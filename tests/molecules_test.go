package tests

import (
	"testing"

	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
)

func TestMoleculesIndex(t *testing.T) {
	kit := kitcla.New()

	// Title
	title := kit.Component.Ccv("h1", "text-3xl font-bold text-gray-900 mb-8", "Molecules Overview")

	// Description
	description := kit.Component.Dcs("prose max-w-none mb-8",
		kit.Component.Ccv("p", "text-lg text-gray-600", "Molecules are simple combinations of atoms that work together as a unit. They represent relatively simple groups of UI elements functioning together."),
	)

	// Component cards
	cards := []goc.HTML{
		createCard(kit, "Alerts", "/?page=molecules/alerts/index", "Alert and notification components with icons and styling variants."),
		createCard(kit, "Messages", "/?page=molecules/messages/index", "Message display components for user feedback and communication."),
		createCard(kit, "Navbars", "/?page=molecules/navbars/index", "Navigation bar components for site-wide navigation."),
		createCard(kit, "Popovers", "/?page=molecules/popovers/index", "Popover components for contextual information and actions."),
		createCard(kit, "Steppers", "/?page=molecules/steppers/index", "Step indicator components for multi-step processes."),
		createCard(kit, "Trees", "/?page=molecules/trees/index", "Tree view components for hierarchical data display."),
	}

	// Grid
	grid := kit.Component.Dcs("grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4", cards...)

	// Main container
	container := kit.Component.Dcs("space-y-8 p-6", title, description, grid)

	// Render and save
	html := goc.RenderRoot(container)
	sup.AddCustomPage(html, "docs/pages/molecules/index.html", "docs/pages/molecules/index.json")
}
