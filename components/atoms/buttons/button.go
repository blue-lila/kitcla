package buttons

import (
	"github.com/blue-lila/kitcla/aria"
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/components/atoms/icons"
	"github.com/blue-lila/kitcla/goc"
)

type Button struct {
	Icon      *icons.Icon
	Component *components.Component
}

const ButtonModKindPrimary = "primary"
const ButtonModKindSecondary = "secondary"
const ButtonModKindTertiary = "tertiary"
const ButtonModKindQuaternary = "quaternary"
const ButtonModKindLink = "link"

const ButtonModSizeXs = "24"
const ButtonModSizeSm = "28"
const ButtonModSizeBase = "32"
const ButtonModSizeLg = "36"
const ButtonModSizeXl = "40"

const HtmlKindA = "a"
const HtmlKindSubmit = "submit"

type ButtonMod struct {
	Label         string
	Kind          string
	Link          string
	HtmlKind      string
	Icon          string
	Post          bool
	PostedValues  []*PostValue
	Size          string
	TextColor     string
	Disabled      bool
	Title         string
	RemoteFormUrl string
}

type PostValue struct {
	Name  string
	Value string
}

func (this *ButtonMod) AddPostValue(name string, value string) {
	this.PostedValues = append(this.PostedValues, &PostValue{
		Name:  name,
		Value: value,
	})
}

func (this *Button) Mod() *ButtonMod {
	return this.modDefaulting(nil)
}

func (this *Button) modDefaulting(mod *ButtonMod) *ButtonMod {
	if mod != nil {
		return mod
	}
	return &ButtonMod{}
}

func (this *Button) PrimaryLink(label string, link string, mod *ButtonMod) goc.HTML {
	return this.H(&ButtonMod{
		Label:    label,
		Kind:     ButtonModKindPrimary,
		Size:     ButtonModSizeLg,
		Link:     link,
		HtmlKind: HtmlKindA,
	})
}

func (this *Button) SecondaryLink(label string, link string, mod *ButtonMod) goc.HTML {
	mod = this.modDefaulting(mod)
	mod.Label = label
	mod.Kind = ButtonModKindSecondary
	mod.Size = ButtonModSizeLg
	mod.Link = link
	mod.HtmlKind = HtmlKindA
	return this.H(mod)
}

func (this *Button) SecondaryIconLink(icon string, link string, mod *ButtonMod) goc.HTML {
	mod = this.modDefaulting(mod)
	mod.Kind = ButtonModKindSecondary
	mod.Size = ButtonModSizeLg
	mod.Link = link
	mod.Icon = icon
	mod.HtmlKind = HtmlKindA
	return this.H(mod)
}

func (this *Button) SecondarySubmit(label string, mod *ButtonMod) goc.HTML {
	mod = this.modDefaulting(mod)
	mod.Label = label
	mod.Kind = ButtonModKindSecondary
	mod.Size = ButtonModSizeLg
	mod.HtmlKind = HtmlKindSubmit
	return this.H(mod)
}

func (this *Button) SecondaryIconSubmit(icon string, mod *ButtonMod) goc.HTML {
	mod = this.modDefaulting(mod)
	mod.Kind = ButtonModKindSecondary
	mod.Size = ButtonModSizeLg
	mod.Icon = icon
	mod.HtmlKind = HtmlKindSubmit
	return this.H(mod)
}

func (this *Button) SecondaryIconPost(icon string, link string, mod *ButtonMod) goc.HTML {
	mod = this.modDefaulting(mod)

	mod.Link = link
	mod.Icon = icon
	mod.Kind = ButtonModKindSecondary
	mod.Size = ButtonModSizeLg
	mod.HtmlKind = HtmlKindSubmit
	mod.Post = true

	return this.H(mod)
}

func (this *Button) TableIconLink(icon string, link string, mod *ButtonMod) goc.HTML {
	return this.H(&ButtonMod{
		Kind:     ButtonModKindSecondary,
		Size:     ButtonModSizeBase,
		Link:     link,
		Icon:     icon,
		HtmlKind: HtmlKindA,
	})
}

func (this *Button) TableLink(label string, link string, mod *ButtonMod) goc.HTML {
	return this.H(&ButtonMod{
		Kind:     ButtonModKindSecondary,
		Size:     ButtonModSizeBase,
		Link:     link,
		Label:    label,
		HtmlKind: HtmlKindA,
	})
}

