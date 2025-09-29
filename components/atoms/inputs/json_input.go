package inputs

import (
	"encoding/json"
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/goc"
)

type JsonInput struct {
	Component *components.Component
}

type JsonInputMod struct {
	Name  string
	Value json.RawMessage
}

func (this *JsonInput) JsonInput(name string, value json.RawMessage) goc.HTML {
	return this.H(&JsonInputMod{Name: name, Value: value})
}

func (this *JsonInput) H(mod *JsonInputMod) goc.HTML {
	return this.Component.Cav("input",
		goc.Attr{"type": "text", "name": mod.Name, "value": string(mod.Value), "class": "border border-scale-3 px-3 py-2"},
	)
}
