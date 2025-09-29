package cells

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"testing"
)

func TestRichTextCellRichTextCell(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Cells.RichTextCell

	h := component.RichTextCell("")
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/rich_text_cell_rich_text_cell_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/rich_text_cell_rich_text_cell_1.html")
	sup.AddPage("RichTextCell", html)

}
