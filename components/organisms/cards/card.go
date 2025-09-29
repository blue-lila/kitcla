package cards

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/goc"
)

type Card struct {
	Component *components.Component
}

type CardMod struct {
	Title           string
	HasHeader       bool
	HasButtonArea   bool
	FullBodyControl bool
	Content         goc.HTML
	ButtonArea      goc.HTML
	Status          string
}

const StatusSuccess = "success"
const StatusWarning = "warning"
const StatusError = "error"
const StatusNoData = "no_data"

func (this *Card) FormCard(content goc.HTML) goc.HTML {
	return this.H(&CardMod{Content: content})
}

func (this *Card) ShowCard(content goc.HTML) goc.HTML {
	return this.H(&CardMod{Content: content})
}

func (this *Card) EmptyCard(content goc.HTML) goc.HTML {
	return this.H(&CardMod{Content: content,
		FullBodyControl: true})
}

func (this *Card) EmptyCardWithSmallMargins(content goc.HTML) goc.HTML {
	return this.H(&CardMod{Content: this.Component.W("p-4", content),
		FullBodyControl: true})
}

func (this *Card) SimpleCard(title string, content goc.HTML) goc.HTML {
	return this.H(&CardMod{
		Title:           title,
		Content:         content,
		HasHeader:       true,
		FullBodyControl: true,
	})
}

func (this *Card) SimpleCardWithButtons(title string, content goc.HTML, buttonArea goc.HTML) goc.HTML {
	return this.H(&CardMod{
		Title:           title,
		Content:         content,
		HasHeader:       true,
		FullBodyControl: true,
		HasButtonArea:   true,
		ButtonArea:      buttonArea,
	})
}

func (this *Card) SimpleCardWithStatus(title string, content goc.HTML, status string) goc.HTML {
	return this.H(&CardMod{
		Title:           title,
		Content:         content,
		HasHeader:       true,
		FullBodyControl: true,
		Status:          status,
	})
}

func (this *Card) H(mod *CardMod) goc.HTML {
	return this.outer(mod)
}

func (this *Card) outer(mod *CardMod) goc.HTML {
	margins := ""
	if !mod.FullBodyControl {
		margins = " px-10 py-7"
	}

	status := mod.Status
	base := "w-full bg-white border border-gray-200 rounded-lg shadow-sm"
	statusCss := ""

	switch status {
	case StatusSuccess:
		statusCss += " border-t-4 border-t-green-600"
	case StatusWarning:
		statusCss += " border-t-4 border-t-yellow-600"
	case StatusError:
		statusCss += " border-t-4 border-t-red-600"
	case StatusNoData:
		statusCss += " border-t-4 border-t-gray-600"
	}

	return this.Component.Dcs(base+margins+statusCss,
		this.header(mod),
		this.body(mod),
	)
}

func (this *Card) body(mod *CardMod) goc.HTML {
	return mod.Content
}

func (this *Card) header(mod *CardMod) goc.HTML {
	if mod.HasHeader == false {
		return goc.HTML{}
	}
	var items []goc.HTML
	items = append(items, this.title(mod))
	if mod.HasButtonArea == true {
		items = append(items, mod.ButtonArea)
	}
	return this.Component.Dcs("px-3 py-4 flex justify-between items-center border-b border-gray-200",
		items...,
	)
}

func (this *Card) title(mod *CardMod) goc.HTML {
	return this.Component.Ccv("h2", "text-lg font-semibold text-gray-800", mod.Title)
}
