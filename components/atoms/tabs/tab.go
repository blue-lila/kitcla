package tabs

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/goc"
)

type Tab struct {
	Component *components.Component
}

type TabMod struct {
	Items []*TabItem
}

type TabItem struct {
	Text   string
	Value  string
	Active bool
}

func (this *Tab) Tab(mod *TabMod) goc.HTML {
	return this.H(mod)
}

func (this *Tab) H(mod *TabMod) goc.HTML {
	return this.Component.Dcs("border-b border-gray-200", this.nav(mod))
}

func (this *Tab) nav(mod *TabMod) goc.HTML {
	if mod == nil {
		return goc.HTML{}
	}
	var set []goc.HTML
	for _, item := range mod.Items {
		set = append(set, this.button(item, mod))
	}
	return this.Component.Dcs("flex space-x-1", set...)
}

func (this *Tab) button(item *TabItem, mod *TabMod) goc.HTML {
	css := "py-4 px-1 inline-flex items-center gap-x-2 border-b-2 text-sm whitespace-nowrap" +
		" hover:text-blue-600 focus:outline-none focus:text-blue-600 disabled:opacity-50" +
		" disabled:pointer-events-none"

	if item.Active {
		css += " font-semibold border-blue-600 text-blue-600"
	} else {
		css += " border-transparent text-gray-500"
	}

	return this.Component.Cav("button", goc.Attr{"class": css}, item.Text)
}