func (this *Button) TertiaryIconLink(icon string, link string, mod *ButtonMod) goc.HTML {
	return this.H(&ButtonMod{
		Kind:     ButtonModKindTertiary,
		Size:     ButtonModSizeLg,
		Link:     link,
		Icon:     icon,
		HtmlKind: HtmlKindA,
	})
}

func (this *Button) FormSubmit(label string, mod *ButtonMod) goc.HTML {
	mod = this.modDefaulting(mod)
	mod.Label = label
	mod.Kind = ButtonModKindPrimary
	mod.Size = ButtonModSizeXl
	mod.HtmlKind = HtmlKindSubmit
	return this.H(mod)
}

func (this *Button) RemoteForm(label string, remoteFormUrl string, mod *ButtonMod) goc.HTML {
	mod = this.modDefaulting(mod)
	mod.Label = label
	mod.Kind = ButtonModKindSecondary
	mod.Size = ButtonModSizeLg
	mod.HtmlKind = HtmlKindA
	mod.RemoteFormUrl = remoteFormUrl
	return this.H(mod)
}

func (this *Button) PrimarySubmit(label string, mod *ButtonMod) goc.HTML {
	return this.H(&ButtonMod{
		Label:    label,
		Kind:     ButtonModKindPrimary,
		Size:     ButtonModSizeLg,
		HtmlKind: HtmlKindSubmit,
	})
}

func (this *Button) PrimaryPost(label string, link string, mod *ButtonMod) goc.HTML {
	mod = this.modDefaulting(mod)

	mod.Label = label
	mod.Link = link
	mod.Kind = ButtonModKindPrimary
	mod.Size = ButtonModSizeLg
	mod.HtmlKind = HtmlKindSubmit
	mod.Post = true

	return this.H(mod)
}

func (this *Button) PrimaryIconLink(icon string, link string, mod *ButtonMod) goc.HTML {
	return this.H(&ButtonMod{
		Kind:     ButtonModKindPrimary,
		Size:     ButtonModSizeLg,
		Link:     link,
		Icon:     icon,
		HtmlKind: HtmlKindA,
	})
}

func (this *Button) PrimaryIconSubmit(icon string, mod *ButtonMod) goc.HTML {
	return this.H(&ButtonMod{
		Kind:     ButtonModKindPrimary,
		Size:     ButtonModSizeLg,
		Icon:     icon,
		HtmlKind: HtmlKindSubmit,
	})
}

func (this *Button) PrimaryIconPost(icon string, link string, mod *ButtonMod) goc.HTML {
	mod = this.modDefaulting(mod)

	mod.Link = link
	mod.Icon = icon
	mod.Kind = ButtonModKindPrimary
	mod.Size = ButtonModSizeLg
	mod.HtmlKind = HtmlKindSubmit
	mod.Post = true

	return this.H(mod)
}

func (this *Button) SecondaryPost(label string, link string, mod *ButtonMod) goc.HTML {
	mod = this.modDefaulting(mod)

	mod.Label = label
	mod.Link = link
	mod.Kind = ButtonModKindSecondary
	mod.Size = ButtonModSizeLg
	mod.HtmlKind = HtmlKindSubmit
	mod.Post = true

	return this.H(mod)
}

func (this *Button) H(mod *ButtonMod) goc.HTML {
	if mod.Post == true {
		return this.postButton(mod)
	}
	if mod.HtmlKind == HtmlKindSubmit {
		return this.submitButton(mod)
	}
	return this.linkButton(mod)
}

func (this *Button) submitButton(mod *ButtonMod) goc.HTML {
	attributes := goc.Attr{"type": "submit", "class": this.baseCss(mod), "role": aria.AriaRoleButton}
	if mod.Disabled {
		attributes["disabled"] = ""
		attributes[aria.AriaDisabled] = aria.AriaDisabledTrue
	}
	attributes = this.addCommonAttributes(mod, attributes)

	if mod.Icon != "" {
		return this.Component.Cas(
			"button",
			attributes,
			this.Icon.Icon(mod.Icon),
		)
	}

	return this.Component.Cav(
		"button",
		attributes,
		mod.Label,
	)
}

