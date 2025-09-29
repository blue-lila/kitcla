package cells

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"testing"
)

func TestJsonCellJsonCell(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Cells.JsonCell

	h := component.JsonCell([]byte("Hello world"))
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/json_cell_json_cell_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/json_cell_json_cell_1.html")
	sup.AddPage("JsonCell", html)
}
