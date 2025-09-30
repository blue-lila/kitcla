package inputs

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"testing"
)

func TestAdvancedSelectInputAdvancedSelectInput(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Inputs.AdvancedSelectInput

	h := component.AdvancedSelectInput("plant_variety", "heirloom_tomato", "vegetable_varieties")
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/advanced_select_input_advanced_select_input_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/advanced_select_input_advanced_select_input_1.html")
}

func TestAdvancedSelectInputAdvancedSelectInput2(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Inputs.AdvancedSelectInput

	h := component.AdvancedSelectInput2("seed_supplier", "organic_seeds_co", "seed_suppliers", nil)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/advanced_select_input_advanced_select_input2_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/advanced_select_input_advanced_select_input2_1.html")
}
