package buttons

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"testing"
)

func TestButtonAlpGhostlyTertiaryIconLink(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Buttons.ButtonAlp

	h := component.GhostlyTertiaryIconLink("", "", nil)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/button_alp_ghostly_tertiary_icon_link_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/button_alp_ghostly_tertiary_icon_link_1.html")
	sup.AddPage("GhostlyTertiaryIconLink", html)
}

func TestButtonAlpPrimaryLink(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Buttons.ButtonAlp

	h := component.PrimaryLink("", "", nil)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/button_alp_primary_link_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/button_alp_primary_link_1.html")
	sup.AddPage("PrimaryLink", html)
}

func TestButtonAlpSecondaryIconLink(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Buttons.ButtonAlp

	h := component.SecondaryIconLink("", "", nil)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/button_alp_secondary_icon_link_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/button_alp_secondary_icon_link_1.html")
	sup.AddPage("SecondaryIconLink", html)
}

func TestButtonAlpSecondaryLink(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Buttons.ButtonAlp

	h := component.SecondaryLink("", "", nil)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/button_alp_secondary_link_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/button_alp_secondary_link_1.html")
	sup.AddPage("SecondaryLink", html)
}

func TestButtonAlpTertiaryIconLink(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Buttons.ButtonAlp

	h := component.TertiaryIconLink("", "", nil)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/button_alp_tertiary_icon_link_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/button_alp_tertiary_icon_link_1.html")
	sup.AddPage("TertiaryIconLink", html)
}
