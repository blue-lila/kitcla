package inputs

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/goc"
)

type BooleanInputAlp struct {
	Component *components.Component
}

type BooleanInputAlpMod struct {
	XModel string
	Attr   goc.Attr
}

func (this *BooleanInputAlp) BooleanInput(xModel string) goc.HTML {
	return this.H(&BooleanInputAlpMod{XModel: xModel, Attr: goc.Attr{}})
}

func (this *BooleanInputAlp) BooleanInputWithAttr(xModel string, attr goc.Attr) goc.HTML {
	return this.H(&BooleanInputAlpMod{XModel: xModel, Attr: attr})
}

func (this *BooleanInputAlp) H(mod *BooleanInputAlpMod) goc.HTML {
	attr := mod.Attr
	attr["type"] = "checkbox"
	attr["x-model"] = mod.XModel
	attr["class"] = "h-4 w-4 rounded border-scale-3 text-selection"

	return this.Component.Cav("input",
		attr,
	)
}
