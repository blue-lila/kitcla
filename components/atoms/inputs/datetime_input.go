package inputs

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/goc"
)

type DatetimeInput struct {
	Component *components.Component
}

type DatetimeInputMod struct {
	Name  string
	Value string
}

func (this *DatetimeInput) DatetimeInput(name string, value string) goc.HTML {
	return this.H(&DatetimeInputMod{Name: name, Value: value})
}

func (this *DatetimeInput) H(mod *DatetimeInputMod) goc.HTML {
	h := this.Component.Cav("input",
		goc.Attr{"type": "datetime-local", "name": mod.Name, "value": mod.Value, "class": "border border-scale-3 px-3 py-2 rounded-md"},
	)
	return h
}
