package inputs

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/goc"
	"strings"
)

type AdvancedSelectInput struct {
	Component *components.Component
}

type AdvancedSelectInputMod struct {
	Name            string
	Value           string
	AutocompleteKey string
	Filters         []Filter
}

type Filter struct {
	Key   string
	Value string
}

func (this *AdvancedSelectInputMod) AddFilter(key string, value string) {
	this.Filters = append(this.Filters, Filter{Key: key, Value: value})
}

func (this *AdvancedSelectInput) Mod() *AdvancedSelectInputMod {
	return this.modDefaulting(nil)
}

func (this *AdvancedSelectInput) AdvancedSelectInput(name string, value string, key string) goc.HTML {
	mod := this.modDefaulting(nil)
	mod.Name = name
	mod.Value = value
	mod.AutocompleteKey = key
	return this.H(mod)
}

func (this *AdvancedSelectInput) AdvancedSelectInput2(name string, value string, key string, mod *AdvancedSelectInputMod) goc.HTML {
	mod = this.modDefaulting(mod)
	mod.Name = name
	mod.Value = value
	mod.AutocompleteKey = key
	return this.H(mod)
}

func (this *AdvancedSelectInput) modDefaulting(mod *AdvancedSelectInputMod) *AdvancedSelectInputMod {
	if mod == nil {
		return &AdvancedSelectInputMod{}
	}
	return mod
}

func (this *AdvancedSelectInput) hiddenInput(mod *AdvancedSelectInputMod) goc.HTML {
	return this.Component.Ca("input",
		goc.Attr{
			"name":    mod.Name,
			"x-model": "value",
			"type":    "hidden",
		},
	)
}

func (this *AdvancedSelectInput) textInput() goc.HTML {
	return this.Component.Ca("input",
		goc.Attr{
			"x-model":                     "filter",
			"x-transition:leave":          "transition ease-in duration-100",
			"x-transition:leave-start":    "opacity-100",
			"x-transition:leave-end":      "opacity-0",
			"@mousedown":                  "open()",
			"@keydown.enter.stop.prevent": "selectOption()",
			"@keydown.arrow-up.prevent":   "focusPrevOption()",
			"@keydown.arrow-down.prevent": "focusNextOption()",
			"class":                       "p-1 px-2 appearance-none outline-none w-full text-gray-800",
		},
	)
}
func (this *AdvancedSelectInput) arrowBottom() goc.HTML {
	return this.Component.Ca("polyline",
		goc.Attr{"x-show": "!isOpen()", "points": "18 15 12 20 6 15"},
	)
}
func (this *AdvancedSelectInput) arrowUp() goc.HTML {
	return this.Component.Ca("polyline",
		goc.Attr{"x-show": "isOpen()", "points": "18 15 12 9 6 15"},
	)
}
func (this *AdvancedSelectInput) arrow() goc.HTML {
	return this.Component.Cas("svg",
		goc.Attr{"xmlns": "http://www.w3.org/2000/svg", "width": "100%", "height": "100%", "fill": "none", "viewBox": "0 0 24 24", "stroke": "currentColor", "stroke-width": "2", "stroke-linecap": "round", "stroke-linejoin": "round"},
		this.arrowBottom(),
		this.arrowUp(),
	)
}
func (this *AdvancedSelectInput) function4() goc.HTML {
	return this.Component.Cas("div",
		goc.Attr{"@click": "toggle()", "class": "cursor-pointer w-6 h-6 text-gray-600 outline-none focus:outline-none"},
		this.arrow(),
	)
}
func (this *AdvancedSelectInput) function5() goc.HTML {
	return this.Component.Cas("div",
		goc.Attr{"class": "text-gray-300 w-8 py-1 pl-2 pr-1 border-l flex items-center border-gray-200"},
		this.function4(),
	)
}
func (this *AdvancedSelectInput) function6() goc.HTML {
	return this.Component.Cas("div",
		goc.Attr{"@click.away": "close()", "class": "my-2 p-1 bg-white flex border border-gray-200 rounded"},
		this.textInput(),
		this.function5(),
	)
}
func (this *AdvancedSelectInput) function7() goc.HTML {
	return this.Component.Cas("div",
		goc.Attr{"class": "w-full"},
		this.function6(),
	)
}

func (this *AdvancedSelectInput) function11() goc.HTML {
	return this.Component.Ca("span",
		goc.Attr{"x-text": "option.text"},
	)
}

func (this *AdvancedSelectInput) function13() goc.HTML {
	return this.Component.Cas("div",
		goc.Attr{"class": "mx-2 -mt-1"},
		this.function11(),
	)
}
func (this *AdvancedSelectInput) function14() goc.HTML {
	return this.Component.Cas("div",
		goc.Attr{"class": "w-full items-center flex"},
		this.function13(),
	)
}
func (this *AdvancedSelectInput) function15() goc.HTML {
	return this.Component.Cas("div",
		goc.Attr{"class": "flex w-full items-center p-2 pl-2 border-transparent border-l-2 relative hover:border-teal-100"},
		this.function14(),
	)
}
func (this *AdvancedSelectInput) function16() goc.HTML {
	return this.Component.Cas("div",
		goc.Attr{"@click": "onOptionClick(index)", ":class": "classOption(option.id, index)", ":aria-selected": "focusedOptionIndex === index"},
		this.function15(),
	)
}
func (this *AdvancedSelectInput) function17() goc.HTML {
	return this.Component.Cas("template",
		goc.Attr{"x-for": "(option, index) in filteredOptions()", ":key": "index"},
		this.function16(),
	)
}
func (this *AdvancedSelectInput) function18() goc.HTML {
	return this.Component.Cas("div",
		goc.Attr{"class": "flex flex-col w-full"},
		this.function17(),
	)
}
func (this *AdvancedSelectInput) function19() goc.HTML {
	return this.Component.Cas("div",
		goc.Attr{"x-show": "isOpen()", "style": "top: 100%;max-height: 300px;", "class": "absolute shadow bg-white  z-40 w-full lef-0 rounded overflow-y-auto"},
		this.function18(),
	)
}
func (this *AdvancedSelectInput) selectComponent(mod *AdvancedSelectInputMod) goc.HTML {
	return this.Component.Cas("div",
		goc.Attr{
			"x-data": "advancedInput('" + mod.Value + "')",
			"x-init": "fetchOptions('" + mod.AutocompleteKey + "', " + this.getFilters(mod) + ")",
			"class":  "w-full flex flex-col items-center relative",
		},
		this.hiddenInput(mod),
		this.function7(),
		this.function19(),
	)
}

func (this *AdvancedSelectInput) H(mod *AdvancedSelectInputMod) goc.HTML {
	return this.Component.Cas("div",
		goc.Attr{"class": "w-full md:w-1/2 flex flex-col items-center"},
		this.selectComponent(mod),
	)
}

func (this *AdvancedSelectInput) getFilters(mod *AdvancedSelectInputMod) string {
	var ss []string
	for _, filter := range mod.Filters {
		ss = append(ss, "{key:\""+filter.Key+"\",value:\""+filter.Value+"\"}")
	}
	return "[" + strings.Join(ss, ",") + "]"
}
