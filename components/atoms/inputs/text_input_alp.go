package inputs

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/goc"
)

type TextInputAlp struct {
	Component *components.Component
}

type TextInputAlpMod struct {
	XModel     string
	Attributes TextInputAlpAttributes
}

type TextInputAlpAttributes struct {
	H goc.Attr
}

func (this *TextInputAlp) TextInput(xModel string) goc.HTML {
	return this.H(&TextInputAlpMod{XModel: xModel})
}

func (this *TextInputAlp) H(mod *TextInputAlpMod) goc.HTML {
	attr := mod.Attributes.H
	if attr == nil {
		attr = goc.Attr{}
	}
	attr["type"] = "text"
	attr["x-model"] = mod.XModel
	attr["class"] = "border border-scale-3 px-3 py-2 rounded-md"

	return this.Component.Cav("input", attr)
}
