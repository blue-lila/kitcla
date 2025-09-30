package inputs

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"testing"
)

func TestSwitchInputSwitchInput(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Inputs.SwitchInput

	h := component.SwitchInput("auto_watering", true)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/switch_input_switch_input_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/switch_input_switch_input_1.html")
}
