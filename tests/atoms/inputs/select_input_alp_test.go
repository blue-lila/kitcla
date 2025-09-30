package inputs

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"testing"
)

func TestSelectInputAlpSelectInputAlp(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Inputs.SelectInputAlp

	options := component.Mod()
	h := component.SelectInputAlp("garden_section", options)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/select_input_alp_select_input_alp_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/select_input_alp_select_input_alp_1.html")
}

func TestSelectInputAlpSelectInputAlpWithOnChange(t *testing.T) {
	//kit := kitcla.New()
	//component := kit.Atoms.Inputs.SelectInputAlp
	//
	//h := component.SelectInputAlpWithOnChange("", "", "")
	//html := goc.RenderRoot(h)
	//
	//sup.UpdateEqualHtmlFromDataFile(t, html, "./data/select_input_alp_select_input_alp_with_on_change_1.html")
	//sup.AssertEqualHtmlFromDataFile(t, html, "./data/select_input_alp_select_input_alp_with_on_change_1.html")
}
