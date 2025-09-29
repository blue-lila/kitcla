package inputs

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/goc"
)

type SelectInputAlp struct {
	Component *components.Component
}

type SelectInputAlpMod struct {
	XModel     string
	OnChange   string
	Options    []*SelectOption
	SelectAttr goc.Attr
}

func (this *SelectInputAlpMod) AddOption(value string, text string) {
	option := &SelectOption{
		Value: value,
		Text:  text,
	}
	this.Options = append(this.Options, option)
}

func (this *SelectInputAlpMod) AddOptionWithData(value string, text string, data string) {
	option := &SelectOption{
		Value: value,
		Text:  text,
		Data:  data,
	}
	this.Options = append(this.Options, option)
}

func (this *SelectInputAlp) SelectInputAlp(xModel string, mod *SelectInputAlpMod) goc.HTML {
	mod.XModel = xModel
	return this.H(mod)
}

func (this *SelectInputAlp) SelectInputAlpWithOnChange(xModel string, onChange string, mod *SelectInputAlpMod) goc.HTML {
	mod.OnChange = onChange
	mod.XModel = xModel
	return this.H(mod)
}

func (this *SelectInputAlp) Mod() *SelectInputAlpMod {
	return &SelectInputAlpMod{}
}

func (this *SelectInputAlp) H(mod *SelectInputAlpMod) goc.HTML {
	attr := mod.SelectAttr
	if attr == nil {
		attr = goc.Attr{}
	}
	attr["x-model"] = mod.XModel
	if mod.OnChange != "" {
		attr["@change"] = mod.OnChange
	}
	attr["class"] = "border border-scale-3 px-3 py-2 bg-scale-0 rounded-md"
	attr["style"] = "height:42px;"
	return this.Component.Cas("select",
		attr,
		this.options(mod)...,
	)
}

func (this *SelectInputAlp) options(mod *SelectInputAlpMod) []goc.HTML {
	var options []goc.HTML
	for _, option := range mod.Options {
		options = append(options, this.option(option, ""))
	}
	return options
}

func (this *SelectInputAlp) option(option *SelectOption, value string) goc.HTML {
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
