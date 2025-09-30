package cards

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/components/organisms/cards"
	"github.com/blue-lila/kitcla/goc"
	u "github.com/blue-lila/kitcla/goc_utils"
	"github.com/blue-lila/kitcla/sup"
	"github.com/blue-lila/kitcla/sup/placeholder"
	"testing"
)

func TestCardEmptyCard(t *testing.T) {
	kit := kitcla.New()
	component := kit.Organisms.Cards.Card

	h := component.EmptyCard(placeholder.CardContent())
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/card_empty_card_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/card_empty_card_1.html")

	sup.AddPage("EmptyCard", html)
}

func TestCardEmptyCardWithSmallMargins(t *testing.T) {
	kit := kitcla.New()
	component := kit.Organisms.Cards.Card

	h := component.EmptyCardWithSmallMargins(placeholder.CardContent())
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/card_empty_card_with_small_margins_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/card_empty_card_with_small_margins_1.html")

	sup.AddPage("EmptyCardWithSmallMargins", html)
}

func TestCardFormCard(t *testing.T) {
	kit := kitcla.New()
	component := kit.Organisms.Cards.Card

	h := component.FormCard(placeholder.CardContent())
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/card_form_card_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/card_form_card_1.html")

	sup.AddPage("FormCard", html)
}

func TestCardShowCard(t *testing.T) {
	kit := kitcla.New()
	component := kit.Organisms.Cards.Card

	h := component.ShowCard(placeholder.CardContent())
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/card_show_card_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/card_show_card_1.html")

	sup.AddPage("ShowCard", html)
}

func TestCardSimpleCard(t *testing.T) {
	kit := kitcla.New()
	component := kit.Organisms.Cards.Card

	h := component.SimpleCard("Plant Care Guide", placeholder.CardContent())
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/card_simple_card_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/card_simple_card_1.html")

	sup.AddPage("SimpleCard", html)
}

func TestCardSimpleCardWithButtons(t *testing.T) {
	kit := kitcla.New()
	component := kit.Organisms.Cards.Card

	h := component.SimpleCardWithButtons(
		placeholder.CardTitle(),
		placeholder.CardContent(),
		u.Ds(kit.Atoms.Buttons.Button.SecondaryLink("Plant it", "#", nil)))
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/card_simple_card_with_buttons_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/card_simple_card_with_buttons_1.html")

	sup.AddPage("SimpleCardWithButtons", html)
}

func TestCardSimpleCardWithStatus(t *testing.T) {
	kit := kitcla.New()
	component := kit.Organisms.Cards.Card

	h := component.SimpleCardWithStatus("Garden Health Report", placeholder.CardContent(), cards.StatusSuccess)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/card_simple_card_with_status_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/card_simple_card_with_status_1.html")

	sup.AddPage("SimpleCardWithStatus", html)
}
