package cells

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/goc"
)

type TextCell struct {
	Component *components.Component
}

type TextCellMod struct {
	Value string
}

func (this *TextCell) TextCell(value string) goc.HTML {
	return this.H(&TextCellMod{Value: value})
}

func (this *TextCell) H(mod *TextCellMod) goc.HTML {
	return this.Component.Cv("span", mod.Value)
}
