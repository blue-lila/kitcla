package cells

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"testing"
)

func TestBooleanCellBooleanCell(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Cells.BooleanCell

	h := component.BooleanCell(false)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/boolean_cell_boolean_cell_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/boolean_cell_boolean_cell_1.html")
	sup.AddPage("BooleanCell", html)

}

func TestBooleanCellNeutralFalseBooleanCell(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Cells.BooleanCell

	h := component.NeutralFalseBooleanCell(false)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/boolean_cell_neutral_false_boolean_cell_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/boolean_cell_neutral_false_boolean_cell_1.html")
	sup.AddPage("NeutralFalseBooleanCell", html)
}
