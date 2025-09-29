package card_bodies

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"testing"
)

func TestDuoCardBodyDuoCardBody(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.CardBodies.DuoCardBody

	h := component.DuoCardBody(nil, nil, nil)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/duo_card_body_duo_card_body_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/duo_card_body_duo_card_body_1.html")
}
