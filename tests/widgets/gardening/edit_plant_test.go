package gardening

import (
	"testing"

	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/components/atoms/links"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
)

func TestEditPlant(t *testing.T) {
	kit := kitcla.New()

	// Helper function to create styled form fields
	styledField := func(label string, input goc.HTML) goc.HTML {
		return kit.Component.Dcs("flex flex-col space-y-2",
			kit.Component.Ccv("label", "text-sm font-semibold text-gray-700", label),
			input,
		)
	}

	// Navigation breadcrumb
	breadcrumb := kit.Component.Dcs("bg-white border-b shadow-sm",
		kit.Component.Dcs("max-w-7xl mx-auto px-6 py-4",
			kit.Component.Dcs("flex items-center space-x-2 text-sm",
				kit.Atoms.Links.Link.H(&links.LinkMod{Label: "Garden", Url: "/garden", HoverCss: "text-blue-600 hover:text-blue-800"}),
				kit.Component.Dcv("text-gray-400", "›"),
				kit.Atoms.Links.Link.H(&links.LinkMod{Label: "Plants", Url: "/plants", HoverCss: "text-blue-600 hover:text-blue-800"}),
				kit.Component.Dcv("text-gray-400", "›"),
				kit.Atoms.Links.Link.H(&links.LinkMod{Label: "Monstera Deliciosa", Url: "/plants/1", HoverCss: "text-blue-600 hover:text-blue-800"}),
				kit.Component.Dcv("text-gray-400", "›"),
				kit.Component.Dcv("text-gray-700 font-medium", "Edit"),
			),
		),
	)

	// Page header
	pageHeader := kit.Component.Dcs("bg-white border-b",
		kit.Component.Dcs("max-w-7xl mx-auto px-6 py-8",
			kit.Component.Ccv("h1", "text-3xl font-bold text-gray-900 mb-2", "✏️ Edit Plant"),
			kit.Component.Dcv("text-gray-600 text-lg", "Update your plant information and care settings"),
		),
	)

	// Basic Information Form Section
	basicInfoForm := kit.Component.Dcs("mb-8",
		kit.Organisms.Cards.Card.SimpleCard(
			"🌿 Basic Information",
			kit.Component.Ds(
				kit.Component.Dcs("space-y-6",
					kit.Component.Dcs("grid grid-cols-1 lg:grid-cols-2 gap-6",
						styledField(
							"Common Name",
							kit.Component.Cav("input", goc.Attr{
								"type":        "text",
								"name":        "common_name",
								"value":       "Monstera Deliciosa",
								"class":       "block w-full rounded-lg border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 text-gray-900 placeholder-gray-400 px-4 py-3 text-base transition-all duration-200",
								"placeholder": "Enter plant name",
							}),
						),
						styledField(
							"Scientific Name",
							kit.Component.Cav("input", goc.Attr{
								"type":        "text",
								"name":        "scientific_name",
								"value":       "Monstera deliciosa",
								"class":       "block w-full rounded-lg border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 text-gray-900 placeholder-gray-400 px-4 py-3 text-base transition-all duration-200",
								"placeholder": "Enter scientific name",
							}),
						),
					),
					kit.Component.Dcs("grid grid-cols-1 lg:grid-cols-2 gap-6",
						styledField(
							"Plant Type",
							kit.Component.Cas("select", goc.Attr{
								"name":  "plant_type",
								"class": "block w-full rounded-lg border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 text-gray-900 px-4 py-3 text-base transition-all duration-200 bg-white",
							},
								kit.Component.Cav("option", goc.Attr{"value": "houseplant", "selected": ""}, "🪴 Houseplant"),
								kit.Component.Cav("option", goc.Attr{"value": "herb"}, "🌿 Herb"),
								kit.Component.Cav("option", goc.Attr{"value": "flowering"}, "🌸 Flowering Plant"),
								kit.Component.Cav("option", goc.Attr{"value": "succulent"}, "🌵 Succulent"),
								kit.Component.Cav("option", goc.Attr{"value": "tree"}, "🌳 Tree"),
							),
						),
						styledField(
							"Location",
							kit.Component.Cav("input", goc.Attr{
								"type":        "text",
								"name":        "location",
								"value":       "Living Room - East Window",
								"class":       "block w-full rounded-lg border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 text-gray-900 placeholder-gray-400 px-4 py-3 text-base transition-all duration-200",
								"placeholder": "e.g., Living Room - East Window",
							}),
						),
					),
					styledField(
						"Notes",
						kit.Component.Cav("textarea", goc.Attr{
							"name":        "notes",
							"rows":        "4",
							"class":       "block w-full rounded-lg border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 text-gray-900 placeholder-gray-400 px-4 py-3 text-base transition-all duration-200 resize-y",
							"placeholder": "Add any notes about this plant...",
						}, "This Monstera has been thriving since I brought it home. The large fenestrated leaves are beautiful and it's growing quite quickly."),
					),
				),
			),
		),
	)

	// Care Settings Form Section
	careSettingsForm := kit.Component.Dcs("mb-8",
		kit.Organisms.Cards.Card.SimpleCard(
			"💧 Care Settings",
			kit.Component.Ds(
				kit.Component.Dcs("space-y-6",
					kit.Component.Dcs("grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6",
						styledField(
							"Watering Frequency",
							kit.Component.Dcs("relative",
								kit.Component.Cav("input", goc.Attr{
									"type":  "number",
									"name":  "watering_days",
									"value": "7",
									"min":   "1",
									"class": "block w-full rounded-lg border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 text-gray-900 px-4 py-3 pr-16 text-base transition-all duration-200",
								}),
								kit.Component.Dcs("absolute right-4 top-1/2 -translate-y-1/2 text-sm font-medium text-gray-500 pointer-events-none",
									kit.Component.Dcv("", "days"),
								),
							),
						),
						styledField(
							"Pot Size",
							kit.Component.Dcs("relative",
								kit.Component.Cav("input", goc.Attr{
									"type":  "number",
									"name":  "pot_size",
									"value": "12",
									"min":   "1",
									"class": "block w-full rounded-lg border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 text-gray-900 px-4 py-3 pr-20 text-base transition-all duration-200",
								}),
								kit.Component.Dcs("absolute right-4 top-1/2 -translate-y-1/2 text-sm font-medium text-gray-500 pointer-events-none",
									kit.Component.Dcv("", "inches"),
								),
							),
						),
						styledField(
							"Fertilizer Schedule",
							kit.Component.Cas("select", goc.Attr{
								"name":  "fertilizer_frequency",
								"class": "block w-full rounded-lg border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 text-gray-900 px-4 py-3 text-base transition-all duration-200 bg-white",
							},
								kit.Component.Cav("option", goc.Attr{"value": "weekly"}, "Weekly"),
								kit.Component.Cav("option", goc.Attr{"value": "biweekly"}, "Bi-weekly"),
								kit.Component.Cav("option", goc.Attr{"value": "monthly", "selected": ""}, "Monthly"),
								kit.Component.Cav("option", goc.Attr{"value": "seasonal"}, "Seasonal"),
								kit.Component.Cav("option", goc.Attr{"value": "never"}, "Never"),
							),
						),
					),
					kit.Component.Dcs("grid grid-cols-1 md:grid-cols-2 gap-6",
						styledField(
							"Light Requirements",
							kit.Component.Cas("select", goc.Attr{
								"name":  "light_level",
								"class": "block w-full rounded-lg border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 text-gray-900 px-4 py-3 text-base transition-all duration-200 bg-white",
							},
								kit.Component.Cav("option", goc.Attr{"value": "low"}, "🌑 Low Light"),
								kit.Component.Cav("option", goc.Attr{"value": "medium"}, "☁️ Medium Light"),
								kit.Component.Cav("option", goc.Attr{"value": "bright-indirect", "selected": ""}, "🌤️ Bright Indirect"),
								kit.Component.Cav("option", goc.Attr{"value": "bright-direct"}, "☀️ Bright Direct"),
								kit.Component.Cav("option", goc.Attr{"value": "full-sun"}, "🔆 Full Sun"),
							),
						),
						styledField(
							"Soil Type",
							kit.Component.Cav("input", goc.Attr{
								"type":        "text",
								"name":        "soil_type",
								"value":       "Well-draining potting mix",
								"class":       "block w-full rounded-lg border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 text-gray-900 placeholder-gray-400 px-4 py-3 text-base transition-all duration-200",
								"placeholder": "e.g., Well-draining potting mix",
							}),
						),
					),
				),
			),
		),
	)

	// Plant Status Section
	statusForm := kit.Component.Dcs("mb-8",
		kit.Organisms.Cards.Card.SimpleCard(
			"📊 Current Status",
			kit.Component.Ds(
				kit.Component.Dcs("grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6",
					styledField(
						"Health Status",
						kit.Component.Cas("select", goc.Attr{
							"name":  "health_status",
							"class": "block w-full rounded-lg border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 text-gray-900 px-4 py-3 text-base transition-all duration-200 bg-white",
						},
							kit.Component.Cav("option", goc.Attr{"value": "excellent"}, "✨ Excellent"),
							kit.Component.Cav("option", goc.Attr{"value": "good", "selected": ""}, "✅ Good"),
							kit.Component.Cav("option", goc.Attr{"value": "fair"}, "⚠️ Fair"),
							kit.Component.Cav("option", goc.Attr{"value": "poor"}, "🔴 Poor"),
							kit.Component.Cav("option", goc.Attr{"value": "critical"}, "💀 Critical"),
						),
					),
					styledField(
						"Last Watered",
						kit.Component.Cav("input", goc.Attr{
							"type":  "date",
							"name":  "last_watered",
							"value": "2024-03-20",
							"class": "block w-full rounded-lg border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 text-gray-900 px-4 py-3 text-base transition-all duration-200",
						}),
					),
					styledField(
						"Humidity Level",
						kit.Component.Dcs("relative",
							kit.Component.Cav("input", goc.Attr{
								"type":        "number",
								"name":        "humidity",
								"value":       "60",
								"min":         "0",
								"max":         "100",
								"class":       "block w-full rounded-lg border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 text-gray-900 px-4 py-3 pr-12 text-base transition-all duration-200",
								"placeholder": "0-100",
							}),
							kit.Component.Dcs("absolute right-4 top-1/2 -translate-y-1/2 text-sm font-medium text-gray-500 pointer-events-none",
								kit.Component.Dcv("", "%"),
							),
						),
					),
				),
			),
		),
	)

	// Action buttons with sticky footer
	actionButtons := kit.Component.Dcs("sticky bottom-0 bg-white border-t border-gray-200 shadow-lg",
		kit.Component.Dcs("max-w-7xl mx-auto px-6 py-4 flex flex-col sm:flex-row justify-between items-stretch sm:items-center gap-4",
			kit.Component.Dcs("flex flex-col sm:flex-row gap-3 order-2 sm:order-1",
				kit.Atoms.Links.Link.H(&links.LinkMod{Label: "Cancel", Url: "/plants/1", HoverCss: "inline-flex items-center justify-center px-4 py-2 text-gray-700 hover:text-gray-900 font-medium transition-colors"}),
				kit.Atoms.Buttons.Button.SecondaryLink("🗑️ Delete Plant", "/plants/1/delete", nil),
			),
			kit.Component.Dcs("flex flex-col sm:flex-row gap-3 order-1 sm:order-2",
				kit.Atoms.Buttons.Button.FormSubmit("💾 Save Draft", nil),
				kit.Atoms.Buttons.Button.PrimarySubmit("✅ Save Changes", nil),
			),
		),
	)

	// Main form layout
	mainContent := kit.Component.Cas("form",
		goc.Attr{
			"method": "POST",
			"action": "/plants/1",
			"class":  "min-h-screen bg-gradient-to-b from-gray-50 to-gray-100",
		},
		breadcrumb,
		pageHeader,
		kit.Component.Dcs("max-w-5xl mx-auto px-4 sm:px-6 lg:px-8 py-8 pb-32",
			kit.Component.Dcs("mb-8",
				kit.Molecules.Alerts.Alert.InfoAlert(
					"💡 Pro Tip",
					"Regular updates help track your plant's progress and identify patterns in their care needs.",
					nil,
				),
			),
			basicInfoForm,
			careSettingsForm,
			statusForm,
		),
		actionButtons,
	)

	html := goc.RenderRoot(mainContent)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/edit_plant_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/edit_plant_1.html")
	sup.AddPage("EditPlant", html)
}
