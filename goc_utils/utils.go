package goc_utils

import "github.com/blue-lila/kitcla/goc"

// C Component
func C(el string) goc.HTML {
	return goc.H(el)
}

// Cc Component, with css
func Cc(el string, css string) goc.HTML {
	return goc.H(el,
		goc.Attr{"class": css},
	)
}

// Ccv Component, with css and values
func Ccv(el string, css string, values ...string) goc.HTML {
	return goc.H(el,
		goc.Attr{"class": css},
		values,
	)
}

// Cv Component, with values
func Cv(el string, values ...string) goc.HTML {
	return goc.H(el,
		values,
	)
}

// Cs Component, with Subcomponents
func Cs(el string, components ...goc.HTML) goc.HTML {
	return goc.H(el,
		components,
	)
}

// Cav Component, with Attributes and Values
func Cav(el string, attributes goc.Attr, values ...string) goc.HTML {
	return goc.H(el,
		attributes,
		values,
	)
}

// Ca Component, with Attributes
func Ca(el string, attributes goc.Attr) goc.HTML {
	return goc.H(el, attributes)
}

// Cas Component, with Attributes and Subcomponents
func Cas(el string, attributes goc.Attr, components ...goc.HTML) goc.HTML {
	return goc.H(el,
		attributes,
		components,
	)
}

// Ccs Component, with Css and Subcomponents
func Ccs(el string, css string, components ...goc.HTML) goc.HTML {
	return goc.H(el,
		goc.Attr{"class": css},
		components,
	)
}

// Dcs Div component, with Css and Subcomponents
func Dcs(css string, components ...goc.HTML) goc.HTML {
	return Ccs("div", css, components...)
}

// Ds Div component, with Subcomponents
func Ds(components ...goc.HTML) goc.HTML {
	return Cs("div", components...)
}

// Da Div component, with Attributes
func Da(attributes goc.Attr) goc.HTML {
	return Ca("div", attributes)
}

// Dc Div component, with Css
func Dc(css string) goc.HTML {
	return Cc("div", css)
}

// Dv Div component, with Css
func Dv(values ...string) goc.HTML {
	return Cv("div", values...)
}

// Dcv Div component, with Css and values
func Dcv(css string, value string) goc.HTML {
	return Ccv("div", css, value)
}

// Dav Div component, with Attribute and Value
func Dav(attributes goc.Attr, value string) goc.HTML {
	return Cav("div", attributes, value)
}

// Das Div component, with Attributes and Subcomponents
func Das(attributes goc.Attr, components ...goc.HTML) goc.HTML {
	return Cas("div", attributes, components...)
}

func Wrap(css string, components ...goc.HTML) goc.HTML {
	return W(css, components...)
}

func Ti(condition string, component goc.HTML) goc.HTML {
	return Cas("template", goc.Attr{"x-if": condition}, component)
}

func Tf(condition string, component goc.HTML) goc.HTML {
	return Cas("template", goc.Attr{"x-for": condition}, component)
}

func W(css string, components ...goc.HTML) goc.HTML {
	return Ccs("div", css, components...)
}

func Wa(attributes goc.Attr, components ...goc.HTML) goc.HTML {
	return Cas("div", attributes, components...)
}

func ExpHtml(html string) goc.HTML {
	return goc.H("div", goc.UnsafeContent(html))
}

func OrNil(element goc.HTML, isNil bool) []goc.HTML {
	var collection []goc.HTML
	if isNil {
		return collection
	}
	collection = append(collection, element)

	return collection
}

type KVS struct {
	Key   string
	Value string
}

func KeysValuesString(keys []string, values []string) []*KVS {
	var kvs []*KVS
	for i, key := range keys {
		label := &KVS{
			Key:   key,
			Value: values[i],
		}
		kvs = append(kvs, label)
	}
	return kvs
}
