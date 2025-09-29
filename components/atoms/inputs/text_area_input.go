package inputs

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/goc"
	"strconv"
)

type TextAreaInput struct {
	Component *components.Component
}

type TextAreaInputMod struct {
	Name         string
	Value        string
	ColumnsCount int
	RowsCount    int
	Placeholder  string
}

func (this *TextAreaInput) TextAreaInput(name string, value string, mod *TextAreaInputMod) goc.HTML {
	mod = this.modDefaulting(mod)
	mod.Name = name
	mod.Value = value
	return this.H(mod)
}

func (this *TextAreaInput) modDefaulting(mod *TextAreaInputMod) *TextAreaInputMod {
	if mod == nil {
		return &TextAreaInputMod{
			RowsCount: 15,
		}
	}
	return mod
}

func (this *TextAreaInput) Mod() *TextAreaInputMod {
	return this.modDefaulting(nil)
}

func (this *TextAreaInput) H(mod *TextAreaInputMod) goc.HTML {
	rowsCount := strconv.Itoa(mod.RowsCount)

	attrs := goc.Attr{"name": mod.Name, "class": "border border-scale-3 px-3 py-2 rounded-md", "rows": rowsCount}

	if mod.Placeholder != "" {
		attrs["placeholder"] = mod.Placeholder
	}

	return this.Component.Cav("textarea",
		attrs,
		mod.Value,
	)
}
