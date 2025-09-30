package headers

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"testing"
)

func TestHeaderH1(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Headers.Header

	h := component.H1("My Garden Dashboard")
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/header_h1_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/header_h1_1.html")
}

func TestHeaderH2(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Headers.Header

	h := component.H2("Plant Care Center")
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/header_h2_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/header_h2_1.html")
}
