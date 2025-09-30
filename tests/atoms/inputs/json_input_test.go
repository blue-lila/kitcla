package inputs

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"testing"
)

func TestJsonInputJsonInput(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Inputs.JsonInput

	h := component.JsonInput("plant_care_instructions", []byte(`{"watering": "daily", "sunlight": "partial", "fertilizer": "monthly"}`))
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/json_input_json_input_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/json_input_json_input_1.html")
}
