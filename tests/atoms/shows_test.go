package atoms

import (
	"testing"

	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
)

func TestShows(t *testing.T) {
	kit := kitcla.New()

	// Create a showcase of all show components
	title := kit.Component.Ccv("h1", "text-3xl font-bold text-gray-900 mb-8", "Show Components Overview")

	// Text shows section
	textSection := kit.Component.Dcs("space-y-4",
		kit.Component.Ccv("h2", "text-xl font-semibold text-gray-800 mb-4", "Text Display Components"),
		kit.Component.Dcs("grid grid-cols-1 md:grid-cols-2 gap-6",
			kit.Component.Dcs("p-4 bg-gray-50 rounded-lg border",
				kit.Component.Ccv("h3", "text-sm font-medium text-gray-700 mb-3", "Text Show"),
				kit.Atoms.Shows.TextShow.TextShow("garden_description", "This beautiful garden showcases a variety of seasonal plants, from vibrant spring flowers to robust autumn vegetables."),
			),
			kit.Component.Dcs("p-4 bg-gray-50 rounded-lg border",
				kit.Component.Ccv("h3", "text-sm font-medium text-gray-700 mb-3", "Rich Text Show"),
				kit.Atoms.Shows.RichTextShow.RichTextShow("garden_guide", "<p>Welcome to our <strong>comprehensive garden guide</strong>! Learn about <em>plant care</em>, <a href='/watering-tips' class='text-green-600 hover:underline'>watering techniques</a>, and seasonal maintenance.</p>"),
			),
		),
	)

	// Usage examples section
	usageSection := kit.Component.Dcs("space-y-4",
		kit.Component.Ccv("h2", "text-xl font-semibold text-gray-800 mb-4", "Usage Examples"),
		kit.Component.Dcs("space-y-6",
			kit.Component.Dcs("space-y-2",
				kit.Component.Ccv("h3", "text-lg font-medium text-gray-800", "Content Display"),
				kit.Component.Dcs("bg-white border rounded-lg p-4",
					kit.Component.Ccv("h4", "text-base font-semibold text-gray-900 mb-2", "Article Title"),
					kit.Atoms.Shows.TextShow.TextShow("planting_article", "Spring is the perfect time to start your vegetable garden. Begin by preparing the soil with compost and organic matter. Choose vegetables that thrive in your climate zone for the best results."),
				),
			),
			kit.Component.Dcs("space-y-2",
				kit.Component.Ccv("h3", "text-lg font-medium text-gray-800", "Rich Content Display"),
				kit.Component.Dcs("bg-white border rounded-lg p-4",
					kit.Component.Ccv("h4", "text-base font-semibold text-gray-900 mb-2", "Rich Article"),
					kit.Atoms.Shows.RichTextShow.RichTextShow("garden_tips", "<p>Essential <strong>garden maintenance tips</strong> for a <em>thriving garden</em>: <a href='/soil-care' class='text-green-600 hover:underline'>proper soil preparation</a>, regular watering, and seasonal pruning.</p><ul class='list-disc list-inside mt-2'><li>Test soil pH levels regularly</li><li>Mulch around plants to retain moisture</li><li>Rotate crops annually for soil health</li></ul>"),
				),
			),
		),
	)

	// Main container
	container := kit.Component.Dcs("space-y-8 p-6", title, textSection, usageSection)

	// Render and save
	html := goc.RenderRoot(container)
	sup.AddIndexPage("shows", html)
}
