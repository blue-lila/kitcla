package cells

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/goc"

	"strconv"
)

type IntegerCell struct {
	Component *components.Component
}

type IntegerCellMod struct {
	Value int
}

func (this *IntegerCell) IntegerCell(value int) goc.HTML {
	return this.H(&IntegerCellMod{Value: value})
}

func (this *IntegerCell) H(mod *IntegerCellMod) goc.HTML {
	v := strconv.Itoa(mod.Value)
	return this.Component.Cv("span", v)
}
