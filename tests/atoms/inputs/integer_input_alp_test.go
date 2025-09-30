package inputs

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"testing"
)

func TestIntegerInputAlpIntegerInput(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Inputs.IntegerInputAlp

	h := component.IntegerInput("plant_count")
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/integer_input_alp_integer_input_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/integer_input_alp_integer_input_1.html")
}

func TestIntegerInputAlpMini(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Inputs.IntegerInputAlp

	h := component.Mini("seed_count")
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/integer_input_alp_mini_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/integer_input_alp_mini_1.html")
}
