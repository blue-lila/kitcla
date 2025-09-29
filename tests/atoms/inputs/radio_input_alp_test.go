package inputs

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"testing"
)

func TestRadioInputAlpRadioInput(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Inputs.RadioInputAlp

	h := component.RadioInput(nil)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/radio_input_alp_radio_input_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/radio_input_alp_radio_input_1.html")
}
