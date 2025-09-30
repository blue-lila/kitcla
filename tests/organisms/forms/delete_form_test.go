package forms

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"testing"
)

func TestDeleteFormDeleteForm(t *testing.T) {
	kit := kitcla.New()
	component := kit.Organisms.Forms.DeleteForm

	h := component.DeleteForm("/plants/delete/1")
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/delete_form_delete_form_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/delete_form_delete_form_1.html")
}
