package inputs

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/goc"

	"strconv"
)

type IntegerInput struct {
	Component *components.Component
}

type IntegerInputMod struct {
	Name  string
	Value int
}

func (this *IntegerInput) IntegerInput(name string, value int) goc.HTML {
	return this.H(&IntegerInputMod{Name: name, Value: value})
}

func (this *IntegerInput) H(mod *IntegerInputMod) goc.HTML {
	v := strconv.Itoa(mod.Value)
	return this.Component.Cav("input",
		goc.Attr{"type": "number", "name": mod.Name, "value": v, "class": "border border-scale-3 px-3 py-2 rounded-md"},
	)
}
