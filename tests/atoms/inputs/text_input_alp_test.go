package inputs

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"testing"
)

func TestTextInputAlpTextInput(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Inputs.TextInputAlp

	h := component.TextInput("garden_notes")
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/text_input_alp_text_input_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/text_input_alp_text_input_1.html")
}
