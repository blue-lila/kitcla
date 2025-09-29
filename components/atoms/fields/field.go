package fields

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/goc"
)

type Field struct {
	Component *components.Component
}

type FieldMod struct {
	Label  string
	Input  goc.HTML
	Hidden bool
}

func (this *Field) Field(label string, input goc.HTML) goc.HTML {
	return this.H(&FieldMod{Label: label, Input: input})
}

func (this *Field) HiddenField(label string, input goc.HTML) goc.HTML {
	return this.H(&FieldMod{Label: label, Input: input, Hidden: true})
}

func (this *Field) H(mod *FieldMod) goc.HTML {
	if mod.Hidden == true {
		return mod.Input
	}
	return this.Component.Ccs("div", "flex flex-col", this.label(mod.Label), mod.Input)
}

func (this *Field) label(label string) goc.HTML {
	return this.Component.Ccv("label", "font-bold", label)
}
