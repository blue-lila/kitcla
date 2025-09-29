package buttons

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/components/organisms/key_point"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"testing"
)

func TestKeyPointKeyPoint(t *testing.T) {
	kit := kitcla.New()
	component := kit.Organisms.KeyPoint.KeyPoint

	h := component.KeyPoint("Plants", "36", nil)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/key_point_key_point_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/key_point_key_point_1.html")
	sup.AddPage("KeyPoint", html)
}

func TestKeyPointKeyPointWithIcon(t *testing.T) {
	kit := kitcla.New()
	component := kit.Organisms.KeyPoint.KeyPoint

	h := component.KeyPointWithIcon("Plants", "36", "seedling", nil)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/key_point_key_point_with_icon_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/key_point_key_point_with_icon_1.html")
	sup.AddPage("KeyPointWithIcon", html)
}

func TestKeyPointKeyPointWithIndicator(t *testing.T) {
	kit := kitcla.New()
	component := kit.Organisms.KeyPoint.KeyPoint

	h := component.KeyPointWithIndicator("Plants", "36", key_point.IndicatorStateWarning, nil)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/key_point_key_point_with_indicator_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/key_point_key_point_with_indicator_1.html")
	sup.AddPage("KeyPointWithIndicator", html)
}

func TestKeyPointKeyPointWithOverValue(t *testing.T) {
	kit := kitcla.New()
	component := kit.Organisms.KeyPoint.KeyPoint

	h := component.KeyPointWithOverValue("Plants", "36", "100 trees", nil)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/key_point_key_point_with_over_value_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/key_point_key_point_with_over_value_1.html")
	sup.AddPage("KeyPointWithOverValue", html)
}

func TestKeyPointStyle1(t *testing.T) {
	kit := kitcla.New()
	component := kit.Organisms.KeyPoint.KeyPoint

	mod := component.Mod()
	mod.IndicatorState = key_point.IndicatorStateWarning
	mod.OverValue = "100 trees"
	mod.IconLabel = "seedling"
	h := component.KeyPoint("Plants", "36", mod)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/key_point_style1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/key_point_style1.html")
	sup.AddPage("Style1", html)
}
