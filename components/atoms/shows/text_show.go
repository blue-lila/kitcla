package shows

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/goc"
)

type TextShow struct {
	Component *components.Component
}

type TextShowMod struct {
	Name  string
	Value string
}

func (this *TextShow) TextShow(name string, value string) goc.HTML {
	return this.H(&TextShowMod{Name: name, Value: value})
}

func (this *TextShow) H(mod *TextShowMod) goc.HTML {
	return this.Component.Cv("div", mod.Value)
}
