package links

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/components/atoms/icons"
	"github.com/blue-lila/kitcla/goc"
)

type Link struct {
	Icon      *icons.Icon
	Component *components.Component
}

type LinkMod struct {
	Label    string
	Url      string
	Icon     string
	HoverCss string
	Submit   bool
}

func (this *Link) Link(label string, url string) goc.HTML {
	return this.H(&LinkMod{
		Label: label,
		Url:   url,
	})
}

func (this *Link) SubmitLink(label string) goc.HTML {
	return this.H(&LinkMod{
		Label:  label,
		Submit: true,
	})
}

func (this *Link) IconLink(icon string, url string) goc.HTML {
	return this.H(&LinkMod{
		Icon: icon,
		Url:  url,
	})
}

func (this *Link) H(mod *LinkMod) goc.HTML {
	css := mod.HoverCss
	if mod.Icon != "" && mod.Label == "" {
		return this.Component.Cas("a", goc.Attr{"href": mod.Url, "class": css}, this.Icon.Icon(mod.Icon))
	}
	if mod.Submit == true {
		return this.Component.Cav("button", goc.Attr{"type": "submit", "class": css}, mod.Label)
	}

	return this.Component.Cav("a", goc.Attr{"href": mod.Url, "class": css}, mod.Label)
}
