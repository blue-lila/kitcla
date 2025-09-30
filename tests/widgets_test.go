package tests

import (
	"testing"

	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
)

func TestWidgetsIndex(t *testing.T) {
	kit := kitcla.New()

	// Title
	title := kit.Component.Ccv("h1", "text-3xl font-bold text-gray-900 mb-8", "Widgets Overview")

	// Description
	description := kit.Component.Dcs("prose max-w-none mb-8",
		kit.Component.Ccv("p", "text-lg text-gray-600", "Widgets are complete, self-contained application features that combine organisms, molecules, and atoms into full user experiences."),
	)

	// Component cards
	cards := []goc.HTML{
		createCard(kit, "Gardening", "/?page=widgets/gardening/index", "Complete gardening management widget with plant tracking, care schedules, and harvest logging."),
	}

	// Grid
	grid := kit.Component.Dcs("grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4", cards...)

	// Main container
	container := kit.Component.Dcs("space-y-8 p-6", title, description, grid)

	// Render and save
	html := goc.RenderRoot(container)
	sup.AddCustomPage(html, "docs/pages/widgets/index.html", "docs/pages/widgets/index.json")
}
