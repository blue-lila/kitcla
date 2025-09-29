package cells

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"testing"
	"time"
)

func TestTimeCellTimeCell(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Cells.TimeCell

	now := time.Date(2012, time.February, 4, 22, 21, 20, 0, time.UTC)
	h := component.TimeCell(now)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/time_cell_time_cell_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/time_cell_time_cell_1.html")
	sup.AddPage("TimeCell", html)
}
