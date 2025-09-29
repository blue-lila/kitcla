package steppers

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/goc"
)

type Stepper struct {
	Component *components.Component
}

const TextPositionBelow = "below"
const TextPositionAside = "aside"

type StepperMod struct {
	Steps        []*Step
	TextPosition string
	Width        string
}

type Step struct {
	Label  string
	Key    string
	Number string
}

func (this *StepperMod) AddNumberedStep(key string, label string, number string) {
	this.Steps = append(this.Steps, &Step{Label: label, Key: key, Number: number})
}

func (this *Stepper) Static(mod *StepperMod) goc.HTML {
	return goc.H("ul",
		goc.Attr{"class": "relative flex flex-row gap-x-2 " + mod.Width},
		this.items(mod),
	)
}

func (this *Stepper) Mod() *StepperMod {
	return &StepperMod{}
}

func (this *Stepper) H(mod *StepperMod) goc.HTML {
	return goc.H("ul",
		goc.Attr{"class": "relative flex flex-row gap-x-2"},
		this.items(mod),
	)
}

func (this *Stepper) listItem(mod *StepperMod, step *Step) goc.HTML {
	return goc.H("li",
		goc.Attr{"class": "shrink basis-0 flex-1 group"},
		this.listItemContent(mod, step),
		this.listItemSeparator(mod, step),
	)
}

func (this *Stepper) listItemContent(mod *StepperMod, step *Step) goc.HTML {
	return goc.H("div",
		goc.Attr{"class": "min-w-7 min-h-7 w-full inline-flex items-center text-xs align-middle"},
		this.listItemCircle(mod, step),
		this.listItemBar(mod),
	)
}

func (this *Stepper) listItemCircle(mod *StepperMod, step *Step) goc.HTML {
	value := ""
	if step.Number != "" {
		value = step.Number
	}

	return goc.H("span",
		goc.Attr{"class": "size-7 flex justify-center items-center shrink-0 bg-gray-100 font-medium text-gray-800 rounded-full"},
		value,
	)
}

func (this *Stepper) listItemBar(mod *StepperMod) goc.HTML {
	return goc.H("div",
		goc.Attr{"class": "ms-2 w-full h-px flex-1 bg-gray-200 group-last:hidden"},
	)
}

func (this *Stepper) listItemSeparator(mod *StepperMod, step *Step) goc.HTML {
	return goc.H("div",
		goc.Attr{"class": "mt-3"},
		this.listItemLabel(mod, step),
	)
}

func (this *Stepper) listItemLabel(mod *StepperMod, step *Step) goc.HTML {
	return goc.H("span",
		goc.Attr{"class": "block text-sm font-medium text-gray-800"},
		step.Label,
	)
}

func (this *Stepper) items(mod *StepperMod) []goc.HTML {
	var set []goc.HTML
	for _, step := range mod.Steps {
		set = append(set, this.listItem(mod, step))
	}
	return set
}
