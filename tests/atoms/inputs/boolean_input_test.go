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

	h := component.BooleanInput("", false)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/boolean_input_boolean_input_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/boolean_input_boolean_input_1.html")
}

func TestBooleanInputBooleanInputWithAttr(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Inputs.BooleanInput

	h := component.BooleanInputWithAttr("", false, placeholder.GocAttrValue())
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/boolean_input_boolean_input_with_attr_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/boolean_input_boolean_input_with_attr_1.html")
}
