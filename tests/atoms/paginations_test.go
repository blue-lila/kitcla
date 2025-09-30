package atoms

import (
	"testing"

	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/components/atoms/paginations"
	"github.com/blue-lila/kitcla/dat"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
)

func TestPaginations(t *testing.T) {
	kit := kitcla.New()

	// Create a showcase of pagination components with gardening theme
	title := kit.Component.Ccv("h1", "text-3xl font-bold text-gray-900 mb-8", "Garden Pagination Components")

	// Helper function to create pagination
	createPagination := func(currentPage, totalItems, perPage int, baseUrl string) goc.HTML {
		return kit.Atoms.Paginations.Pagination.H(&paginations.PaginationMod{
			Pagination: &dat.Pagination{
				CurrentPage: currentPage,
				ItemsCount:  totalItems,
				PerPage:     perPage,
			},
			BaseUrl: baseUrl,
		})
	}

	// Basic pagination section
	basicSection := kit.Component.Dcs("space-y-4",
		kit.Component.Ccv("h2", "text-xl font-semibold text-gray-800 mb-4", "Garden Plant List Pagination"),
		kit.Component.Dcs("space-y-6",
			kit.Component.Dcs("flex justify-center p-6 bg-white border rounded-lg",
				createPagination(3, 150, 10, "/plants?page="),
			),
			kit.Component.Ccv("p", "text-sm text-gray-600 text-center", "Showing page 3 of 15 (150 plants total)"),
		),
	)

	// Pagination states section
	statesSection := kit.Component.Dcs("space-y-4",
		kit.Component.Ccv("h2", "text-xl font-semibold text-gray-800 mb-4", "Garden Pagination States"),
		kit.Component.Dcs("space-y-6",
			kit.Component.Dcs("space-y-2",
				kit.Component.Ccv("h3", "text-lg font-medium text-gray-800", "First Page - New Garden"),
				kit.Component.Dcs("flex justify-center p-4 bg-white border rounded-lg",
					createPagination(1, 25, 5, "/garden/plants?page="),
				),
				kit.Component.Ccv("p", "text-xs text-gray-500 text-center", "5 plants per page, 25 total plants"),
			),
			kit.Component.Dcs("space-y-2",
				kit.Component.Ccv("h3", "text-lg font-medium text-gray-800", "Middle Page - Growing Collection"),
				kit.Component.Dcs("flex justify-center p-4 bg-white border rounded-lg",
					createPagination(4, 120, 8, "/vegetable-garden?page="),
				),
				kit.Component.Ccv("p", "text-xs text-gray-500 text-center", "8 vegetables per page, 120 total vegetables"),
			),
			kit.Component.Dcs("space-y-2",
				kit.Component.Ccv("h3", "text-lg font-medium text-gray-800", "Last Page - Complete Garden"),
				kit.Component.Dcs("flex justify-center p-4 bg-white border rounded-lg",
					createPagination(8, 60, 8, "/herb-garden?page="),
				),
				kit.Component.Ccv("p", "text-xs text-gray-500 text-center", "8 herbs per page, 60 total herbs"),
			),
		),
	)

	// Usage examples section
	usageSection := kit.Component.Dcs("space-y-4",
		kit.Component.Ccv("h2", "text-xl font-semibold text-gray-800 mb-4", "Garden Usage Examples"),
		kit.Component.Dcs("bg-green-50 border border-green-200 rounded-lg p-4",
			kit.Component.Ccv("h3", "text-lg font-semibold text-green-800 mb-2", "Garden Application Note"),
			kit.Component.Ccv("p", "text-green-700", "Pagination components are essential for browsing large plant collections, garden inventories, care schedules, and harvest records. They help gardeners navigate through extensive plant databases efficiently."),
		),
	)

	// Main container
	container := kit.Component.Dcs("space-y-8 p-6", title, basicSection, statesSection, usageSection)

	// Render and save
	html := goc.RenderRoot(container)
	sup.AddIndexPage("paginations", html)
}
