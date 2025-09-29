package buttons

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/components/atoms/icons"
	"github.com/blue-lila/kitcla/goc"
)

type ButtonAlp struct {
	Icon      *icons.Icon
	Button    *Button
	Component *components.Component
}

type ButtonAlpMod struct {
	Label     string
	Kind      string
	OnClick   string
	HtmlKind  string
	Icon      string
	Post      bool
	Size      string
	TextColor string
	Disabled  bool
}

func (this *ButtonAlp) PrimaryLink(label string, onClick string, mod *ButtonAlpMod) goc.HTML {
	return this.H(&ButtonAlpMod{
		Label:    label,
		Kind:     ButtonModKindPrimary,
		Size:     ButtonModSizeLg,
		OnClick:  onClick,
		HtmlKind: "a",
	})
}

func (this *ButtonAlp) SecondaryLink(label string, onClick string, mod *ButtonAlpMod) goc.HTML {
	return this.H(&ButtonAlpMod{
		Label:    label,
		Kind:     ButtonModKindSecondary,
		Size:     ButtonModSizeLg,
		OnClick:  onClick,
		HtmlKind: "a",
	})
}

func (this *ButtonAlp) SecondaryIconLink(icon string, onClick string, mod *ButtonAlpMod) goc.HTML {
	return this.H(&ButtonAlpMod{
		Kind:     ButtonModKindSecondary,
		Size:     ButtonModSizeLg,
		OnClick:  onClick,
		Icon:     icon,
		HtmlKind: "a",
	})
}

func (this *ButtonAlp) TertiaryIconLink(icon string, onClick string, mod *ButtonAlpMod) goc.HTML {
	return this.H(&ButtonAlpMod{
		Kind:     ButtonModKindTertiary,
		Size:     ButtonModSizeLg,
		OnClick:  onClick,
		Icon:     icon,
		HtmlKind: "a",
	})
}

func (this *ButtonAlp) GhostlyTertiaryIconLink(icon string, onClick string, mod *ButtonAlpMod) goc.HTML {
	return this.H(&ButtonAlpMod{
		Kind:      ButtonModKindTertiary,
		Size:      ButtonModSizeLg,
		OnClick:   onClick,
		Icon:      icon,
		HtmlKind:  "a",
		TextColor: "text-gray-500 hover:text-gray-700",
	})
}

func (this *ButtonAlp) H(mod *ButtonAlpMod) goc.HTML {
	return this.linkButton(mod)
}

func (this *ButtonAlp) linkButton(mod *ButtonAlpMod) goc.HTML {
	converted := this.convert(mod)

	if mod.Icon != "" {
		return this.Component.Cas(
			"a",
			goc.Attr{
				"class":  this.Button.baseCss(converted),
				"@click": mod.OnClick},
			this.Icon.Icon(mod.Icon),
		)
	}

	return this.Component.Cav(
		"a",
		goc.Attr{"class": this.Button.baseCss(converted), "@click": mod.OnClick},
		mod.Label,
	)
}

func (this *ButtonAlp) convert(mod *ButtonAlpMod) *ButtonMod {
	return &ButtonMod{
		Label:     mod.Label,
		Kind:      mod.Kind,
		HtmlKind:  mod.HtmlKind,
		Icon:      mod.Icon,
		Post:      mod.Post,
		Size:      mod.Size,
		TextColor: mod.TextColor,
		Disabled:  mod.Disabled,
	}
}
