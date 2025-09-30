package atoms

import (
	"testing"

	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
)

func TestButtons(t *testing.T) {
	kit := kitcla.New()

	// Create a showcase of all button types
	title := kit.Component.Ccv("h1", "text-3xl font-bold text-gray-900 mb-8", "Button Components Overview")

	// Helper function to create table row
	tableRow := func(method, kind, htmlType, description string, rendering goc.HTML) goc.HTML {
		return kit.Component.Cs("tr",
			kit.Component.Ccv("td", "border px-4 py-2 font-mono text-sm", method),
			kit.Component.Ccv("td", "border px-4 py-2", kind),
			kit.Component.Ccv("td", "border px-4 py-2", htmlType),
			kit.Component.Ccv("td", "border px-4 py-2", description),
			kit.Component.Cas("td", goc.Attr{"class": "border px-4 py-2"}, rendering),
		)
	}

	// Table header
	tableHeader := kit.Component.Cs("thead",
		kit.Component.Cs("tr",
			kit.Component.Ccv("th", "border px-4 py-2 bg-gray-100 text-left", "Method"),
			kit.Component.Ccv("th", "border px-4 py-2 bg-gray-100 text-left", "Kind"),
			kit.Component.Ccv("th", "border px-4 py-2 bg-gray-100 text-left", "HTML Type"),
			kit.Component.Ccv("th", "border px-4 py-2 bg-gray-100 text-left", "Description"),
			kit.Component.Ccv("th", "border px-4 py-2 bg-gray-100 text-left", "Rendering"),
		),
	)

	// Primary buttons section
	primaryTitle := kit.Component.Ccv("h2", "text-2xl font-semibold text-gray-800 mb-4 mt-8", "Primary Buttons")
	primaryRows := kit.Component.Cs("tbody",
		tableRow("PrimaryLink", "Primary", "Link (a)", "Standard link button with primary styling", kit.Atoms.Buttons.Button.PrimaryLink("Primary Link", "#", nil)),
		tableRow("PrimarySubmit", "Primary", "Submit", "Submit button for forms with primary styling", kit.Atoms.Buttons.Button.PrimarySubmit("Primary Submit", nil)),
		tableRow("PrimaryPost", "Primary", "POST Form", "POST request button with primary styling", kit.Atoms.Buttons.Button.PrimaryPost("Primary Post", "#", nil)),
		tableRow("PrimaryIconLink", "Primary", "Link (a)", "Icon-only link button with primary styling", kit.Atoms.Buttons.Button.PrimaryIconLink("tree", "#", nil)),
		tableRow("PrimaryIconSubmit", "Primary", "Submit", "Icon-only submit button with primary styling", kit.Atoms.Buttons.Button.PrimaryIconSubmit("tree", nil)),
		tableRow("PrimaryIconPost", "Primary", "POST Form", "Icon-only POST request button with primary styling", kit.Atoms.Buttons.Button.PrimaryIconPost("tree", "#", nil)),
	)
	primaryTable := kit.Component.Cas("table", goc.Attr{"class": "w-full border-collapse"}, tableHeader, primaryRows)
	primarySection := kit.Component.Dcs("space-y-4", primaryTitle, primaryTable)

	// Secondary buttons section
	secondaryTitle := kit.Component.Ccv("h2", "text-2xl font-semibold text-gray-800 mb-4 mt-8", "Secondary Buttons")
	secondaryRows := kit.Component.Cs("tbody",
		tableRow("SecondaryLink", "Secondary", "Link (a)", "Standard link button with secondary styling", kit.Atoms.Buttons.Button.SecondaryLink("Secondary Link", "#", nil)),
		tableRow("SecondarySubmit", "Secondary", "Submit", "Submit button for forms with secondary styling", kit.Atoms.Buttons.Button.SecondarySubmit("Secondary Submit", nil)),
		tableRow("SecondaryPost", "Secondary", "POST Form", "POST request button with secondary styling", kit.Atoms.Buttons.Button.SecondaryPost("Secondary Post", "#", nil)),
		tableRow("SecondaryIconLink", "Secondary", "Link (a)", "Icon-only link button with secondary styling", kit.Atoms.Buttons.Button.SecondaryIconLink("tree", "#", nil)),
		tableRow("SecondaryIconSubmit", "Secondary", "Submit", "Icon-only submit button with secondary styling", kit.Atoms.Buttons.Button.SecondaryIconSubmit("tree", nil)),
		tableRow("SecondaryIconPost", "Secondary", "POST Form", "Icon-only POST request button with secondary styling", kit.Atoms.Buttons.Button.SecondaryIconPost("tree", "#", nil)),
	)
	secondaryTable := kit.Component.Cas("table", goc.Attr{"class": "w-full border-collapse"}, tableHeader, secondaryRows)
	secondarySection := kit.Component.Dcs("space-y-4", secondaryTitle, secondaryTable)

	// Tertiary buttons section
	tertiaryTitle := kit.Component.Ccv("h2", "text-2xl font-semibold text-gray-800 mb-4 mt-8", "Tertiary Buttons")
	tertiaryRows := kit.Component.Cs("tbody",
		tableRow("TertiaryIconLink", "Tertiary (Ghost)", "Link (a)", "Icon-only link button with ghost/tertiary styling", kit.Atoms.Buttons.Button.TertiaryIconLink("tree", "#", nil)),
	)
	tertiaryTable := kit.Component.Cas("table", goc.Attr{"class": "w-full border-collapse"}, tableHeader, tertiaryRows)
	tertiarySection := kit.Component.Dcs("space-y-4", tertiaryTitle, tertiaryTable)

	// Table buttons section
	tableButtonsTitle := kit.Component.Ccv("h2", "text-2xl font-semibold text-gray-800 mb-4 mt-8", "Table Buttons")
	tableButtonsRows := kit.Component.Cs("tbody",
		tableRow("TableLink", "Secondary", "Link (a)", "Compact link button for use in tables (base size)", kit.Atoms.Buttons.Button.TableLink("Table Link", "#", nil)),
		tableRow("TableIconLink", "Secondary", "Link (a)", "Compact icon-only button for use in tables (base size)", kit.Atoms.Buttons.Button.TableIconLink("tree", "#", nil)),
	)
	tableButtonsTable := kit.Component.Cas("table", goc.Attr{"class": "w-full border-collapse"}, tableHeader, tableButtonsRows)
	tableButtonsSection := kit.Component.Dcs("space-y-4", tableButtonsTitle, tableButtonsTable)

	// Special purpose buttons section
	specialTitle := kit.Component.Ccv("h2", "text-2xl font-semibold text-gray-800 mb-4 mt-8", "Special Purpose Buttons")
	specialRows := kit.Component.Cs("tbody",
		tableRow("FormSubmit", "Primary", "Submit", "Form submit button with XL size for main forms", kit.Atoms.Buttons.Button.FormSubmit("Form Submit", nil)),
		tableRow("RemoteForm", "Secondary", "Link (a)", "Button that triggers remote form loading via JavaScript", kit.Atoms.Buttons.Button.RemoteForm("Remote Form", "/remote", nil)),
	)
	specialTable := kit.Component.Cas("table", goc.Attr{"class": "w-full border-collapse"}, tableHeader, specialRows)
	specialSection := kit.Component.Dcs("space-y-4", specialTitle, specialTable)

	// Alpine.js buttons section
	buttonAlpTitle := kit.Component.Ccv("h2", "text-2xl font-semibold text-gray-800 mb-4 mt-8", "Alpine.js Buttons")
	buttonAlpRows := kit.Component.Cs("tbody",
		tableRow("ButtonAlp.PrimaryLink", "Primary", "Link (a)", "Alpine.js reactive link button with primary styling", kit.Atoms.Buttons.ButtonAlp.PrimaryLink("Primary Alpine", "", nil)),
		tableRow("ButtonAlp.SecondaryLink", "Secondary", "Link (a)", "Alpine.js reactive link button with secondary styling", kit.Atoms.Buttons.ButtonAlp.SecondaryLink("Secondary Alpine", "", nil)),
		tableRow("ButtonAlp.SecondaryIconLink", "Secondary", "Link (a)", "Alpine.js reactive icon-only button with secondary styling", kit.Atoms.Buttons.ButtonAlp.SecondaryIconLink("tree", "", nil)),
		tableRow("ButtonAlp.TertiaryIconLink", "Tertiary", "Link (a)", "Alpine.js reactive icon-only button with tertiary styling", kit.Atoms.Buttons.ButtonAlp.TertiaryIconLink("tree", "", nil)),
		tableRow("ButtonAlp.QuaternaryIconLink", "Quaternary", "Link (a)", "Alpine.js reactive icon-only button with quaternary styling", kit.Atoms.Buttons.ButtonAlp.QuaternaryIconLink("tree", "", nil)),
	)
	buttonAlpTable := kit.Component.Cas("table", goc.Attr{"class": "w-full border-collapse"}, tableHeader, buttonAlpRows)
	buttonAlpSection := kit.Component.Dcs("space-y-4", buttonAlpTitle, buttonAlpTable)

	// Main container
	container := kit.Component.Dcs("space-y-8 p-6", title, primarySection, secondarySection, tertiarySection, tableButtonsSection, specialSection, buttonAlpSection)

	// Render and save
	html := goc.RenderRoot(container)
	sup.AddIndexPage("buttons", html)
}
