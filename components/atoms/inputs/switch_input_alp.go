package inputs

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/goc"
)

type SwitchInputAlp struct {
	Component *components.Component
}

type SwitchInputAlpMod struct {
	XModel string
}

func (this *SwitchInputAlp) SwitchInput(xModel string, mod *SwitchInputAlpMod) goc.HTML {
	mod = this.modDefaulting(mod)
	mod.XModel = xModel
	return this.H(mod)
}

func (this *SwitchInputAlp) modDefaulting(mod *SwitchInputAlpMod) *SwitchInputAlpMod {
	if mod != nil {
		return mod
	}
	return &SwitchInputAlpMod{}
}

func (this *SwitchInputAlp) H(mod *SwitchInputAlpMod) goc.HTML {
	return this.toggleDiv(mod)
}

func (this *SwitchInputAlp) toggleInput(mod *SwitchInputAlpMod) goc.HTML {
	return this.Component.Cav("input",
		goc.Attr{
			"type": "checkbox", "x-model": mod.XModel,
			"class": "checked:right-0 checked:bg-selection absolute block w-6 h-6 rounded-full bg-white border-scale-4 border-2 appearance-none cursor-pointer",
		},
	)
}

func (this *SwitchInputAlp) toggleContainer(mod *SwitchInputAlpMod) goc.HTML {
	return this.Component.Ca("div", goc.Attr{
		"class": "toggle-label block overflow-hidden h-6 rounded-full bg-scale-2 cursor-pointer border border-scale-3",
		"for":   "toggle",
	})
}

func (this *SwitchInputAlp) toggleDiv(mod *SwitchInputAlpMod) goc.HTML {
	return this.Component.Cas("label", goc.Attr{
		"class": "relative inline-block w-12 mr-2 align-middle select-none transition duration-200 ease-in",
	}, this.toggleInput(mod), this.toggleContainer(mod))
}
