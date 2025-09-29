package resources

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/goc"
)

type Resource struct {
	Component *components.Component
}

type ResourceMod struct {
	Js  string
	Css string
}

func (this *Resource) ResourceJsCss(js string, css string) goc.HTML {
	return this.H(&ResourceMod{Js: js, Css: css})
}

func (this *Resource) H(mod *ResourceMod) goc.HTML {
	var deps []goc.HTML
	if mod.Js != "" {
		deps = append(deps, this.Component.Ca("script", goc.Attr{"type": "text/javascript", "src": mod.Js}))
	}
	if mod.Js != "" {
		deps = append(deps, this.Component.Ca("link", goc.Attr{"rel": "stylesheet", "type": "text/css", "href": mod.Css}))
	}
	return this.Component.Cs("div", deps...)
}
