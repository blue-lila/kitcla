package inputs

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/goc"

	"strconv"
)

type HiddenInput struct {
	Component *components.Component
}

type HiddenInputMod struct {
	Name  string
	Value string
}

func (this *HiddenInput) HiddenInput(name string, value string) goc.HTML {
	return this.H(&HiddenInputMod{Name: name, Value: value})
}

func (this *HiddenInput) IntegerHiddenInput(name string, value int) goc.HTML {
	return this.H(&HiddenInputMod{Name: name, Value: strconv.Itoa(value)})
}

func (this *HiddenInput) H(mod *HiddenInputMod) goc.HTML {
	h := this.Component.Cav("input",
		goc.Attr{"type": "hidden", "name": mod.Name, "value": mod.Value, "hidden": "hidden"},
	)
	return h
}
