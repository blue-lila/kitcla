package cells

import (
	"fmt"
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/goc"

	"strconv"
)

type DecimalCell struct {
	Component *components.Component
}

type DecimalCellMod struct {
	Value     float64
	Precision int
}

func (this *DecimalCell) DecimalCell(value float64) goc.HTML {
	return this.H(&DecimalCellMod{Value: value, Precision: 2})
}

func (this *DecimalCell) H(mod *DecimalCellMod) goc.HTML {
	precisionS := strconv.Itoa(mod.Precision)
	format := "%." + precisionS + "f"
	s := fmt.Sprintf(format, mod.Value)
	return this.Component.Cv("span", s)
}
