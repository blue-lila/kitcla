package inputs

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"github.com/blue-lila/kitcla/sup/placeholder"
	"testing"
)

func TestBooleanInputAlpBooleanInput(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Inputs.BooleanInputAlp

	h := component.BooleanInput("is_watered")
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/boolean_input_alp_boolean_input_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/boolean_input_alp_boolean_input_1.html")
}

func TestBooleanInputAlpBooleanInputWithAttr(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Inputs.BooleanInputAlp

	h := component.BooleanInputWithAttr("requires_fertilizer", placeholder.GocAttrValue())
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/boolean_input_alp_boolean_input_with_attr_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/boolean_input_alp_boolean_input_with_attr_1.html")
}
