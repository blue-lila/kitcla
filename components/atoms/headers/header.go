package headers

import (
	"strconv"

	"github.com/blue-lila/kitcla/aria"
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/goc"
)

type Header struct {
	Component *components.Component
}

type HeaderMod struct {
	Label string
	Level int
}

func (this *Header) H1(label string) goc.HTML {
	return this.H(&HeaderMod{
		Label: label,
		Level: 1,
	})
}

func (this *Header) H2(label string) goc.HTML {
	return this.H(&HeaderMod{
		Label: label,
		Level: 2,
	})
}

func (this *Header) css(mod *HeaderMod) string {
	if mod.Level == 1 {
		return "text-xl font-semibold text-gray-800"
	}
	if mod.Level == 2 {
		return "text-lg font-semibold text-gray-800"
	}
	if mod.Level == 3 {
		return "text-base font-semibold text-gray-800"
	}
	panic("Unknown header level")
}

func (this *Header) H(mod *HeaderMod) goc.HTML {
	class := this.css(mod)
	el := "h" + strconv.Itoa(mod.Level)

	return this.Component.Cav(
		el,
		goc.Attr{
			"class":        class,
			"role":         aria.AriaRoleHeading,
			aria.AriaLevel: strconv.Itoa(mod.Level),
		},
		mod.Label,
	)
}
