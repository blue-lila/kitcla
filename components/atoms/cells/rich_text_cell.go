package cells

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/goc"
)

type RichTextCell struct {
	Component *components.Component
}

type RichTextCellMod struct {
	Value string
}

func (this *RichTextCell) RichTextCell(value string) goc.HTML {
	return this.H(&RichTextCellMod{Value: value})
}

func (this *RichTextCell) H(mod *RichTextCellMod) goc.HTML {
	return this.Component.Cav("a", goc.Attr{
		"class":     "underline cursor-pointer",
		"data-text": string(mod.Value),
		// Using `onclick` instead of `@click`,because we may not be into
		// an x-data and alpinejs events works on in a alpine js object
		"onclick": "showTextIntoModal(this)"},
		"Open text")
}
