package fields

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"github.com/blue-lila/kitcla/sup/placeholder"
	"testing"
)

func TestFieldField(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Fields.Field

	h := component.Field("Watering Instructions", placeholder.GocHtmlValue())
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/field_field_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/field_field_1.html")
}

func TestFieldHiddenField(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Fields.Field

	h := component.HiddenField("plant_species_id", placeholder.GocHtmlValue())
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/field_hidden_field_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/field_hidden_field_1.html")
}
