package inputs

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/goc"

	"math/rand"
	"strconv"
)

type RadioInputAlp struct {
	Component *components.Component
}

type RadioInputAlpMod struct {
	XModel      string
	Items       []*RadioItem
	Orientation string
	OnChange    string
	Name        string
}

type RadioItem struct {
	Label string
	Value string
	Name  string
}

func (this *RadioInputAlpMod) AddRadioItem(label string, value string, name string) {
	item := &RadioItem{Label: label, Value: value, Name: name}
	if name == "" {
		item.Name = strconv.Itoa(rand.Int())
	}
	this.Items = append(this.Items, item)
}

func (this *RadioInputAlp) Mod(xModel string, name string) *RadioInputAlpMod {
	return &RadioInputAlpMod{
		XModel: xModel,
		Name:   name,
	}
}

func (this *RadioInputAlp) RadioInput(mod *RadioInputAlpMod) goc.HTML {
	return this.H(mod)
}

func (this *RadioInputAlp) H(mod *RadioInputAlpMod) goc.HTML {
	css := "flex flex-row space-x-6"
	if mod != nil && mod.Orientation == "vertical" {
		css = "flex flex-col space-y-6"
	}

	return this.Component.Dcs(css, this.radios(mod)...)
}

func (this *RadioInputAlp) input(mod *RadioInputAlpMod, value *RadioItem) goc.HTML {
	css := "shrink-0 mt-0.5 border-gray-200 rounded-full text-blue-600 focus:ring-blue-500 " +
		"disabled:opacity-50 disabled:pointer-events-none"

	attr := goc.Attr{
		"type":    "radio",
		"x-model": mod.XModel,
		"class":   css,
		"id":      value.Name,
		"value":   value.Value,
		"name":    mod.Name,
	}
	if mod.OnChange != "" {
		attr["@change"] = mod.OnChange
	}
	return this.Component.Ca("input", attr)
}

func (this *RadioInputAlp) radios(mod *RadioInputAlpMod) []goc.HTML {
	var items []goc.HTML
	if mod == nil {
		return items
	}
	for _, value := range mod.Items {
		items = append(items, this.radio(mod, value))
	}
	return items
}

func (this *RadioInputAlp) radio(mod *RadioInputAlpMod, value *RadioItem) goc.HTML {
	return this.Component.Dcs("flex", this.input(mod, value), this.label(mod, value))
}

func (this *RadioInputAlp) label(mod *RadioInputAlpMod, value *RadioItem) goc.HTML {
	css := "text-sm text-gray-500 ms-2"
	return this.Component.Cav("label",
		goc.Attr{"for": value.Name, "class": css},
		value.Label,
	)
}
