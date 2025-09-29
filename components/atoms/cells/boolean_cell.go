package cells

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/components/atoms/icons"
	"github.com/blue-lila/kitcla/goc"
)

type BooleanCell struct {
	Component *components.Component
	Icon      *icons.Icon
}

type BooleanCellMod struct {
	Value      bool
	TrueColor  string
	FalseColor string
}

func (this *BooleanCell) BooleanCell(value bool) goc.HTML {
	return this.H(&BooleanCellMod{Value: value, TrueColor: "success", FalseColor: "warning"})
}

func (this *BooleanCell) NeutralFalseBooleanCell(value bool) goc.HTML {
	return this.H(&BooleanCellMod{Value: value, TrueColor: "success", FalseColor: "scale-6"})
}

func (this *BooleanCell) H(mod *BooleanCellMod) goc.HTML {
	color := "text-" + mod.TrueColor
	// Embed values for css pruner: text-success, text-warning, text-scale-6
	icon2 := this.Icon.Icon(icons.IconFasCircleCheck)
	if mod.Value == false {
		color = "text-" + mod.FalseColor
		icon2 = this.Icon.Icon(icons.IconFasCircleXmark)

	}
	return this.Component.Ccs("div", "h-full flex items-center", this.Component.W(color, icon2))
}
