package gardening

import (
	"testing"

	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
)

func TestGardenOverview(t *testing.T) {
	kit := kitcla.New()

	// Header section
	header := kit.Component.Dcs("bg-green-50 p-6",
		kit.Component.Ccv("h1", "text-3xl font-bold text-green-800 mb-2", "🌱 My Garden Overview"),
		kit.Component.Dcv("text-green-600", "Welcome to your digital garden dashboard"),
	)

	// Quick stats cards
	statsGrid := kit.Component.Dcs("grid grid-cols-2 lg:grid-cols-4 gap-4 mb-8",
		kit.Component.Dcs("bg-white p-4 rounded-lg shadow border-l-4 border-green-400",
			kit.Component.Dcv("text-2xl mb-2", "🌿"),
			kit.Component.Dcv("text-2xl font-bold text-gray-800", "24"),
			kit.Component.Dcv("text-sm text-gray-500", "Total Plants"),
		),
		kit.Component.Dcs("bg-white p-4 rounded-lg shadow border-l-4 border-blue-400",
			kit.Component.Dcv("text-2xl mb-2", "💧"),
			kit.Component.Dcv("text-2xl font-bold text-gray-800", "8"),
			kit.Component.Dcv("text-sm text-gray-500", "Need Watering"),
		),
		kit.Component.Dcs("bg-white p-4 rounded-lg shadow border-l-4 border-yellow-400",
			kit.Component.Dcv("text-2xl mb-2", "🌞"),
			kit.Component.Dcv("text-2xl font-bold text-gray-800", "3"),
			kit.Component.Dcv("text-sm text-gray-500", "In Bloom"),
		),
		kit.Component.Dcs("bg-white p-4 rounded-lg shadow border-l-4 border-red-400",
			kit.Component.Dcv("text-2xl mb-2", "⚠️"),
			kit.Component.Dcv("text-2xl font-bold text-gray-800", "2"),
			kit.Component.Dcv("text-sm text-gray-500", "Need Attention"),
		),
	)

	// Recent activities alert
	recentAlert := kit.Component.Dcs("mb-6",
		kit.Molecules.Alerts.Alert.SuccessAlert(
			"Recent Activity",
			"Watered tomatoes and basil yesterday. Fertilized roses last week.",
			nil,
		),
	)

	// Quick actions buttons
	quickActions := kit.Component.Dcs("bg-white p-6 rounded-lg shadow mb-8",
		kit.Component.Ccv("h2", "text-xl font-semibold text-gray-800 mb-4", "Quick Actions"),
		kit.Component.Dcs("flex flex-wrap gap-3",
			kit.Atoms.Buttons.Button.PrimaryLink("View All Plants", "/plants", nil),
			kit.Atoms.Buttons.Button.SecondaryLink("Add New Plant", "/plants/new", nil),
			kit.Atoms.Buttons.Button.SecondaryLink("Watering Schedule", "/schedule", nil),
			kit.Atoms.Buttons.Button.SecondaryLink("Garden Journal", "/journal", nil),
		),
	)

	// Upcoming tasks
	upcomingTasks := kit.Component.Dcs("bg-white p-6 rounded-lg shadow",
		kit.Component.Ccv("h2", "text-xl font-semibold text-gray-800 mb-4", "Upcoming Tasks"),
		kit.Component.Dcs("space-y-3",
			kit.Component.Dcs("flex items-center p-3 bg-blue-50 rounded-md",
				kit.Component.Dcv("text-blue-600 text-lg mr-3", "💧"),
				kit.Component.Ds(
					kit.Component.Dcv("font-medium text-gray-800", "Water the succulents"),
					kit.Component.Dcv("text-sm text-blue-600", "Today • High Priority"),
				),
			),
			kit.Component.Dcs("flex items-center p-3 bg-yellow-50 rounded-md",
				kit.Component.Dcv("text-yellow-600 text-lg mr-3", "🌞"),
				kit.Component.Ds(
					kit.Component.Dcv("font-medium text-gray-800", "Check tomato growth"),
					kit.Component.Dcv("text-sm text-yellow-600", "Tomorrow • Medium Priority"),
				),
			),
			kit.Component.Dcs("flex items-center p-3 bg-green-50 rounded-md",
				kit.Component.Dcv("text-green-600 text-lg mr-3", "🌿"),
				kit.Component.Ds(
					kit.Component.Dcv("font-medium text-gray-800", "Prune the roses"),
					kit.Component.Dcv("text-sm text-green-600", "This weekend • Low Priority"),
				),
			),
		),
	)

	// Main layout
	mainContent := kit.Component.Dcs("min-h-screen bg-gray-50",
		header,
		kit.Component.Dcs("max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 pb-8",
			statsGrid,
			recentAlert,
			quickActions,
			upcomingTasks,
		),
	)

	html := goc.RenderRoot(mainContent)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/garden_overview_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/garden_overview_1.html")
	sup.AddPage("GardenOverview", html)
}
