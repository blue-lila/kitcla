package menus

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/components/molecules/navbars"
	"github.com/blue-lila/kitcla/goc"
)

type Menu struct {
	Navbar    *navbars.Navbar
	Component *components.Component
}

type MenuMod struct {
	SelectedKey string
	Items       []*MenuItem
}

type MenuItem struct {
	Key  string
	Name string
	Url  string
}

func (this *MenuMod) AddItem(key string, name string, url string) {
	this.Items = append(this.Items, &MenuItem{
		Key:  key,
		Name: name,
		Url:  url,
	})
}

func (this *Menu) Mod() *MenuMod {
	return &MenuMod{}
}

func (this *Menu) TopMenu(mod *MenuMod) goc.HTML {
	return this.H(mod)
}

func (this *Menu) H(mod *MenuMod) goc.HTML {
	navbarMod := this.convertToNavbarMod(mod)
	return this.Navbar.H(navbarMod)
}

func (this *Menu) convertToNavbarMod(mod *MenuMod) *navbars.NavbarMod {
	navbarMod := &navbars.NavbarMod{}
	for _, item := range mod.Items {
		selected := mod.SelectedKey == item.Key
		navbarItem := &navbars.NavbarItem{
			Label:    item.Name,
			Selected: selected,
			Link:     item.Url,
		}
		navbarMod.Items = append(navbarMod.Items, navbarItem)
	}
	return navbarMod
}
