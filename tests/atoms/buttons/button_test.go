package buttons

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"testing"
)

func TestButtonH(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Buttons.Button
	mod := kit.Atoms.Buttons.Button.Mod()
	mod.Kind = "primary"
	mod.Label = "Test Button"
	mod.Size = "36"

	h := component.H(mod)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/button_h.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/button_h.html")
	sup.AddPage("H", html)
}

func TestButtonPrimaryLink(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Buttons.Button

	h := component.PrimaryLink("Plant it", "#", nil)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/button_primary_link_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/button_primary_link_1.html")
	sup.AddPage("PrimaryLink", html)
}

func TestButtonPrimarySubmit(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Buttons.Button

	h := component.PrimarySubmit("Plant it", nil)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/button_primary_submit_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/button_primary_submit_1.html")
	sup.AddPage("PrimarySubmit", html)
}

func TestButtonPrimaryPost(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Buttons.Button

	h := component.PrimaryPost("Plant it", "#", nil)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/button_primary_post_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/button_primary_post_1.html")
	sup.AddPage("PrimaryPost", html)
}

func TestButtonPrimaryIconLink(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Buttons.Button

	h := component.PrimaryIconLink("tree", "#", nil)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/button_primary_icon_link_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/button_primary_icon_link_1.html")
	sup.AddPage("PrimaryIconLink", html)
}

func TestButtonPrimaryIconSubmit(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Buttons.Button

	h := component.PrimaryIconSubmit("tree", nil)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/button_primary_icon_submit_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/button_primary_icon_submit_1.html")
	sup.AddPage("PrimaryIconSubmit", html)
}

func TestButtonPrimaryIconPost(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Buttons.Button

	h := component.PrimaryIconPost("tree", "#", nil)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/button_primary_icon_post_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/button_primary_icon_post_1.html")
	sup.AddPage("PrimaryIconPost", html)
}

func TestButtonSecondaryIconLink(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Buttons.Button

	h := component.SecondaryIconLink("tree", "#", nil)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/button_secondary_icon_link_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/button_secondary_icon_link_1.html")
	sup.AddPage("SecondaryIconLink", html)
}

func TestButtonSecondarySubmit(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Buttons.Button

	h := component.SecondarySubmit("Plant it", nil)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/button_secondary_submit_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/button_secondary_submit_1.html")
	sup.AddPage("SecondarySubmit", html)
}

func TestButtonSecondaryIconSubmit(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Buttons.Button

	h := component.SecondaryIconSubmit("tree", nil)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/button_secondary_icon_submit_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/button_secondary_icon_submit_1.html")
	sup.AddPage("SecondaryIconSubmit", html)
}

func TestButtonSecondaryIconPost(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Buttons.Button

	h := component.SecondaryIconPost("tree", "#", nil)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/button_secondary_icon_post_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/button_secondary_icon_post_1.html")
	sup.AddPage("SecondaryIconPost", html)
}

func TestButtonSecondaryLink(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Buttons.Button

	h := component.SecondaryLink("Plant it", "#", nil)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/button_secondary_link_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/button_secondary_link_1.html")
	sup.AddPage("SecondaryLink", html)
}

func TestButtonSecondaryPost(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Buttons.Button

	h := component.SecondaryPost("Plant it", "#", nil)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/button_secondary_post_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/button_secondary_post_1.html")
	sup.AddPage("SecondaryPost", html)
}

func TestButtonTertiaryIconLink(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Buttons.Button

	h := component.TertiaryIconLink("tree", "#", nil)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/button_tertiary_icon_link_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/button_tertiary_icon_link_1.html")
	sup.AddPage("TertiaryIconLink", html)
}
