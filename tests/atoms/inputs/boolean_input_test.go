package inputs

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"github.com/blue-lila/kitcla/sup/placeholder"
	"testing"
)

func TestBooleanInputBooleanInput(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Inputs.BooleanInput

	h := component.BooleanInput("needs_watering", true)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/boolean_input_boolean_input_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/boolean_input_boolean_input_1.html")
}

func TestBooleanInputBooleanInputWithAttr(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Inputs.BooleanInput

	h := component.BooleanInputWithAttr("organic_certified", true, placeholder.GocAttrValue())
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/boolean_input_boolean_input_with_attr_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/boolean_input_boolean_input_with_attr_1.html")
}
