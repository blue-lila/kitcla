package card_headers

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/goc"
)

type CardHeader struct {
	Component *components.Component
}

type CardHeaderMod struct {
	LeftArea goc.HTML
	Title    goc.HTML
}

func (this *CardHeader) CardHeader(title goc.HTML) goc.HTML {
	return this.H(&CardHeaderMod{
		Title: title,
	})
}

func (this *CardHeader) CardHeaderWithLeftArea(title goc.HTML, leftArea goc.HTML) goc.HTML {
	return this.H(&CardHeaderMod{
		LeftArea: leftArea,
		Title:    title,
	})
}

func (this *CardHeader) H(mod *CardHeaderMod) goc.HTML {
	return this.Component.Dcs("border-b p-4 flex flex-row justify-between",
		mod.Title,
		mod.LeftArea,
	)
}
