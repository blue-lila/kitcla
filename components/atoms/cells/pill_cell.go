package cells

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/goc"
)

type PillCell struct {
	Component *components.Component
}

type PillCellMod struct {
	Value string
	Color string
}

func (this *PillCell) PillCell(value string) goc.HTML {
	return this.H(&PillCellMod{Value: value})
}

func (this *PillCell) H(mod *PillCellMod) goc.HTML {
	if mod.Value == "" {
		return goc.HTML{}
	}
	return this.Component.Ccv("span", "py-1 px-3 bg-gray-100 rounded-full", mod.Value)
}
