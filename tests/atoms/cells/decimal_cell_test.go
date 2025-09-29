package cells

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"testing"
)

func TestDecimalCellDecimalCell(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Cells.DecimalCell

	h := component.DecimalCell(0.0)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/decimal_cell_decimal_cell_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/decimal_cell_decimal_cell_1.html")
	sup.AddPage("DecimalCell", html)
}
