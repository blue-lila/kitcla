package card_wrappers

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"github.com/blue-lila/kitcla/sup/placeholder"
	"testing"
)

func TestCardWrapperCardWrapper(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.CardWrapper.CardWrapper

	h := component.CardWrapper(placeholder.GocHtmlValue())
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/card_wrapper_card_wrapper_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/card_wrapper_card_wrapper_1.html")
}

func TestCardWrapperCardWrapperWithHeader(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.CardWrapper.CardWrapper

	h := component.CardWrapperWithHeader(placeholder.GocHtmlValue(), placeholder.GocHtmlValue())
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/card_wrapper_card_wrapper_with_header_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/card_wrapper_card_wrapper_with_header_1.html")
}
