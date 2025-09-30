package atoms

import (
	"testing"

	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/components/atoms/tabs"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
)

func TestTabs(t *testing.T) {
	kit := kitcla.New()

	// Create a showcase of tab components with gardening theme
	title := kit.Component.Ccv("h1", "text-3xl font-bold text-gray-900 mb-8", "Garden Tab Components")

	// Helper function to create tab items
	createTabs := func(tabData []struct {
		text   string
		active bool
	}) goc.HTML { var items []*tabs.TabItem; for _, tab := range tabData {
		items = append(items, &tabs.TabItem{
			Text:   tab.text,
			Value:  tab.text,
			Active: tab.active,
		})
	}; return kit.Atoms.Tabs.Tab.Tab(&tabs.TabMod{Items: items}) }

	// Basic tabs section
	basicSection := kit.Component.Dcs("space-y-4",
		kit.Component.Ccv("h2", "text-xl font-semibold text-gray-800 mb-4", "Garden Management Tabs"),
		kit.Component.Dcs("border rounded-lg p-4 bg-white",
			createTabs([]struct {
				text   string
				active bool
			}{
				{"My Plants", true},
				{"Care Schedule", false},
				{"Garden Tools", false},
			}),
		),
	)

	// Tab variations section
	variationsSection := kit.Component.Dcs("space-y-4",
		kit.Component.Ccv("h2", "text-xl font-semibold text-gray-800 mb-4", "Garden Tab States"),
		kit.Component.Dcs("grid grid-cols-1 gap-6",
			kit.Component.Dcs("border rounded-lg p-4 bg-white",
				kit.Component.Ccv("h3", "text-sm font-medium text-gray-700 mb-3", "Active Garden Tab"),
				createTabs([]struct {
					text   string
					active bool
				}{
					{"Vegetable Garden", true},
				}),
			),
			kit.Component.Dcs("border rounded-lg p-4 bg-white",
				kit.Component.Ccv("h3", "text-sm font-medium text-gray-700 mb-3", "Multiple Garden Sections"),
				createTabs([]struct {
					text   string
					active bool
				}{
					{"Vegetables", false},
					{"Herbs", true},
					{"Flowers", false},
				}),
			),
		),
	)

	// Usage examples section
	usageSection := kit.Component.Dcs("space-y-4",
		kit.Component.Ccv("h2", "text-xl font-semibold text-gray-800 mb-4", "Garden Navigation Examples"),
		kit.Component.Dcs("bg-white border rounded-lg p-6",
			kit.Component.Ccv("h3", "text-lg font-medium text-gray-800 mb-4", "Garden Dashboard Navigation"),
			kit.Component.Dcs("space-y-4",
				createTabs([]struct {
					text   string
					active bool
				}{
					{"Overview", false},
					{"Plant Care", true},
					{"Watering Schedule", false},
					{"Harvest Log", false},
				}),
				kit.Component.Dcs("p-4 bg-green-50 rounded border border-green-200",
					kit.Component.Ccv("h4", "font-medium text-green-800 mb-2", "Plant Care Section"),
					kit.Component.Ccv("p", "text-green-700", "Here you would see detailed plant care instructions, watering reminders, fertilizing schedules, and growth tracking for your garden."),
				),
			),
		),
	)

	// Main container
	container := kit.Component.Dcs("space-y-8 p-6", title, basicSection, variationsSection, usageSection)

	// Render and save
	html := goc.RenderRoot(container)
	sup.AddIndexPage("tabs", html)
}
