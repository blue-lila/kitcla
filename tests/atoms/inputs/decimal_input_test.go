package inputs

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"testing"
)

func TestDecimalInputDecimalInput(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Inputs.DecimalInput

	h := component.DecimalInput("", 0.0)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/decimal_input_decimal_input_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/decimal_input_decimal_input_1.html")
}
