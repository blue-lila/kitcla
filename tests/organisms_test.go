package tests

import (
	"testing"

	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
)

func TestOrganismsIndex(t *testing.T) {
	kit := kitcla.New()

	// Title
	title := kit.Component.Ccv("h1", "text-3xl font-bold text-gray-900 mb-8", "Organisms Overview")

	// Description
	description := kit.Component.Dcs("prose max-w-none mb-8",
		kit.Component.Ccv("p", "text-lg text-gray-600", "Organisms are complex UI components composed of groups of molecules and atoms. They form distinct sections of an interface."),
	)

	// Component cards
	cards := []goc.HTML{
		createCard(kit, "Cards", "/?page=organisms/cards/index", "Card components for content grouping and display."),
		createCard(kit, "Duo Cards", "/?page=organisms/duo_cards/index", "Dual-panel card components with split content views."),
		createCard(kit, "Forms", "/?page=organisms/forms/index", "Complete form components with validation and submission."),
		createCard(kit, "Grids", "/?page=organisms/grids/index", "Grid layout components for structured content display."),
		createCard(kit, "Key Point", "/?page=organisms/key_point/index", "Key point highlight components for important information."),
		createCard(kit, "Logs", "/?page=organisms/logs/index", "Log display components for activity and event tracking."),
		createCard(kit, "Menus", "/?page=organisms/menus/index", "Menu components for navigation and actions."),
		createCard(kit, "Modals", "/?page=organisms/modals/index", "Modal dialog components for focused interactions."),
		createCard(kit, "Navbars", "/?page=organisms/navbars/index", "Complex navigation bar components with multiple features."),
		createCard(kit, "Navmenus", "/?page=organisms/navmenus/index", "Navigation menu components for hierarchical navigation."),
		createCard(kit, "Sidebars", "/?page=organisms/sidebars/index", "Sidebar components for secondary navigation and content."),
		createCard(kit, "Tables", "/?page=organisms/tables/index", "Table components for structured data display and interaction."),
	}

	// Grid
	grid := kit.Component.Dcs("grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4", cards...)

	// Main container
	container := kit.Component.Dcs("space-y-8 p-6", title, description, grid)

	// Render and save
	html := goc.RenderRoot(container)
	sup.AddCustomPage(html, "docs/pages/organisms/index.html", "docs/pages/organisms/index.json")
}
