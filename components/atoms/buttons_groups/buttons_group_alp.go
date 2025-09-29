package buttons_groups

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/goc"
)

type ButtonsGroupAlp struct {
	Component *components.Component
}

type ButtonsGroupAlpMod struct {
	SelectedKey string
	Buttons     []*Button
}

func (this *ButtonsGroupAlpMod) AddButton(key string, label string, onClick string) {
	this.Buttons = append(this.Buttons, &Button{
		Key:     key,
		Label:   label,
		OnClick: onClick,
	})
}

type Button struct {
	Key     string
	Label   string
	OnClick string
}

func (this *ButtonsGroupAlp) Mod(selectedKey string) *ButtonsGroupAlpMod {
	return &ButtonsGroupAlpMod{
		SelectedKey: selectedKey,
	}
}

func (this *ButtonsGroupAlp) H(mod *ButtonsGroupAlpMod) goc.HTML {
	return this.group(mod)
}

func (this *ButtonsGroupAlp) group(mod *ButtonsGroupAlpMod) goc.HTML {
	return this.Component.Das(
		goc.Attr{"class": "inline-flex rounded-lg", "x-data": "{selectedKey: '" + mod.SelectedKey + "'}"},
		this.buttons(mod)...,
	)
}

func (this *ButtonsGroupAlp) buttons(mod *ButtonsGroupAlpMod) []goc.HTML {
	var buttons []goc.HTML

	for _, button := range mod.Buttons {
		buttons = append(buttons, this.button(mod, button))
	}
	return buttons
}

func (this *ButtonsGroupAlp) button(mod *ButtonsGroupAlpMod, button *Button) goc.HTML {
	css := "py-3 px-4 inline-flex items-center gap-x-2 "
	css += "-ms-px first:rounded-s-lg first:ms-0 last:rounded-e-lg text-sm font-medium focus:z-10 "
	css += "border border-gray-200 text-gray-800 shadow-sm hover:bg-gray-50 focus:outline-none "
	css += "focus:bg-gray-50 disabled:opacity-50 disabled:pointer-events-none"

	bindClass := "selectedKey == '" + button.Key + "' ? 'bg-gray-50' : 'bg-white'"

	return this.Component.Cav("button", goc.Attr{
		"@click": "selectedKey = '" + button.Key + "' ;" + button.OnClick,
		"class":  css,
		":class": bindClass,
	}, button.Label)
}
