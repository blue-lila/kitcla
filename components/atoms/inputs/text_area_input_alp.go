package inputs

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/goc"
)

type TextAreaInputAlp struct {
	Component *components.Component
}

type TextAreaInputAlpMod struct {
	XModel     string
	Attributes TextAreaInputAlpAttributes
}

type TextAreaInputAlpAttributes struct {
	H goc.Attr
}

func (this *TextAreaInputAlp) TextAreaInput(xModel string, mod *TextAreaInputAlpMod) goc.HTML {
	mod = this.modDefaulting(mod)
	mod.XModel = xModel
	return this.H(mod)
}

func (this *TextAreaInputAlp) modDefaulting(mod *TextAreaInputAlpMod) *TextAreaInputAlpMod {
	if mod == nil {
		return &TextAreaInputAlpMod{}
	}
	return mod
}

func (this *TextAreaInputAlp) Mod() *TextAreaInputAlpMod {
	return &TextAreaInputAlpMod{}
}

func (this *TextAreaInputAlp) H(mod *TextAreaInputAlpMod) goc.HTML {
	attr := mod.Attributes.H
	if attr == nil {
		attr = goc.Attr{}
	}
	attr["x-model"] = mod.XModel
	attr["class"] = "border border-scale-3 px-3 py-2 rounded-md"

	return this.Component.Cav("textarea", attr)
}
