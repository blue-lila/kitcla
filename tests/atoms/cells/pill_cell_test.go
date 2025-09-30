package cells

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"testing"
)

func TestPillCellPillCell(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Cells.PillCell

	h := component.PillCell("Organic")
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/pill_cell_pill_cell_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/pill_cell_pill_cell_1.html")
	sup.AddPage("PillCell", html)
}
