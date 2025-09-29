package popovers

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"github.com/blue-lila/kitcla/sup/placeholder"
	"testing"
)

func TestPopoverGhostlyIconPopover(t *testing.T) {
	kit := kitcla.New()
	component := kit.Molecules.Popovers.Popover

	h := component.GhostlyIconPopover("", placeholder.GocHtmlValue())
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/popover_ghostly_icon_popover_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/popover_ghostly_icon_popover_1.html")
}

func TestPopoverIconPopover(t *testing.T) {
	kit := kitcla.New()
	component := kit.Molecules.Popovers.Popover

	h := component.IconPopover("", placeholder.GocHtmlValue())
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/popover_icon_popover_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/popover_icon_popover_1.html")
}

func TestPopoverIconPopoverWithFixedWidth(t *testing.T) {
	kit := kitcla.New()
	component := kit.Molecules.Popovers.Popover

	h := component.IconPopoverWithFixedWidth("", placeholder.GocHtmlValue(), "")
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/popover_icon_popover_with_fixed_width_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/popover_icon_popover_with_fixed_width_1.html")
}

func TestPopoverPopover(t *testing.T) {
	kit := kitcla.New()
	component := kit.Molecules.Popovers.Popover

	h := component.Popover(placeholder.GocHtmlValue(), placeholder.GocHtmlValue())
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/popover_popover_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/popover_popover_1.html")
}

func TestPopoverTextPopover(t *testing.T) {
	kit := kitcla.New()
	component := kit.Molecules.Popovers.Popover

	h := component.TextPopover("", placeholder.GocHtmlValue())
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/popover_text_popover_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/popover_text_popover_1.html")
}
