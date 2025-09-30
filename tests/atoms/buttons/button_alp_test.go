package buttons

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"testing"
)

func TestButtonAlpQuaternaryIconLink(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Buttons.ButtonAlp

	h := component.QuaternaryIconLink("leaf", "view_plant_details()", nil)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/button_alp_quaternary_icon_link_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/button_alp_quaternary_icon_link_1.html")
	sup.AddPage("QuaternaryIconLink", html)
}

func TestButtonAlpPrimaryLink(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Buttons.ButtonAlp

	h := component.PrimaryLink("Add to Garden", "add_plant_to_garden()", nil)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/button_alp_primary_link_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/button_alp_primary_link_1.html")
	sup.AddPage("PrimaryLink", html)
}

func TestButtonAlpSecondaryIconLink(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Buttons.ButtonAlp

	h := component.SecondaryIconLink("droplet", "schedule_watering()", nil)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/button_alp_secondary_icon_link_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/button_alp_secondary_icon_link_1.html")
	sup.AddPage("SecondaryIconLink", html)
}

func TestButtonAlpSecondaryLink(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Buttons.ButtonAlp

	h := component.SecondaryLink("Care Guide", "view_care_instructions()", nil)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/button_alp_secondary_link_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/button_alp_secondary_link_1.html")
	sup.AddPage("SecondaryLink", html)
}

func TestButtonAlpTertiaryIconLink(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Buttons.ButtonAlp

	h := component.TertiaryIconLink("sun", "set_sunlight_schedule()", nil)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/button_alp_tertiary_icon_link_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/button_alp_tertiary_icon_link_1.html")
	sup.AddPage("TertiaryIconLink", html)
}
