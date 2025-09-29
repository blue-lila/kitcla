package dropdowns

import (
	"github.com/blue-lila/kitcla/aria"
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/components/atoms/icons"
	"github.com/blue-lila/kitcla/goc"
)

type Dropdown struct {
	Icon      *icons.Icon
	Component *components.Component
}

type DropdownMod struct {
	Button goc.HTML
	Items  []goc.HTML
}

func (this *Dropdown) Dropdown(button goc.HTML, items []goc.HTML) goc.HTML {
	return this.H(&DropdownMod{Button: button, Items: items})
}

func (this *Dropdown) EllipsisDropdown(items []goc.HTML) goc.HTML {
	button := this.Component.Ccs("div",
		"hover:text-scale-10 inline-flex h-6 w-6 items-center justify-center border border-scale-5 rounded-full cursor-pointer text-scale-7",
		this.Icon.Icon(icons.IconFasEllipsisVertical),
	)
	return this.H(&DropdownMod{Button: button, Items: items})
}

func (this *Dropdown) H(mod *DropdownMod) goc.HTML {
	return this.Component.Cas("div",
		goc.Attr{
			"class":       "flex relative w-8 h-8",
			"x-data":      "{open:false}",
			"@click.away": "open = false",
		},
		this.buttonContainer(mod),
		this.menuContainer(mod),
	)
}

func (this *Dropdown) buttonContainer(mod *DropdownMod) goc.HTML {
	return this.Component.Cas("div",
		goc.Attr{
			"class":                "w-8 h-8 flex items-center justify-center",
			"@click":               "open = !open",
			"role":                 aria.AriaRoleButton,
			aria.AriaHaspopup:      aria.AriaHaspopupMenu,
			aria.AriaExpanded:      "false",
			"x-bind:aria-expanded": "open",
		},
		mod.Button,
	)
}

func (this *Dropdown) menuContainer(mod *DropdownMod) goc.HTML {
	return this.Component.Cas("div",
		goc.Attr{
			"class":  "absolute z-30 mt-8 px-2 -left-48 -top-8 w-48",
			"x-show": "open",
			"@click": "open = !open",
		},
		this.menu(mod),
	)
}

func (this *Dropdown) menuItem(item goc.HTML) goc.HTML {
	return this.Component.Cas("div", goc.Attr{"class": "bg-scale-0 py-2 px-3 space-y-2 hover:bg-scale-2", "role": aria.AriaRoleMenuitem}, item)
}

func (this *Dropdown) menu(mod *DropdownMod) goc.HTML {
	var menuItem []goc.HTML

	for _, item := range mod.Items {
		menuItem = append(menuItem, this.menuItem(item))
	}

	return this.Component.Cas("div",
		goc.Attr{
			"class": "rounded-lg shadow-lg ring-1 ring-black ring-opacity-5 overflow-hidden",
			"role":  aria.AriaRoleMenu,
		},
		menuItem...,
	)
}
