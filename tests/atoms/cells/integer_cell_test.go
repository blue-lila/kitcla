package cells

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"testing"
)

func TestIntegerCellIntegerCell(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Cells.IntegerCell

	h := component.IntegerCell(0)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/integer_cell_integer_cell_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/integer_cell_integer_cell_1.html")
	sup.AddPage("IntegerCell", html)
}
