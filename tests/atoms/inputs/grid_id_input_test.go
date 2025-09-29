package inputs

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"testing"
)

func TestButtonH(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Inputs.GridIdInput
	mod := component.Mod()

	h := component.H(mod)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/grid_id_input_h.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/grid_id_input_h.html")
	sup.AddPage("H", html)
}
