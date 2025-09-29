package navbars

import (
	"github.com/blue-lila/kitcla/aria"
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/goc"
)

type Navbar struct {
	Component *components.Component
}

type NavbarMod struct {
	Items []*NavbarItem
}

type NavbarItem struct {
	Label    string
	Selected bool
	Link     string
}

func (this *Navbar) H(mod *NavbarMod) goc.HTML {
	return this.Component.Cas("nav", goc.Attr{"class": "flex flex-row space-x-4", "role": aria.AriaRoleNavigation},
		this.items(mod)...,
	)
}

func (this *Navbar) navbarItem(mod *NavbarMod, item *NavbarItem) goc.HTML {
	css := "hover:text-gray-500 rounded px-2 py-1 cursor-pointer"
	if item.Selected {
		css += " bg-gray-200"
	}
	if item.Link != "" {
		return this.Component.Cav("a", goc.Attr{"href": item.Link, "class": css}, item.Label)
	}
	return this.Component.Dcv(css, item.Label)
}

func (this *Navbar) items(mod *NavbarMod) []goc.HTML {
	var set []goc.HTML

	for _, item := range mod.Items {
		set = append(set, this.navbarItem(mod, item))
	}
	return set
}
