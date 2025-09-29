package paginations

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/dat"
	"github.com/blue-lila/kitcla/goc"
	"net/url"
	"sort"
	"strconv"
)

type Pagination struct {
	Component *components.Component
}

type PaginationMod struct {
	Pagination *dat.Pagination
	BaseUrl    string
}

func (this *Pagination) H(mod *PaginationMod) goc.HTML {
	var pages []int

	pagination := mod.Pagination

	if pagination.PerPage < 1 {
		panic("Invalid Mod: pagination per page cannot be lower than 1")
	}
	if pagination.ItemsCount <= pagination.PerPage {
		return goc.HTML{}
	}

	lastPage := ((pagination.ItemsCount - 1) / (pagination.PerPage)) + 1

	pages = append(pages, pagination.CurrentPage)
	pages = this.addSurroundingPages(pages, pagination.CurrentPage, lastPage)
	pages = this.addFirstPages(lastPage, pages, pagination.CurrentPage)
	pages = this.addLastPages(lastPage, pages, pagination.CurrentPage)

	pages = this.unique(pages)
	sort.Ints(pages)

	var inter []goc.HTML
	inter = append(inter, this.paginationPrevArrowButton(mod, lastPage))

	inter = append(inter, this.buildPaginationItems(mod, pages, lastPage)...)

	inter = append(inter, this.paginationNextArrowButton(mod, lastPage))

	return this.Component.Cas("nav",
		goc.Attr{
			"aria-label": "Pagination navigation",
			"role":       "navigation",
		},
		this.Component.Ccs("div", "flex -space-x-px text-scale-7", inter...),
	)
}

func (this *Pagination) addLastPages(lastPage int, pages []int, currentPage int) []int {
	var pages2 []int
	j := lastPage
	for j > 0 {
		if j <= lastPage-3 {
			break
		}
		pages2 = append(pages2, j)
		j--
	}
	pages2 = this.reverse(pages2)
	for _, p := range pages2 {
		pages = append(pages, p)
	}
	return pages
}

func (this *Pagination) addFirstPages(lastPage int, pages []int, currentPage int) []int {
	i := 1
	for i <= lastPage {
		if i > 3 {
			break
		}
		pages = append(pages, i)
		i++
	}
	return pages
}

func (this *Pagination) Mod() *PaginationMod {
	return &PaginationMod{}
}

func (this *Pagination) paginationPageButton(mod *PaginationMod, page int) goc.HTML {
	urlChanged := this.changePageValue(mod, "set", page)
	css := "bg-scale-0 hover:bg-scale-1"
	ariaLabel := "Go to page " + strconv.Itoa(page)
	ariaCurrent := ""

	if page == mod.Pagination.CurrentPage {
		css = "bg-blue-50 text-blue-600"
		ariaLabel = "Current page " + strconv.Itoa(page)
		ariaCurrent = "page"
	}

	attrs := goc.Attr{
		"href":       urlChanged,
		"class":      "border px-4 py-2 cursor-pointer inline-flex items-center text-sm font-medium transition-colors " + css,
		"aria-label": ariaLabel,
	}

	if ariaCurrent != "" {
		attrs["aria-current"] = ariaCurrent
	}

	return this.Component.Cav("a", attrs, strconv.Itoa(page))
}

func (this *Pagination) paginationPrevArrowButton(mod *PaginationMod, lastPage int) goc.HTML {
	isDisabled := mod.Pagination.CurrentPage <= 1
	urlChanged := this.changePageValue(mod, "decrement", 0)

	css := "border px-4 py-2 rounded-l-md bg-scale-0 inline-flex items-center text-sm font-medium transition-colors"

	if isDisabled {
		css += " opacity-50 cursor-not-allowed text-gray-400"
		return this.Component.Cav("span",
			goc.Attr{
				"class":         css,
				"aria-label":    "Previous page (disabled)",
				"aria-disabled": "true",
			},
			"‹")
	}

	css += " cursor-pointer hover:bg-scale-1"
	return this.Component.Cav("a",
		goc.Attr{
			"href":       urlChanged,
			"class":      css,
			"aria-label": "Go to previous page",
		},
		"‹")
}

func (this *Pagination) paginationNextArrowButton(mod *PaginationMod, pageMax int) goc.HTML {
	isDisabled := mod.Pagination.CurrentPage >= pageMax
	urlChanged := this.changePageValue(mod, "increment", pageMax)

	css := "border rounded-r-md px-4 py-2 bg-scale-0 inline-flex items-center text-sm font-medium transition-colors"

	if isDisabled {
		css += " opacity-50 cursor-not-allowed text-gray-400"
		return this.Component.Cav("span",
			goc.Attr{
				"class":         css,
				"aria-label":    "Next page (disabled)",
				"aria-disabled": "true",
			},
			"›")
	}

	css += " cursor-pointer hover:bg-scale-1"
	return this.Component.Cav("a",
		goc.Attr{
			"href":       urlChanged,
			"class":      css,
			"aria-label": "Go to next page",
		},
		"›")
}

func (this *Pagination) changePageValue(mod *PaginationMod, operation string, value int) string {
	urlParsed, err := url.Parse(mod.BaseUrl)
	if err != nil {
		panic("Cannot parse base url")
	}
	values := urlParsed.Query()
	current := 1
	page := values.Get("page")
	if page != "" {
		current, _ = strconv.Atoi(page)
	}
	if operation == "set" {
		current = value
	}
	if operation == "increment" {
		current++
	}
	if operation == "decrement" {
		current--
	}
	if current < 1 {
		current = 1
	}
	if operation == "increment" && current > value {
		current = value
	}
	values.Set("page", strconv.Itoa(current))
	urlParsed.RawQuery = values.Encode()
	return urlParsed.String()
}

func (this *Pagination) unique(intSlice []int) []int {
	keys := make(map[int]bool)
	var list []int
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func (this *Pagination) reverse(intSlice []int) []int {
	for i, j := 0, len(intSlice)-1; i < j; i, j = i+1, j-1 {
		intSlice[i], intSlice[j] = intSlice[j], intSlice[i]
	}
	return intSlice
}

func (this *Pagination) addSurroundingPages(pages []int, page int, lastPage int) []int {
	if page-1 > 0 {
		pages = append(pages, page-1)
	}
	if page+1 <= lastPage {
		pages = append(pages, page+1)
	}
	return pages
}

func (this *Pagination) buildPaginationItems(mod *PaginationMod, pages []int, lastPage int) []goc.HTML {
	var items []goc.HTML

	if len(pages) == 0 {
		return items
	}

	prevPage := 0
	for i, page := range pages {
		// Add ellipsis if there's a gap greater than 1
		if i > 0 && page > prevPage+1 {
			items = append(items, this.paginationEllipsis())
		}

		items = append(items, this.paginationPageButton(mod, page))
		prevPage = page
	}

	return items
}

func (this *Pagination) paginationEllipsis() goc.HTML {
	return this.Component.Cav("span",
		goc.Attr{
			"class":       "border px-4 py-2 bg-scale-0 inline-flex items-center text-sm font-medium text-gray-500",
			"aria-hidden": "true",
		},
		"...")
}
