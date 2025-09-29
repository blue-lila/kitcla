package inputs

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/goc"
)

type GridIdInput struct {
	Component *components.Component
}

type GridIdInputMod struct {
	Name    string
	Value   string
	Options []GridIdInputModOption
}

type GridIdInputModOption struct {
	Value string
	Text  string
}

func (this *GridIdInputMod) AddOption(value string, text string) {
	this.Options = append(this.Options, GridIdInputModOption{Value: value, Text: text})
}

func (this *GridIdInput) GridIdInput(name string, value string, mod *GridIdInputMod) goc.HTML {
	mod = this.modDefaulting(mod)
	mod.Name = name
	mod.Value = value
	return this.H(mod)
}

func (this *GridIdInput) modDefaulting(mod *GridIdInputMod) *GridIdInputMod {
	if mod == nil {
		return &GridIdInputMod{}
	}
	return mod
}

func (this *GridIdInput) H(mod *GridIdInputMod) goc.HTML {
	var set []goc.HTML

	for _, option := range mod.Options {
		set = append(set, this.gridItem(option, mod))
	}

	return this.Component.Dcs("grid grid-cols-3 gap-3", set...)
}

func (this *GridIdInput) Mod() *GridIdInputMod {
	return &GridIdInputMod{}
}

func (this *GridIdInput) gridItem(option GridIdInputModOption, mod *GridIdInputMod) goc.HTML {
	return this.Component.Ccs("label", "relative block cursor-pointer",
		this.htmlInput(option, mod),
		this.card(option, mod),
	)
}

func (this *GridIdInput) htmlInput(option GridIdInputModOption, mod *GridIdInputMod) goc.HTML {
	attrs := goc.Attr{
		"type":  "radio",
		"name":  mod.Name,
		"value": option.Value,
		"class": "peer sr-only",
	}
	if mod.Value == option.Value {
		attrs["checked"] = "checked"
	}
	return this.Component.Cas("input", attrs)
}

func (this *GridIdInput) card(option GridIdInputModOption, mod *GridIdInputMod) goc.HTML {
	return this.Component.Dcv("flex items-center justify-center h-16 rounded-lg border border-gray-200 bg-gray-50 text-sm font-medium text-gray-700 peer-checked:bg-blue-600 peer-checked:text-white peer-checked:border-blue-600 transition", option.Text)
}
