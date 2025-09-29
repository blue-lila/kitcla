package sidebars

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/goc"
)

type Sidebar struct {
	Component *components.Component
}

type SidebarItem struct {
	Label string
	Url   string
}

type SidebarMod struct {
	Items []*SidebarItem
}

func (this *Sidebar) Sidebar(items []*SidebarItem) goc.HTML {
	return this.H(&SidebarMod{Items: items})
}

func (this *Sidebar) H(mod *SidebarMod) goc.HTML {
	return goc.H("div",
		goc.Attr{"class": "flex-1 flex flex-col min-h-0 bg-scale-10"},
		this.LogoContainer(),
		this.LinkContainer(mod),
	)
}

func (this *Sidebar) LogoImg() goc.HTML {
	return this.Component.Ccv("div", "text-3xl ml-4", "🎋")
}

func (this *Sidebar) LogoTitle() goc.HTML {
	return this.Component.Cav("a", goc.Attr{"class": "text-scale-0 ml-4 text-xl", "href": "/app/"}, "Dashboard")
}

func (this *Sidebar) LogoContainer() goc.HTML {
	return goc.H("div",
		goc.Attr{"class": "flex items-center h-16 flex-shrink-0 px-4 bg-scale-10"},
		this.LogoImg(),
		this.LogoTitle(),
	)
}

func (this *Sidebar) MenuItemLink(item *SidebarItem) goc.HTML {
	return goc.H("a",
		goc.Attr{"href": item.Url, "class": "bg-scale-10 text-white group flex items-center px-2 py-2 text-sm font-medium rounded-md"},
		item.Label,
	)
}
func (this *Sidebar) MenuItems(mod *SidebarMod) goc.HTML {
	var items []goc.HTML

	for _, item := range mod.Items {
		items = append(items, this.MenuItemLink(item))
	}

	return goc.H("nav",
		goc.Attr{"class": "flex-1 px-2 py-4 space-y-1"},
		items,
	)
}
func (this *Sidebar) LinkContainer(mod *SidebarMod) goc.HTML {
	return goc.H("div",
		goc.Attr{"class": "flex-1 flex flex-col overflow-y-auto"},
		this.MenuItems(mod),
	)
}
