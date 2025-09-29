package modals

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"testing"
)

func TestModalSimpleModal(t *testing.T) {
	kit := kitcla.New()
	component := kit.Organisms.Modals.Modal

	h := component.SimpleModal()
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/modal_simple_modal_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/modal_simple_modal_1.html")

	sup.AddPage("SimpleModal", html)
}
