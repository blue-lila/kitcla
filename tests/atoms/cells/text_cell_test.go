package cells

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"testing"
)

func TestTextCellTextCell(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Cells.TextCell

	h := component.TextCell("Cherry Tomatoes")
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/text_cell_text_cell_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/text_cell_text_cell_1.html")
	sup.AddPage("TextCell", html)
}
