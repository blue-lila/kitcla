package resources

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"testing"
)

func TestResourceResourceJsCss(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Resources.Resource

	h := component.ResourceJsCss("", "")
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/resource_resource_js_css_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/resource_resource_js_css_1.html")
}
