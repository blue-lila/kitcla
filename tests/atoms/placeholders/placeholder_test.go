package placeholders

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"testing"
)

func TestPlaceholderPlaceholder(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Placeholders.Placeholder

	h := component.Placeholder("Label", goc.H("div"))
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/placeholder_placeholder_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/placeholder_placeholder_1.html")
}
