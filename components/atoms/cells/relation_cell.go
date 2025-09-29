package cells

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/goc"
)

type RelationCell struct {
	Component *components.Component
}

type RelationCellMod struct {
	Value string
	Text  string
	Url   string
}

func (this *RelationCell) RelationCell(value string, text string, url string) goc.HTML {
	return this.H(&RelationCellMod{Value: value, Text: text, Url: url})
}

func (this *RelationCell) H(mod *RelationCellMod) goc.HTML {
	return this.Component.Cav("a", goc.Attr{"href": mod.Url}, mod.Text)
}
