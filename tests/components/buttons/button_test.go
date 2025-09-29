package buttons

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"testing"
)

func TestButtonPrimaryLink(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Buttons.Button

	h := component.PrimaryLink("My Label", "test.com", nil)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/button_primary_link_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/button_primary_link_1.html")
}
