package navmenus

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"testing"
)

func TestButtonH(t *testing.T) {
	kit := kitcla.New()
	component := kit.Organisms.Navmenus.Navmenu
	mod := component.Mod()

	h := component.H(mod)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/navmenu_h.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/navmenu_h.html")
	sup.AddPage("H", html)
}
