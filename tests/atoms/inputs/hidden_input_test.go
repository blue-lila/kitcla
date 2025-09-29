package inputs

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"testing"
)

func TestHiddenInputHiddenInput(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Inputs.HiddenInput

	h := component.HiddenInput("", "")
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/hidden_input_hidden_input_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/hidden_input_hidden_input_1.html")
}

func TestHiddenInputIntegerHiddenInput(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Inputs.HiddenInput

	h := component.IntegerHiddenInput("", 0)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/hidden_input_integer_hidden_input_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/hidden_input_integer_hidden_input_1.html")
}
