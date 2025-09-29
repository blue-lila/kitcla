package inputs

import (
	"fmt"
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/goc"

	"strconv"
)

type DecimalInput struct {
	Component *components.Component
}

type DecimalInputMod struct {
	Name      string
	Value     float64
	Precision int
}

func (this *DecimalInput) DecimalInput(name string, value float64) goc.HTML {
	return this.H(&DecimalInputMod{Name: name, Value: value, Precision: 2})
}

func (this *DecimalInput) H(mod *DecimalInputMod) goc.HTML {
	precisionS := strconv.Itoa(mod.Precision)
	format := "%." + precisionS + "f"

	v := fmt.Sprintf(format, mod.Value)
	return this.Component.Cav("input",
		goc.Attr{"type": "number", "step": "any", "name": mod.Name, "value": v, "class": "border border-scale-3 px-3 py-2 rounded-md"},
	)
}
