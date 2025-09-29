package inputs

import (
	"github.com/blue-lila/kitcla/aria"
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/goc"
)

type SelectInput struct {
	Component *components.Component
}

type SelectOption struct {
	Value string
	Text  string
	Data  string
}

type SelectInputMod struct {
	Name       string
	Value      string
	Options    []*SelectOption
	SelectAttr goc.Attr
}

func (this *SelectInputMod) AddOption(value string, text string) {
	option := &SelectOption{
		Value: value,
		Text:  text,
	}
	this.Options = append(this.Options, option)
}

func (this *SelectInputMod) AddOptionWithData(value string, text string, data string) {
	option := &SelectOption{
		Value: value,
		Text:  text,
		Data:  data,
	}
	this.Options = append(this.Options, option)
}

func (this *SelectInput) SelectInput(name string, value string, mod *SelectInputMod) goc.HTML {
	mod.Name = name
	mod.Value = value
	return this.H(mod)
}

func (this *SelectInput) Mod() *SelectInputMod {
	return &SelectInputMod{}
}

func (this *SelectInput) H(mod *SelectInputMod) goc.HTML {
	attr := mod.SelectAttr
	if attr == nil {
		attr = goc.Attr{}
	}
	attr["name"] = mod.Name
	attr["value"] = mod.Value
	attr["class"] = "border border-scale-3 px-3 py-2 bg-scale-0 rounded-md"
	attr["style"] = "height:42px;"
	attr["role"] = aria.AriaRoleCombobox
	return this.Component.Cas("select",
		attr,
		this.options(mod)...,
	)
}

func (this *SelectInput) options(mod *SelectInputMod) []goc.HTML {
	var options []goc.HTML
	for _, option := range mod.Options {
		options = append(options, this.option(option, mod.Value))
	}
	return options
}

func (this *SelectInput) option(option *SelectOption, value string) goc.HTML {
	values := goc.Attr{"value": option.Value}
	if option.Value == value {
		values["selected"] = "selected"
	}
	if option.Data != "" {
		values["data-options"] = option.Data
	}
	return this.Component.Cav("option",
		values,
		option.Text,
	)
}
