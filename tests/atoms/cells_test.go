package atoms

import (
	"testing"

	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
)

func TestCells(t *testing.T) {
	kit := kitcla.New()

	// Create a showcase of all cell types
	title := kit.Component.Ccv("h1", "text-3xl font-bold text-gray-900 mb-8", "Cell Components Overview")

	// Text-based cells section
	textSection := kit.Component.Dcs("space-y-4",
		kit.Component.Ccv("h2", "text-xl font-semibold text-gray-800 mb-4", "Text Cells"),
		kit.Component.Dcs("grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4",
			kit.Component.Dcs("p-3 border rounded-lg bg-gray-50",
				kit.Component.Ccv("h3", "text-sm font-medium text-gray-600 mb-2", "Text Cell"),
				kit.Atoms.Cells.TextCell.TextCell("Sample Text"),
			),
			kit.Component.Dcs("p-3 border rounded-lg bg-gray-50",
				kit.Component.Ccv("h3", "text-sm font-medium text-gray-600 mb-2", "Long Text Cell"),
				kit.Atoms.Cells.LongTextCell.LongTextCell("This is a longer text sample"),
			),
			kit.Component.Dcs("p-3 border rounded-lg bg-gray-50",
				kit.Component.Ccv("h3", "text-sm font-medium text-gray-600 mb-2", "Pill Cell"),
				kit.Atoms.Cells.PillCell.PillCell("Active"),
			),
			kit.Component.Dcs("p-3 border rounded-lg bg-gray-50",
				kit.Component.Ccv("h3", "text-sm font-medium text-gray-600 mb-2", "Boolean Cell"),
				kit.Atoms.Cells.BooleanCell.BooleanCell(true),
			),
		),
	)

	// Main container
	container := kit.Component.Dcs("space-y-8 p-6", title, textSection)

	// Render and save
	html := goc.RenderRoot(container)
	sup.AddIndexPage("cells", html)
}
