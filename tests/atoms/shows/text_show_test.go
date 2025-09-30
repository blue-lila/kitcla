package shows

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"testing"
)

func TestTextShowTextShow(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Shows.TextShow

	h := component.TextShow("Plant Status", "Healthy Growth")
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/text_show_text_show_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/text_show_text_show_1.html")
}
