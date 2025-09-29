package cells

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"testing"
)

func TestRelationCellRelationCell(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Cells.RelationCell

	h := component.RelationCell("", "", "")
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/relation_cell_relation_cell_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/relation_cell_relation_cell_1.html")
	sup.AddPage("RelationCell", html)
}
