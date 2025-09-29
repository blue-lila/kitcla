package inputs

import (
	"github.com/blue-lila/kitcla/aria"
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/goc"
)

type TextInput struct {
	Component *components.Component
}

type TextInputMod struct {
	Name  string
	Value string
}

func (this *TextInput) TextInput(name string, value string) goc.HTML {
	return this.H(&TextInputMod{Name: name, Value: value})
}

func (this *TextInput) H(mod *TextInputMod) goc.HTML {
	return this.Component.Cav("input",
		goc.Attr{"type": "text", "name": mod.Name, "value": mod.Value, "class": "border border-scale-3 px-3 py-2 rounded-md", "role": aria.AriaRoleTextbox},
	)
}
