package cells

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/goc"
)

type JsonCell struct {
	Component *components.Component
}

type JsonCellMod struct {
	Value []byte
}

func (this *JsonCell) JsonCell(value []byte) goc.HTML {
	return this.H(&JsonCellMod{Value: value})
}

func (this *JsonCell) H(mod *JsonCellMod) goc.HTML {
	return this.Component.Cav("a", goc.Attr{
		"class":     "underline cursor-pointer",
		"data-text": string(mod.Value),
		// Using `onclick` instead of `@click`,because we may not be into
		// an x-data and alpinejs events works on in a alpine js object
		"onclick": "showTextIntoModal(this)"},
		"Open text")
}
