package cells

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/goc"
)

type LongTextCell struct {
	Component *components.Component
}

type LongTextCellMod struct {
	Value string
}

func (this *LongTextCell) LongTextCell(value string) goc.HTML {

	return this.H(&LongTextCellMod{Value: value})
}

func (this *LongTextCell) H(mod *LongTextCellMod) goc.HTML {
	return this.Component.Cav("a", goc.Attr{
		"class":  "underline",
		"@click": "showTextIntoModal"},
		"Open text")
}