func (this *Button) linkButton(mod *ButtonMod) goc.HTML {
	attributes := goc.Attr{"class": this.baseCss(mod), "role": aria.AriaRoleButton}
	if mod.Link != "" {
		attributes["href"] = mod.Link
	}

	if mod.Disabled {
		attributes["href"] = "#"
		attributes[aria.AriaDisabled] = aria.AriaDisabledTrue
	}

	attributes = this.addCommonAttributes(mod, attributes)

	if mod.Icon != "" {
		return this.Component.Cas(
			"a",
			attributes,
			this.Icon.Icon(mod.Icon),
		)
	}

	return this.Component.Cav(
		"a",
		attributes,
		mod.Label,
	)
}

func (this *Button) postButton(mod *ButtonMod) goc.HTML {
	return this.Component.Cas("form", goc.Attr{"action": mod.Link, "method": "POST", "role": aria.AriaRoleForm}, this.postValues(mod), this.submitButton(mod))
}

func (this *Button) baseCss(mod *ButtonMod) string {
	return this.cssFromKind(mod) + " " + this.cssFromSize(mod) + " " + this.cssFromState(mod)
}

func (this *Button) cssFromKind(mod *ButtonMod) string {
	base := "inline-flex items-center font-medium "

	pointer := "cursor-pointer"
	if mod.Disabled && mod.HtmlKind == HtmlKindA {
		pointer = "cursor-not-allowed"
	}
	base += pointer + " "

	switch mod.Kind {
	case ButtonModKindPrimary: // Solid
		return base + "border rounded-lg border-transparent bg-blue-600 text-white hover:bg-blue-700"
	case ButtonModKindSecondary: // Light
		color := "text-gray-800 "
		if mod.Icon != "" {
			color = "text-gray-700 "
		}
		return base + color + "rounded-lg border border-gray-200 bg-white shadow-sm hover:bg-gray-50"
	case ButtonModKindTertiary: // Ghost
		textColor := this.textColor(mod, " text-blue-600 hover:text-blue-800")
		return base + "rounded-lg border border-transparent hover:bg-blue-100 " + textColor
	case ButtonModKindQuaternary: // Soft
		return base + "rounded-lg border border-transparent bg-blue-100 text-blue-800 hover:bg-blue-200"
	default:
		panic("Unknown kind")
	}
}

func (this *Button) textColor(mod *ButtonMod, def string) string {
	if mod.TextColor != "" {
		return mod.TextColor
	}
	return def
}

func (this *Button) cssFromSize(mod *ButtonMod) string {
	switch mod.Size {
	case ButtonModSizeXs:
		return "text-xs h-6 px-2 py-1" // h=24px
	case ButtonModSizeSm:
		return "text-sm h-7 px-2 py-1" // h=28px
	case ButtonModSizeBase:
		return "px-2.5 py-1.5 text-sm h-8" // h=32px
	case ButtonModSizeLg:
		return "px-3 py-2 text-base h-9" // h=36px
	case ButtonModSizeXl:
		return "px-4 py-2.5 text-base h-10" // h=40px
	default:
		panic("Unknown size")
	}
}

func (this *Button) postValues(mod *ButtonMod) goc.HTML {
	if len(mod.PostedValues) == 0 {
		return goc.HTML{}
	}
	var values []goc.HTML
	for _, postedValue := range mod.PostedValues {
		values = append(values, this.Component.Cas("input", goc.Attr{"type": "hidden", "name": postedValue.Name, "value": postedValue.Value}))
	}
	return this.Component.Cs("span", values...)
}

func (this *Button) cssFromState(mod *ButtonMod) string {
	css := ""
	if mod.Disabled && mod.HtmlKind == HtmlKindSubmit {
		css = "disabled:opacity-50 disabled:cursor-not-allowed"
	}
	if mod.Disabled && mod.HtmlKind == HtmlKindA {
		css = "opacity-50"
	}
	return css
}

func (this *Button) addCommonAttributes(mod *ButtonMod, attributes goc.Attr) goc.Attr {
	if mod.Title != "" {
		attributes["title"] = mod.Title
	}
	if mod.RemoteFormUrl != "" {
		attributes["onclick"] = "G_getRemoteForm('" + mod.RemoteFormUrl + "')"
	}
	// Add aria-label for icon-only buttons
	if mod.Icon != "" && mod.Label == "" && mod.Title != "" {
		attributes[aria.AriaLabel] = mod.Title
	}
	return attributes
}
