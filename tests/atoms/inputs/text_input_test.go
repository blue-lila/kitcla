package inputs

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"testing"
)

func TestTextInputTextInput(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Inputs.TextInput

	h := component.TextInput("plant_name", "Cherry Tomatoes")
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/text_input_text_input_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/text_input_text_input_1.html")
}
