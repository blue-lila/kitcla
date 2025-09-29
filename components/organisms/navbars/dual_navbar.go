package navbars

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/components/atoms/icons"
	"github.com/blue-lila/kitcla/components/organisms/menus"
	"github.com/blue-lila/kitcla/goc"
)

type DualNavbar struct {
	Component *components.Component
	Icon      *icons.Icon
}

type BreadcrumbItem struct {
	Label string
	Link  string
}

type DualNavbarMod struct {
	MenuItems       []*menus.MenuItem
	BreadcrumbItems []*BreadcrumbItem
	OpenWhenCreated bool
}

func (this *DualNavbar) Mod() *DualNavbarMod {
	return &DualNavbarMod{}
}

func (this *DualNavbar) DualNavbar(mod *DualNavbarMod) goc.HTML {
	return this.H(mod)
}

func (this *DualNavbar) H(mod *DualNavbarMod) goc.HTML {
	initialState := "false"
	if mod.OpenWhenCreated {
		initialState = "true"
	}

	return this.Component.Das(
		goc.Attr{
			"class":       "relative",
			"x-data":      "{ open: " + initialState + " }",
			"@mouseenter": "open = true",
			"@mouseleave": "open = false",
		},
		this.menuButton(mod),
		this.dropdownContent(mod),
	)
}

func (this *DualNavbar) menuButton(mod *DualNavbarMod) goc.HTML {
	return this.Component.Cas("button",
		goc.Attr{"class": "px-4 py-2 bg-white text-gray-700 rounded border hover:bg-gray-50 transition-colors flex items-center justify-center"},
		this.Icon.Icon(icons.IconFasBars),
	)
}

func (this *DualNavbar) dropdownContent(mod *DualNavbarMod) goc.HTML {
	return this.Component.Das(
		goc.Attr{
			"class":        "absolute left-0 min-w-max bg-white rounded-lg shadow-lg border z-50",
			"x-show":       "open",
			"x-transition": "",
		},
		this.dropdownGrid(mod),
	)
}

func (this *DualNavbar) dropdownGrid(mod *DualNavbarMod) goc.HTML {
	return this.Component.Dcs("flex flex-row p-6 divide-x divide-gray-200",
		this.menuColumn(mod),
		this.breadcrumbColumn(mod),
	)
}

func (this *DualNavbar) menuColumn(mod *DualNavbarMod) goc.HTML {
	menuLinks := this.buildMenuLinks(mod, mod.MenuItems)
	return this.Component.Dcs("flex flex-col space-y-2 pr-6 w-1/2",
		append([]goc.HTML{this.menuColumnTitle(mod)}, menuLinks...)...)
}

func (this *DualNavbar) breadcrumbColumn(mod *DualNavbarMod) goc.HTML {
	breadcrumbLinks := this.buildBreadcrumbLinks(mod, mod.BreadcrumbItems)
	return this.Component.Dcs("flex flex-col space-y-2 pl-6 w-1/2",
		append([]goc.HTML{this.breadcrumbColumnTitle(mod)}, breadcrumbLinks...)...)
}

func (this *DualNavbar) menuColumnTitle(mod *DualNavbarMod) goc.HTML {
	return this.Component.Dcv("font-semibold text-gray-900 w-64 px-3 py-1", "Menu")
}

func (this *DualNavbar) breadcrumbColumnTitle(mod *DualNavbarMod) goc.HTML {
	return this.Component.Dcv("font-semibold text-gray-900 w-64 px-3 py-1", "Breadcrumb")
}

func (this *DualNavbar) buildMenuLinks(mod *DualNavbarMod, menuItems []*menus.MenuItem) []goc.HTML {
	var menuLinks []goc.HTML
	for _, link := range menuItems {
		menuLinks = append(menuLinks, this.menuItem(mod, link))
	}
	return menuLinks
}

func (this *DualNavbar) menuItem(mod *DualNavbarMod, link *menus.MenuItem) goc.HTML {
	return this.Component.Dcs("flex pb-1 -ml-3",
		this.menuItemContent(mod, link),
	)
}

func (this *DualNavbar) menuItemContent(mod *DualNavbarMod, link *menus.MenuItem) goc.HTML {
	return this.Component.Dcs("grow pt-0.5 ml-3",
		this.menuItemLink(mod, link),
	)
}

func (this *DualNavbar) menuItemLink(mod *DualNavbarMod, link *menus.MenuItem) goc.HTML {
	return this.Component.Cav("a",
		goc.Attr{"href": link.Url, "class": "text-sm font-medium text-gray-800 hover:text-blue-600 hover:bg-gray-100 px-3 py-2 rounded transition-colors block"},
		link.Name)
}

func (this *DualNavbar) buildBreadcrumbLinks(mod *DualNavbarMod, breadcrumbItems []*BreadcrumbItem) []goc.HTML {
	var breadcrumbLinks []goc.HTML
	for i, link := range breadcrumbItems {
		isLast := i == len(breadcrumbItems)-1
		breadcrumbLinks = append(breadcrumbLinks, this.breadcrumbItem(mod, link, isLast))
	}
	return breadcrumbLinks
}

func (this *DualNavbar) breadcrumbItem(mod *DualNavbarMod, link *BreadcrumbItem, isLast bool) goc.HTML {
	return this.Component.Dcs("flex",
		this.breadcrumbDotContainer(mod, isLast),
		this.breadcrumbItemContent(mod, link),
	)
}

func (this *DualNavbar) breadcrumbDotContainer(mod *DualNavbarMod, isLast bool) goc.HTML {
	afterClass := "relative"
	if !isLast {
		afterClass += " after:absolute after:top-7 after:bottom-0 after:start-3.5 after:w-px after:-translate-x-[0.5px] after:bg-gray-200"
	}
	return this.Component.Dcs(afterClass,
		this.breadcrumbDotWrapper(mod),
	)
}

func (this *DualNavbar) breadcrumbDotWrapper(mod *DualNavbarMod) goc.HTML {
	return this.Component.Dcs("relative z-10 flex h-7 w-7 mt-1.5 items-center justify-center",
		this.breadcrumbDot(mod),
	)
}

func (this *DualNavbar) breadcrumbDot(mod *DualNavbarMod) goc.HTML {
	return this.Component.Dcv("h-2 w-2 rounded-full bg-gray-400", "")
}

func (this *DualNavbar) breadcrumbItemContent(mod *DualNavbarMod, link *BreadcrumbItem) goc.HTML {
	return this.Component.Dcs("grow pt-0.5 pb-1 ml-4",
		this.breadcrumbItemLink(mod, link),
	)
}

func (this *DualNavbar) breadcrumbItemLink(mod *DualNavbarMod, link *BreadcrumbItem) goc.HTML {
	return this.Component.Cav("a",
		goc.Attr{"href": link.Link, "class": "text-sm font-medium text-gray-800 hover:text-blue-600 hover:bg-gray-100 px-3 py-2 rounded transition-colors block"},
		link.Label)
}
