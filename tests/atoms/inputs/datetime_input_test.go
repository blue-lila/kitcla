package inputs

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"testing"
)

func TestDatetimeInputDatetimeInput(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Inputs.DatetimeInput

	h := component.DatetimeInput("", "")
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/datetime_input_datetime_input_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/datetime_input_datetime_input_1.html")
}
