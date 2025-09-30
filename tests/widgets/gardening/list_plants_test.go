package gardening

import (
	"testing"

	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/components/atoms/links"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
)

func TestListPlants(t *testing.T) {
	kit := kitcla.New()

	// Navigation breadcrumb
	breadcrumb := kit.Component.Dcs("bg-white border-b px-6 py-4",
		kit.Component.Dcs("max-w-7xl mx-auto flex items-center space-x-2 text-sm",
			kit.Atoms.Links.Link.H(&links.LinkMod{Label: "Garden", Url: "/garden", HoverCss: "text-blue-600 hover:text-blue-800"}),
			kit.Component.Dcv("text-gray-400", "›"),
			kit.Component.Dcv("text-gray-600 font-medium", "Plants"),
		),
	)

	// Page header with search and actions
	pageHeader := kit.Component.Dcs("bg-white border-b",
		kit.Component.Dcs("max-w-7xl mx-auto px-6 py-8",
			kit.Component.Dcs("flex flex-col sm:flex-row justify-between items-start sm:items-center gap-6 mb-8",
				kit.Component.Ds(
					kit.Component.Ccv("h1", "text-3xl font-bold text-gray-800 mb-2", "🌱 My Plants"),
					kit.Component.Dcv("text-gray-600 text-sm", "Manage your plant collection"),
				),
				kit.Atoms.Buttons.Button.PrimaryLink("Add New Plant", "/plants/new", nil),
			),
			kit.Component.Dcs("flex flex-col sm:flex-row gap-4",
				kit.Component.Dcs("flex-1",
					kit.Atoms.Inputs.TextInputAlp.TextInput("search"),
				),
				kit.Component.Dcv("text-sm bg-gray-100 border border-gray-300 rounded-md px-4 py-2 whitespace-nowrap", "Filter: All Plants"),
				kit.Atoms.Buttons.Button.SecondaryLink("Reset Filters", "#", nil),
			),
		),
	)

	// Filter summary
	filterSummary := kit.Component.Dcs("bg-blue-50 border-b",
		kit.Component.Dcs("max-w-7xl mx-auto px-6 py-4",
			kit.Molecules.Alerts.Alert.InfoAlert(
				"Search Results",
				"Showing 12 plants. Use filters above to narrow down results.",
				nil,
			),
		),
	)

	// Plants table
	plantsTable := kit.Component.Dcs("bg-white rounded-lg shadow-sm border border-gray-200 mb-8 overflow-hidden",
		kit.Component.Cas("table", goc.Attr{"class": "min-w-full divide-y divide-gray-200"},
			kit.Component.Cas("thead", goc.Attr{"class": "bg-gray-50"},
				kit.Component.Cs("tr",
					kit.Component.Cav("th", goc.Attr{"class": "px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"}, "Plant"),
					kit.Component.Cav("th", goc.Attr{"class": "px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"}, "Status"),
					kit.Component.Cav("th", goc.Attr{"class": "px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"}, "Last Watered"),
					kit.Component.Cav("th", goc.Attr{"class": "px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"}, "Next Due"),
					kit.Component.Cav("th", goc.Attr{"class": "px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"}, "Location"),
					kit.Component.Cav("th", goc.Attr{"class": "px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"}, "Actions"),
				),
			),
			kit.Component.Cas("tbody", goc.Attr{"class": "bg-white divide-y divide-gray-200"},
				kit.Component.Cs("tr",
					kit.Component.Cas("td", goc.Attr{"class": "px-6 py-4 whitespace-nowrap"},
						kit.Component.Dcs("flex items-center space-x-3",
							kit.Component.Dcv("text-2xl", "🌿"),
							kit.Component.Ds(
								kit.Atoms.Links.Link.H(&links.LinkMod{Label: "Monstera Deliciosa", Url: "/plants/1", HoverCss: "font-medium text-blue-600 hover:text-blue-800"}),
								kit.Component.Dcv("text-sm text-gray-500", "Swiss Cheese Plant"),
							),
						),
					),
					kit.Component.Cas("td", goc.Attr{"class": "px-6 py-4 whitespace-nowrap"}, kit.Component.Dcv("inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-green-100 text-green-800", "Healthy")),
					kit.Component.Cav("td", goc.Attr{"class": "px-6 py-4 whitespace-nowrap text-sm text-gray-900"}, "2 days ago"),
					kit.Component.Cav("td", goc.Attr{"class": "px-6 py-4 whitespace-nowrap text-sm text-gray-900"}, "5-8 days"),
					kit.Component.Cav("td", goc.Attr{"class": "px-6 py-4 whitespace-nowrap text-sm text-gray-900"}, "Living Room"),
					kit.Component.Cas("td", goc.Attr{"class": "px-6 py-4 whitespace-nowrap text-sm text-gray-500"},
						kit.Component.Dcs("flex space-x-2",
							kit.Atoms.Buttons.Button.TertiaryIconLink("pen-to-square", "/plants/1/edit", nil),
							kit.Atoms.Links.Link.H(&links.LinkMod{Label: "View", Url: "/plants/1", HoverCss: "text-blue-600 hover:text-blue-800"}),
						),
					),
				),
				kit.Component.Cs("tr",
					kit.Component.Cas("td", goc.Attr{"class": "px-6 py-4 whitespace-nowrap"},
						kit.Component.Dcs("flex items-center space-x-3",
							kit.Component.Dcv("text-2xl", "🌹"),
							kit.Component.Ds(
								kit.Atoms.Links.Link.H(&links.LinkMod{Label: "Garden Rose", Url: "/plants/2", HoverCss: "font-medium text-blue-600 hover:text-blue-800"}),
								kit.Component.Dcv("text-sm text-gray-500", "Rosa gallica"),
							),
						),
					),
					kit.Component.Cas("td", goc.Attr{"class": "px-6 py-4 whitespace-nowrap"}, kit.Component.Dcv("inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-yellow-100 text-yellow-800", "Needs Water")),
					kit.Component.Cav("td", goc.Attr{"class": "px-6 py-4 whitespace-nowrap text-sm text-gray-900"}, "7 days ago"),
					kit.Component.Cav("td", goc.Attr{"class": "px-6 py-4 whitespace-nowrap text-sm text-gray-900"}, "Today"),
					kit.Component.Cav("td", goc.Attr{"class": "px-6 py-4 whitespace-nowrap text-sm text-gray-900"}, "Back Garden"),
					kit.Component.Cas("td", goc.Attr{"class": "px-6 py-4 whitespace-nowrap text-sm text-gray-500"},
						kit.Component.Dcs("flex space-x-2",
							kit.Atoms.Buttons.Button.TertiaryIconLink("pen-to-square", "/plants/2/edit", nil),
							kit.Atoms.Links.Link.H(&links.LinkMod{Label: "View", Url: "/plants/2", HoverCss: "text-blue-600 hover:text-blue-800"}),
						),
					),
				),
				kit.Component.Cs("tr",
					kit.Component.Cas("td", goc.Attr{"class": "px-6 py-4 whitespace-nowrap"},
						kit.Component.Dcs("flex items-center space-x-3",
							kit.Component.Dcv("text-2xl", "🌵"),
							kit.Component.Ds(
								kit.Atoms.Links.Link.H(&links.LinkMod{Label: "Jade Plant", Url: "/plants/3", HoverCss: "font-medium text-blue-600 hover:text-blue-800"}),
								kit.Component.Dcv("text-sm text-gray-500", "Crassula ovata"),
							),
						),
					),
					kit.Component.Cas("td", goc.Attr{"class": "px-6 py-4 whitespace-nowrap"}, kit.Component.Dcv("inline-flex px-2 py-1 text-xs font-semibold rounded-full bg-green-100 text-green-800", "Healthy")),
					kit.Component.Cav("td", goc.Attr{"class": "px-6 py-4 whitespace-nowrap text-sm text-gray-900"}, "10 days ago"),
					kit.Component.Cav("td", goc.Attr{"class": "px-6 py-4 whitespace-nowrap text-sm text-gray-900"}, "3-5 days"),
					kit.Component.Cav("td", goc.Attr{"class": "px-6 py-4 whitespace-nowrap text-sm text-gray-900"}, "Office"),
					kit.Component.Cas("td", goc.Attr{"class": "px-6 py-4 whitespace-nowrap text-sm text-gray-500"},
						kit.Component.Dcs("flex space-x-2",
							kit.Atoms.Buttons.Button.TertiaryIconLink("pen-to-square", "/plants/3/edit", nil),
							kit.Atoms.Links.Link.H(&links.LinkMod{Label: "View", Url: "/plants/3", HoverCss: "text-blue-600 hover:text-blue-800"}),
						),
					),
				),
			),
		),
	)

	// Pagination
	pagination := kit.Component.Dcs("bg-white px-6 py-5 rounded-lg shadow-sm border border-gray-200",
		kit.Component.Dcs("flex flex-col sm:flex-row justify-between items-center gap-4",
			kit.Component.Dcv("text-sm text-gray-700", "Showing 1 to 12 of 24 results"),
			kit.Component.Dcs("flex space-x-2",
				kit.Atoms.Links.Link.H(&links.LinkMod{Label: "1", Url: "/plants?page=1", HoverCss: "px-4 py-2 text-sm font-medium bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors"}),
				kit.Atoms.Links.Link.H(&links.LinkMod{Label: "2", Url: "/plants?page=2", HoverCss: "px-4 py-2 text-sm font-medium bg-gray-100 text-gray-700 rounded-md hover:bg-gray-200 transition-colors"}),
				kit.Atoms.Links.Link.H(&links.LinkMod{Label: "3", Url: "/plants?page=3", HoverCss: "px-4 py-2 text-sm font-medium bg-gray-100 text-gray-700 rounded-md hover:bg-gray-200 transition-colors"}),
			),
		),
	)

	// Main layout with Alpine.js state
	mainContent := kit.Component.Cas("div",
		goc.Attr{
			"x-data": "{ search: '', filter: 'all' }",
			"class":  "min-h-screen bg-gray-50",
		},
		breadcrumb,
		pageHeader,
		filterSummary,
		kit.Component.Dcs("max-w-7xl mx-auto px-6 py-10",
			plantsTable,
			pagination,
		),
	)

	html := goc.RenderRoot(mainContent)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/list_plants_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/list_plants_1.html")
	sup.AddPage("ListPlants", html)
}
