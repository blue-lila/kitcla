package inputs

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"testing"
)

func TestIntegerInputIntegerInput(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Inputs.IntegerInput

	h := component.IntegerInput("", 0)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/integer_input_integer_input_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/integer_input_integer_input_1.html")
}
