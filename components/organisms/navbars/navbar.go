package navbars

import (
	"github.com/blue-lila/kitcla/aria"
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/components/atoms/icons"
	"github.com/blue-lila/kitcla/components/atoms/links"
	"github.com/blue-lila/kitcla/goc"
)

type Navbar struct {
	Icon      *icons.Icon
	Link      *links.Link
	Component *components.Component
}

type NavbarMod struct {
	Items []*links.LinkMod
}

func (this *Navbar) Navbar(item []*links.LinkMod) goc.HTML {
	return this.H(&NavbarMod{
		Items: item,
	})
}

func (this *Navbar) H(mod *NavbarMod) goc.HTML {
	var subs []goc.HTML
	for i, item := range mod.Items {
		subs = append(subs, this.item(item))
		if i != len(mod.Items)-1 {
			subs = append(subs, this.Icon.IconWithSize(icons.IconChrevronRight, "3"))
		}
	}

	return this.Component.Cas("nav", goc.Attr{"class": "flex text-scale-7 items-center font-medium space-x-4", "role": aria.AriaRoleNavigation},
		subs...,
	)
}

func (this *Navbar) item(item *links.LinkMod) goc.HTML {
	item.HoverCss = "hover:text-scale-10"
	return this.Component.W("flex items-center", this.Link.H(item))
}
