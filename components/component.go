package components

import (
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/goc_debug"
)

type Component struct {
	Data struct {
		Debug bool
	}
}

// H Base for all components
func (this *Component) H(el string, attrs ...interface{}) goc.HTML {
	if this.Data.Debug {
		attrs = append(attrs, goc_debug.DebugInfo())
	}
	return goc.H(el, attrs...)
}

// C Component, just the html tag
func (this *Component) C(el string) goc.HTML {
	return this.H(el)
}

// Cc Component, with css
func (this *Component) Cc(el string, css string) goc.HTML {
	return this.H(el,
		goc.Attr{"class": css},
	)
}

// Ccv Component, with css and values
func (this *Component) Ccv(el string, css string, values ...string) goc.HTML {
	return this.H(el,
		goc.Attr{"class": css},
		values,
	)
}

// Cv Component, with values
func (this *Component) Cv(el string, values ...string) goc.HTML {
	return this.H(el,
		values,
	)
}

// Cs Component, with Subcomponents
func (this *Component) Cs(el string, components ...goc.HTML) goc.HTML {
	return this.H(el,
		components,
	)
}

// Cav Component, with Attributes and Values
func (this *Component) Cav(el string, attributes goc.Attr, values ...string) goc.HTML {
	return this.H(el,
		attributes,
		values,
	)
}

// Ca Component, with Attributes
func (this *Component) Ca(el string, attributes goc.Attr) goc.HTML {
	return this.H(el, attributes)
}

// Cas Component, with Attributes and Subcomponents
func (this *Component) Cas(el string, attributes goc.Attr, components ...goc.HTML) goc.HTML {
	return this.H(el,
		attributes,
		components,
	)
}

// Ccs Component, with Css and Subcomponents
func (this *Component) Ccs(el string, css string, components ...goc.HTML) goc.HTML {
	return this.H(el,
		goc.Attr{"class": css},
		components,
	)
}

// Dcs Div component, with Css and Subcomponents
func (this *Component) Dcs(css string, components ...goc.HTML) goc.HTML {
	return this.Ccs("div", css, components...)
}

// Ds Div component, with Subcomponents
func (this *Component) Ds(components ...goc.HTML) goc.HTML {
	return this.Cs("div", components...)
}

// Da Div component, with Attributes
func (this *Component) Da(attributes goc.Attr) goc.HTML {
	return this.Ca("div", attributes)
}

// Dc Div component, with Css
func (this *Component) Dc(css string) goc.HTML {
	return this.Cc("div", css)
}

// Dv Div component, with Css
func (this *Component) Dv(values ...string) goc.HTML {
	return this.Cv("div", values...)
}

// Dcv Div component, with Css and values
func (this *Component) Dcv(css string, value string) goc.HTML {
	return this.Ccv("div", css, value)
}

// Dav Div component, with Attribute and Value
func (this *Component) Dav(attributes goc.Attr, value string) goc.HTML {
	return this.Cav("div", attributes, value)
}

// Das Div component, with Attributes and Subcomponents
func (this *Component) Das(attributes goc.Attr, components ...goc.HTML) goc.HTML {
	return this.Cas("div", attributes, components...)
}

func (this *Component) Wrap(css string, components ...goc.HTML) goc.HTML {
	return this.W(css, components...)
}

func (this *Component) Ti(condition string, component goc.HTML) goc.HTML {
	return this.Cas("template", goc.Attr{"x-if": condition}, component)
}

func (this *Component) Tf(condition string, component goc.HTML) goc.HTML {
	return this.Cas("template", goc.Attr{"x-for": condition}, component)
}

func (this *Component) W(css string, components ...goc.HTML) goc.HTML {
	return this.Ccs("div", css, components...)
}

func (this *Component) Wa(attributes goc.Attr, components ...goc.HTML) goc.HTML {
	return this.Cas("div", attributes, components...)
}

func (this *Component) ExpHtml(html string) goc.HTML {
	return this.H("div", goc.UnsafeContent(html))
}

func (this *Component) OrNil(element goc.HTML, isNil bool) []goc.HTML {
	var collection []goc.HTML
	if isNil {
		return collection
	}
	collection = append(collection, element)

	return collection
}
