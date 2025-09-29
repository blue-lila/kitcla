package key_point

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/components/atoms/icons"
	"github.com/blue-lila/kitcla/goc"
)

type KeyPoint struct {
	Icon      *icons.Icon
	Component *components.Component
}

type KeyPointMod struct {
	Label          string
	IconLabel      string
	Value          string
	Text           string
	SubText        string
	OverValue      string
	IndicatorState string
	IconValue      string
}

const IndicatorStateSuccess = "success"
const IndicatorStateWarning = "warning"
const IndicatorStateError = "error"
const IndicatorStateNoData = "no_data"

func (this *KeyPoint) Mod() *KeyPointMod {
	return this.modDefaulting(nil)
}

func (this *KeyPoint) modDefaulting(mod *KeyPointMod) *KeyPointMod {
	if mod != nil {
		return mod
	}
	return &KeyPointMod{}
}

func (this *KeyPoint) KeyPoint(label string, value string, mod *KeyPointMod) goc.HTML {
	mod = this.modDefaulting(mod)
	mod.Label = label
	mod.Value = value
	return this.H(mod)
}

func (this *KeyPoint) KeyPointWithIcon(label string, value string, icon string, mod *KeyPointMod) goc.HTML {
	mod = this.modDefaulting(mod)
	mod.Label = label
	mod.Value = value
	mod.IconLabel = icon
	return this.H(mod)
}

func (this *KeyPoint) KeyPointWithOverValue(label string, value string, overValue string, mod *KeyPointMod) goc.HTML {
	mod = this.modDefaulting(mod)
	mod.Label = label
	mod.Value = value
	mod.OverValue = overValue
	return this.H(mod)
}

func (this *KeyPoint) KeyPointWithIndicator(label string, value string, indicatorState string, mod *KeyPointMod) goc.HTML {
	mod = this.modDefaulting(mod)
	mod.Label = label
	mod.Value = value
	mod.IndicatorState = indicatorState
	return this.H(mod)
}

func (this *KeyPoint) H(mod *KeyPointMod) goc.HTML {
	return this.Component.Dcs("flex flex-col", this.label(mod), this.value(mod))
}

func (this *KeyPoint) label(mod *KeyPointMod) goc.HTML {
	var set []goc.HTML

	if mod.IconLabel != "" {
		set = append(set, this.Icon.Icon(mod.IconLabel))
	}
	set = append(set, this.Component.Dv(mod.Label))

	return this.Component.Dcs("text-gray-500 flex flex-row space-x-1 items-center", set...)
}

func (this *KeyPoint) value(mod *KeyPointMod) goc.HTML {
	var set []goc.HTML

	if mod.IndicatorState != "" {
		set = append(set, this.indicatorState(mod))
	}
	set = append(set, this.Component.Dv(mod.Value))
	if mod.OverValue != "" {
		set = append(set, this.overValue(mod))
	}

	return this.Component.Dcs("text-black flex flex-row space-x-1 items-center", set...)
}

func (this *KeyPoint) indicatorState(mod *KeyPointMod) goc.HTML {
	color := ""
	switch mod.IndicatorState {
	case IndicatorStateSuccess:
		color = "bg-green-600"
	case IndicatorStateWarning:
		color = "bg-yellow-600"
	case IndicatorStateError:
		color = "bg-red-600"
	case IndicatorStateNoData:
		color = "bg-gray-600"
	default:
		panic("invalid indicator state")
	}
	return this.Component.Dc("h-3 w-3 h-3 w-3 rounded-full " + color)
}

func (this *KeyPoint) overValue(mod *KeyPointMod) goc.HTML {
	return this.Component.Dcs("flex flex-row space-x-1 text-gray-500", this.Component.Dv("/"), this.Component.Dv(mod.OverValue))
}
