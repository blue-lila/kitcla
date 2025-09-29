package inputs

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/goc"
)

type SwitchInput struct {
	Component *components.Component
}

type SwitchInputMod struct {
	Name     string
	Value    bool
	Checkbox bool
}

func (this *SwitchInput) SwitchInput(name string, value bool) goc.HTML {
	return this.H(&SwitchInputMod{Name: name, Value: value})
}

func (this *SwitchInput) H(mod *SwitchInputMod) goc.HTML {
	checked := ""
	if mod.Value {
		checked = "checked"
	}
	return this.toggleDiv(mod, checked)
}

func (this *SwitchInput) toggleInput(mod *SwitchInputMod, checked string) goc.HTML {
	return this.Component.Cav("input",
		goc.Attr{
			"type": "checkbox", "name": mod.Name, checked: checked, "value": "true",
			"class": "checked:right-0 checked:bg-selection absolute block w-6 h-6 rounded-full bg-white border-scale-4 border-2 appearance-none cursor-pointer",
		},
	)
}

func (this *SwitchInput) toggleContainer(mod *SwitchInputMod) goc.HTML {
	return this.Component.Ca("div", goc.Attr{
		"class": "toggle-label block overflow-hidden h-6 rounded-full bg-scale-2 cursor-pointer border border-scale-3",
		"for":   "toggle",
	})
}

func (this *SwitchInput) toggleDiv(mod *SwitchInputMod, checked string) goc.HTML {
	return this.Component.Cas("label", goc.Attr{
		"class": "relative inline-block w-12 mr-2 align-middle select-none transition duration-200 ease-in",
	}, this.toggleInput(mod, checked), this.toggleContainer(mod))
}
