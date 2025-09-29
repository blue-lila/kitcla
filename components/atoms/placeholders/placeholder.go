package placeholders

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/goc"
)

type Placeholder struct {
	Component *components.Component
}

type PlaceholderMod struct {
	Label  string
	Input  goc.HTML
	Hidden bool
}

func (this *Placeholder) Placeholder(label string, input goc.HTML) goc.HTML {
	return this.H(&PlaceholderMod{})
}

func (this *Placeholder) H(mod *PlaceholderMod) goc.HTML {
	return this.Component.Dv("Hello world")
}
