package inputs

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/components/atoms/icons"
	"github.com/blue-lila/kitcla/goc"
)

type IntegerInputAlp struct {
	Icon      *icons.Icon
	Component *components.Component
}

type IntegerInputAlpMod struct {
	XModel       string
	OnMinusClick string
	OnPlusClick  string
}

func (this *IntegerInputAlp) IntegerInput(xModel string) goc.HTML {
	return this.H(&IntegerInputAlpMod{XModel: xModel})
}

func (this *IntegerInputAlp) H(mod *IntegerInputAlpMod) goc.HTML {
	return this.Component.Cav("input",
		goc.Attr{"type": "number", "x-model": mod.XModel, "class": "border border-scale-3 px-3 py-2 rounded-md"},
	)
}

func (this *IntegerInputAlp) Mini(xModel string) goc.HTML {
	return this.Component.Dcs("py-2 px-3 inline-block bg-white border border-gray-200 rounded-lg",
		this.mini1(&IntegerInputAlpMod{XModel: xModel}),
	)
}

func (this *IntegerInputAlp) mini1(mod *IntegerInputAlpMod) goc.HTML {
	return this.Component.Dcs("flex items-center gap-x-1.5",
		this.miniMinusButton(mod),
		this.miniInput(mod),
		this.miniPlusButton(mod),
	)
}

func (this *IntegerInputAlp) miniButtonCss() string {
	return "size-6 inline-flex justify-center items-center gap-x-2 text-sm font-medium rounded-md border " +
		"border-gray-200 bg-white text-gray-800 shadow-sm hover:bg-gray-50 focus:outline-none " +
		"focus:bg-gray-50 disabled:opacity-50 disabled:pointer-events-none"
}

func (this *IntegerInputAlp) miniMinusButton(mod *IntegerInputAlpMod) goc.HTML {
	return this.Component.Das(goc.Attr{
		"class": this.miniButtonCss(),
	}, this.Icon.IconWithSize(icons.IconFasMinus, "3"))
}

func (this *IntegerInputAlp) miniInput(mod *IntegerInputAlpMod) goc.HTML {
	css := "p-0 w-6 bg-transparent border-0 text-gray-800 text-center" +
		" focus:ring-0 [&::-webkit-inner-spin-button]:appearance-none [&::-webkit-outer-spin-button]:appearance-none"

	return this.Component.Cas("input",
		goc.Attr{
			"class":          css,
			"style":          "-moz-appearance: textfield;",
			"type":           "number",
			"x-model.number": mod.XModel,
		},
	)
}

func (this *IntegerInputAlp) miniPlusButton(mod *IntegerInputAlpMod) goc.HTML {
	return this.Component.Dcs(this.miniButtonCss(), this.Icon.IconWithSize(icons.IconFasPlus, "3"))
}
