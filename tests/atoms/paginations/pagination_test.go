package paginations

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/dat"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"testing"
)

func TestButtonH(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Paginations.Pagination
	mod := component.Mod()

	// Set up test data
	mod.Pagination = &dat.Pagination{
		ItemsCount:  250,
		PerPage:     10,
		CurrentPage: 10,
	}
	mod.BaseUrl = "/test"

	h := component.H(mod)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/pagination_h.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/pagination_h.html")
	sup.AddPage("H", html)
}

func TestPaginationEllipsisGaps(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Paginations.Pagination
	mod := component.Mod()

	mod.Pagination = &dat.Pagination{
		ItemsCount:  200,
		PerPage:     10,
		CurrentPage: 10,
	}
	mod.BaseUrl = "/test"

	h := component.H(mod)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/pagination_ellipsis_gaps.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/pagination_ellipsis_gaps.html")
	sup.AddPage("EllipsisGaps", html)
}

func TestPaginationFirstPage(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Paginations.Pagination
	mod := component.Mod()

	mod.Pagination = &dat.Pagination{
		ItemsCount:  100,
		PerPage:     10,
		CurrentPage: 1,
	}
	mod.BaseUrl = "/test"

	h := component.H(mod)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/pagination_first_page.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/pagination_first_page.html")
	sup.AddPage("FirstPage", html)
}

func TestPaginationLastPage(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Paginations.Pagination
	mod := component.Mod()

	mod.Pagination = &dat.Pagination{
		ItemsCount:  100,
		PerPage:     10,
		CurrentPage: 10, // Last page
	}
	mod.BaseUrl = "/test"

	h := component.H(mod)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/pagination_last_page.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/pagination_last_page.html")
	sup.AddPage("LastPage", html)
}
