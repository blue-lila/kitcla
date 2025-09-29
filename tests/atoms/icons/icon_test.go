package icons

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"testing"
)

func TestIconIcon(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Icons.Icon

	h := component.Icon("tree")
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/icon_icon_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/icon_icon_1.html")
}

func TestIconIconWithSize(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Icons.Icon

	h := component.IconWithSize("tree", "4")
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/icon_icon_with_size_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/icon_icon_with_size_1.html")
}
