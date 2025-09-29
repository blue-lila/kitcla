package inputs

import (
	"github.com/blue-lila/kitcla/aria"
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/goc"
)

type BooleanInput struct {
	Component *components.Component
}

type BooleanInputMod struct {
	Name  string
	Value bool
	Attr  goc.Attr
}

func (this *BooleanInput) BooleanInput(name string, value bool) goc.HTML {
	return this.H(&BooleanInputMod{Name: name, Value: value, Attr: goc.Attr{}})
}

func (this *BooleanInput) BooleanInputWithAttr(name string, value bool, attr goc.Attr) goc.HTML {
	return this.H(&BooleanInputMod{Name: name, Value: value, Attr: attr})
}

func (this *BooleanInput) H(mod *BooleanInputMod) goc.HTML {
	attr := mod.Attr
	if mod.Value {
		attr["checked"] = "checked"
	}
	attr["type"] = "checkbox"
	attr["value"] = "true"
	attr["name"] = mod.Name
	attr["class"] = "h-4 w-4 rounded border-scale-3 text-selection"
	attr["role"] = aria.AriaRoleCheckbox

	return this.Component.Cav("input",
		attr,
	)
}
