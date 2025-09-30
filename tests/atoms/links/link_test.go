package links

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"testing"
)

func TestLinkIconLink(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Links.Link

	h := component.IconLink("trowel", "/tools")
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/link_icon_link_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/link_icon_link_1.html")
}

func TestLinkLink(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Links.Link

	h := component.Link("Plant Library", "/plants")
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/link_link_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/link_link_1.html")
}

func TestLinkSubmitLink(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Links.Link

	h := component.SubmitLink("Save Garden Plan")
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/link_submit_link_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/link_submit_link_1.html")
}
