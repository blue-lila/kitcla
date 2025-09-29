package card_wrappers

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/goc"
)

type CardWrapper struct {
	Component *components.Component
}

type CardWrapperMod struct {
	Header goc.HTML
	Body   goc.HTML
	Footer goc.HTML
}

func (this *CardWrapper) CardWrapperWithHeader(header goc.HTML, body goc.HTML) goc.HTML {
	mod := &CardWrapperMod{
		Header: header,
		Body:   body,
	}
	return this.H(mod)
}

func (this *CardWrapper) CardWrapper(body goc.HTML) goc.HTML {
	mod := &CardWrapperMod{
		Body: body,
	}
	return this.H(mod)
}

func (this *CardWrapper) H(mod *CardWrapperMod) goc.HTML {
	return this.Component.Dcs("flex flex-col bg-white border shadow-sm rounded-xl", mod.Header, mod.Body, mod.Footer)
}
