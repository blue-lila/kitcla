package cells

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"testing"
)

func TestLongTextCellLongTextCell(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Cells.LongTextCell

	h := component.LongTextCell("")
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/long_text_cell_long_text_cell_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/long_text_cell_long_text_cell_1.html")
	sup.AddPage("LongTextCell", html)
}
