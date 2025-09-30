package gardening

import (
	"testing"

	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/components/atoms/links"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
)

func TestShowPlant(t *testing.T) {
	kit := kitcla.New()

	// Navigation breadcrumb
	breadcrumb := kit.Component.Dcs("bg-white border-b px-6 py-3",
		kit.Component.Dcs("flex items-center space-x-2 text-sm",
			kit.Atoms.Links.Link.H(&links.LinkMod{Label: "Garden", Url: "/garden", HoverCss: "text-blue-600 hover:text-blue-800"}),
			kit.Component.Dcv("text-gray-400", "›"),
			kit.Atoms.Links.Link.H(&links.LinkMod{Label: "Plants", Url: "/plants", HoverCss: "text-blue-600 hover:text-blue-800"}),
			kit.Component.Dcv("text-gray-400", "›"),
			kit.Component.Dcv("text-gray-600", "Monstera Deliciosa"),
		),
	)

	// Plant header with actions
	plantHeader := kit.Component.Dcs("bg-white px-6 py-6 border-b",
		kit.Component.Dcs("flex justify-between items-start",
			kit.Component.Ds(
				kit.Component.Ccv("h1", "text-3xl font-bold text-gray-800 mb-2", "🌿 Monstera Deliciosa"),
				kit.Component.Dcv("text-gray-600", "Added on March 15, 2024 • Last watered 2 days ago"),
			),
			kit.Component.Dcs("flex space-x-3",
				kit.Atoms.Buttons.Button.PrimaryLink("Edit Plant", "/plants/1/edit", nil),
				kit.Atoms.Buttons.Button.SecondaryLink("Add Photo", "/plants/1/photos", nil),
				kit.Atoms.Buttons.Button.SecondaryLink("Delete", "/plants/1/delete", nil),
			),
		),
	)

	// Plant status card
	statusCard := kit.Component.Dcs("mb-6",
		kit.Organisms.Cards.Card.SimpleCard(
			"Plant Status",
			kit.Component.Ds(
				kit.Component.Dcs("grid grid-cols-2 gap-4 mb-4",
					kit.Component.Ds(
						kit.Component.Dcv("text-sm font-medium text-gray-500", "Health Status"),
						kit.Component.Dcv("text-green-600 font-medium", "Healthy"),
					),
					kit.Component.Ds(
						kit.Component.Dcv("text-sm font-medium text-gray-500", "Watering"),
						kit.Component.Dcv("text-yellow-600 font-medium", "Due Soon"),
					),
					kit.Component.Ds(
						kit.Component.Dcv("text-sm font-medium text-gray-500", "Light Level"),
						kit.Component.Dcv("text-gray-800", "Bright Indirect"),
					),
					kit.Component.Ds(
						kit.Component.Dcv("text-sm font-medium text-gray-500", "Humidity"),
						kit.Component.Dcv("text-gray-800", "60%"),
					),
				),
				kit.Molecules.Alerts.Alert.WarningAlert(
					"Watering Reminder",
					"This plant typically needs water every 7-10 days. Last watered 2 days ago.",
					nil,
				),
			),
		),
	)

	// Plant details grid
	detailsGrid := kit.Component.Dcs("bg-white p-6 rounded-lg shadow mb-6",
		kit.Component.Ccv("h3", "text-lg font-semibold text-gray-800 mb-4", "Plant Details"),
		kit.Component.Dcs("space-y-3",
			kit.Component.Dcs("flex justify-between py-2 border-b border-gray-100",
				kit.Component.Dcv("font-medium text-gray-500", "Common Name"),
				kit.Component.Dcv("text-gray-800", "Swiss Cheese Plant"),
			),
			kit.Component.Dcs("flex justify-between py-2 border-b border-gray-100",
				kit.Component.Dcv("font-medium text-gray-500", "Scientific Name"),
				kit.Component.Dcv("text-gray-800", "Monstera deliciosa"),
			),
			kit.Component.Dcs("flex justify-between py-2 border-b border-gray-100",
				kit.Component.Dcv("font-medium text-gray-500", "Plant Type"),
				kit.Component.Dcv("text-gray-800", "Tropical Houseplant"),
			),
			kit.Component.Dcs("flex justify-between py-2 border-b border-gray-100",
				kit.Component.Dcv("font-medium text-gray-500", "Location"),
				kit.Component.Dcv("text-gray-800", "Living Room - East Window"),
			),
			kit.Component.Dcs("flex justify-between py-2 border-b border-gray-100",
				kit.Component.Dcv("font-medium text-gray-500", "Pot Size"),
				kit.Component.Dcv("text-gray-800", "12 inches"),
			),
			kit.Component.Dcs("flex justify-between py-2 border-b border-gray-100",
				kit.Component.Dcv("font-medium text-gray-500", "Soil Type"),
				kit.Component.Dcv("text-gray-800", "Well-draining potting mix"),
			),
			kit.Component.Dcs("flex justify-between py-2 border-b border-gray-100",
				kit.Component.Dcv("font-medium text-gray-500", "Fertilizer"),
				kit.Component.Dcv("text-gray-800", "Monthly liquid fertilizer"),
			),
			kit.Component.Dcs("flex justify-between py-2",
				kit.Component.Dcv("font-medium text-gray-500", "Growth Rate"),
				kit.Component.Dcv("text-gray-800", "Fast"),
			),
		),
	)

	// Care history table
	careHistoryCard := kit.Component.Dcs("mb-6",
		kit.Organisms.Cards.Card.SimpleCard(
			"Recent Care History",
			kit.Component.Dcs("overflow-x-auto",
				kit.Component.Cas("table", goc.Attr{"class": "min-w-full divide-y divide-gray-200"},
					kit.Component.Cas("thead", goc.Attr{"class": "bg-gray-50"},
						kit.Component.Cas("tr", nil,
							kit.Component.Cav("th", goc.Attr{"class": "px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"}, "Date"),
							kit.Component.Cav("th", goc.Attr{"class": "px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"}, "Action"),
							kit.Component.Cav("th", goc.Attr{"class": "px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"}, "Notes"),
							kit.Component.Cav("th", goc.Attr{"class": "px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"}, "Next Due"),
						),
					),
					kit.Component.Cas("tbody", goc.Attr{"class": "bg-white divide-y divide-gray-200"},
						kit.Component.Cs("tr",
							kit.Component.Cav("td", goc.Attr{"class": "px-6 py-4 whitespace-nowrap text-sm text-gray-900"}, "Mar 20, 2024"),
							kit.Component.Cas("td", goc.Attr{"class": "px-6 py-4 whitespace-nowrap"}, kit.Component.Dcv("inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-blue-100 text-blue-800", "Watered")),
							kit.Component.Cav("td", goc.Attr{"class": "px-6 py-4 text-sm text-gray-900"}, "Soil was dry 2 inches down"),
							kit.Component.Cav("td", goc.Attr{"class": "px-6 py-4 whitespace-nowrap text-sm text-gray-900"}, "Mar 27-30"),
						),
						kit.Component.Cs("tr",
							kit.Component.Cav("td", goc.Attr{"class": "px-6 py-4 whitespace-nowrap text-sm text-gray-900"}, "Mar 15, 2024"),
							kit.Component.Cas("td", goc.Attr{"class": "px-6 py-4 whitespace-nowrap"}, kit.Component.Dcv("inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-green-100 text-green-800", "Fertilized")),
							kit.Component.Cav("td", goc.Attr{"class": "px-6 py-4 text-sm text-gray-900"}, "Monthly feeding with liquid fertilizer"),
							kit.Component.Cav("td", goc.Attr{"class": "px-6 py-4 whitespace-nowrap text-sm text-gray-900"}, "Apr 15"),
						),
						kit.Component.Cs("tr",
							kit.Component.Cav("td", goc.Attr{"class": "px-6 py-4 whitespace-nowrap text-sm text-gray-900"}, "Mar 10, 2024"),
							kit.Component.Cas("td", goc.Attr{"class": "px-6 py-4 whitespace-nowrap"}, kit.Component.Dcv("inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-yellow-100 text-yellow-800", "Pruned")),
							kit.Component.Cav("td", goc.Attr{"class": "px-6 py-4 text-sm text-gray-900"}, "Removed yellowing leaves"),
							kit.Component.Cav("td", goc.Attr{"class": "px-6 py-4 whitespace-nowrap text-sm text-gray-900"}, "As needed"),
						),
					),
				),
			),
		),
	)

	// Notes section
	notesCard := kit.Organisms.Cards.Card.SimpleCard(
		"Plant Notes",
		kit.Component.Ds(
			kit.Component.Dcv("text-gray-700 leading-relaxed mb-4", "This Monstera has been thriving since I brought it home. The large fenestrated leaves are beautiful and it's growing quite quickly. I've noticed it responds well to the bright indirect light from the east window."),
			kit.Component.Dcv("text-gray-700 leading-relaxed mb-4", "Recent observation: New aerial roots are developing, which is a good sign of healthy growth."),
			kit.Component.Dcs("mt-4",
				kit.Atoms.Buttons.Button.SecondaryLink("Edit Notes", "/plants/1/notes", nil),
			),
		),
	)

	// Main layout
	mainContent := kit.Component.Dcs("min-h-screen bg-gray-50",
		breadcrumb,
		plantHeader,
		kit.Component.Dcs("max-w-6xl mx-auto px-6 py-8",
			kit.Component.Dcs("grid grid-cols-1 lg:grid-cols-3 gap-6",
				kit.Component.Dcs("lg:col-span-2",
					detailsGrid,
					careHistoryCard,
					notesCard,
				),
				kit.Component.Ds(
					statusCard,
				),
			),
		),
	)

	html := goc.RenderRoot(mainContent)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/show_plant_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/show_plant_1.html")
	sup.AddPage("ShowPlant", html)
}
