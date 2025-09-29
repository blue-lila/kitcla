package inputs

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"testing"
)

func TestSelectInputSelectInput(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Inputs.SelectInput

	options := component.Mod()
	h := component.SelectInput("", "", options)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/select_input_select_input_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/select_input_select_input_1.html")
}
