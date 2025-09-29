package shows

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"testing"
)

func TestRichTextShowRichTextShow(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Shows.RichTextShow

	h := component.RichTextShow("", "")
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/rich_text_show_rich_text_show_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/rich_text_show_rich_text_show_1.html")
}
