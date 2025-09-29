package tabs

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"testing"
)

func TestTabsTab(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Tabs.Tab

	h := component.Tab(nil)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/tabs_tab_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/tabs_tab_1.html")
}
