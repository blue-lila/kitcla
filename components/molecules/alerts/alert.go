package alerts

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/components/atoms/icons"
	"github.com/blue-lila/kitcla/goc"
)

type Alert struct {
	Icon      *icons.Icon
	Component *components.Component
}

const AlertModKindSuccess = "success"
const AlertModKindInfo = "info"
const AlertModKindDanger = "danger"
const AlertModKindWarning = "warning"

type AlertMod struct {
	Kind        string
	WithIcon    bool
	Title       string
	Description string
}

func (this *Alert) SuccessAlert(title string, description string, mod *AlertMod) goc.HTML {
	return this.generateAlert(AlertModKindSuccess, title, description, true, mod)
}

func (this *Alert) WarningAlert(title string, description string, mod *AlertMod) goc.HTML {
	return this.generateAlert(AlertModKindWarning, title, description, true, mod)
}

func (this *Alert) DangerAlert(title string, description string, mod *AlertMod) goc.HTML {
	return this.generateAlert(AlertModKindDanger, title, description, true, mod)
}

func (this *Alert) InfoAlert(title string, description string, mod *AlertMod) goc.HTML {
	return this.generateAlert(AlertModKindInfo, title, description, true, mod)
}

func (this *Alert) generateAlert(kind string, title string, description string, withIcon bool, mod *AlertMod) goc.HTML {
	mod = this.modDefaulting(mod)
	mod.Kind = kind
	mod.Title = title
	mod.Description = description
	mod.WithIcon = withIcon
	return this.H(mod)
}

func (this *Alert) modDefaulting(mod *AlertMod) *AlertMod {
	if mod == nil {
		mod = &AlertMod{}
	}
	return mod
}

func (this *Alert) ModWithDescription(description string) *AlertMod {
	mod := this.modDefaulting(nil)
	mod.Description = description
	return mod
}

func (this *Alert) H(mod *AlertMod) goc.HTML {
	var set []goc.HTML

	if mod.WithIcon {
		set = append(set, this.iconPart(mod))
	}
	set = append(set, this.textualPart(mod))

	return this.Component.Dcs("border text-sm rounded-lg p-4 "+this.coloration(mod),
		this.Component.Dcs("flex", set...),
	)
}

func (this *Alert) coloration(mod *AlertMod) string {
	if mod.Kind == AlertModKindWarning {
		return "bg-yellow-50 border-yellow-200 text-yellow-800"
	}
	if mod.Kind == AlertModKindSuccess {
		return "bg-teal-100 border-teal-200 text-teal-800"
	}
	if mod.Kind == AlertModKindDanger {
		return "bg-red-100 border-red-200 text-red-800"
	}
	if mod.Kind == AlertModKindInfo {
		return "bg-blue-100 border-blue-200 text-blue-800"
	}
	panic("Unknown alert kind coloration")
}

func (this *Alert) icon(mod *AlertMod) goc.HTML {
	if mod.Kind == AlertModKindWarning {
		return this.Icon.Icon(icons.IconFasCircleExclamation)
	}
	if mod.Kind == AlertModKindSuccess {
		return this.Icon.Icon(icons.IconFasCircleCheck)
	}
	if mod.Kind == AlertModKindDanger {
		return this.Icon.Icon(icons.IconFasCircleXmark)
	}
	if mod.Kind == AlertModKindInfo {
		return this.Icon.Icon(icons.IconFasCircleInfo)
	}
	panic("Unknown alert kind icon")
}

func (this *Alert) textualPart(mod *AlertMod) goc.HTML {
	return this.Component.Dcs("ms-4", this.title(mod), this.description(mod))
}

func (this *Alert) iconPart(mod *AlertMod) goc.HTML {
	return this.Component.Dcs("skrink-0 mt-0.5", this.icon(mod))
}

func (this *Alert) title(mod *AlertMod) goc.HTML {
	return this.Component.Ccv("h3", "text-sm font-semibold", mod.Title)
}

func (this *Alert) description(mod *AlertMod) goc.HTML {
	return this.Component.Dcv("mt-1 text-sm", mod.Description)
}
