package goc

import (
	"fmt"
	"html"
	"sort"
	"strings"
)

var selfClosingTags = map[string]int{
	"area":  1,
	"br":    1,
	"hr":    1,
	"image": 1,
	"input": 1,
	"img":   1,
	"link":  1,
	"meta":  1,
}

// Attr is a HTML element attribute
// <a href="#"> => Attr{"href": "#"}
type Attr map[string]string

type HTML struct {
	El    string
	Attrs []interface{}
}

type Debug struct {
	Id   string
	Path string
}

// dangerous contents type
type dangerousContents func() (string, bool)

// UnsafeContent allows injection of JS or HTML from functions
func UnsafeContent(str string) dangerousContents {
	return func() (string, bool) {
		return str, true
	}
}

// H is the base HTML func
func H(el string, attrs ...interface{}) HTML {
	var as []interface{}
	for _, v := range attrs {
		as = append(as, v)
	}
	return HTML{El: el, Attrs: as}
}

func RenderRoot(root HTML) string {
	return Render(root, root.El, root.Attrs...)
}

func Render(h HTML, el string, attrs ...interface{}) string {
	if el == "" {
		return ""
	}

	var contents []string
	attributes := ""
	for _, v := range attrs {
		switch v := v.(type) {
		case string:
			contents = append(contents, escape(v))
		case Attr:
			attributes = attributes + getAttributes(v)
		case []string:
			children := strings.Join(v, "")
			contents = append(contents, escape(children))
		case []HTML:
			children := subItems(v)
			contents = append(contents, children)
		case HTML:
			contents = append(contents, Render(v, v.El, v.Attrs...))
		case dangerousContents:
			t, _ := v()
			contents = append(contents, t)
		case Debug:
			attributes = attributes + getAttributes(Attr{"data-gxi": v.Id})
		case func() string:
			contents = append(contents, escape(v()))
		default:
			contents = append(contents, escape(fmt.Sprintf("%v", v)))
		}
	}

	elc := escape(el)
	if _, ok := selfClosingTags[elc]; ok {
		return "<" + elc + attributes + " />"
	}
	return "<" + elc + attributes + ">" + strings.Join(contents, "") + "</" + elc + ">"
}

func escape(str string) string {
	return html.EscapeString(str)
}

func subItems(attrs []HTML) string {
	var res []string
	for _, v := range attrs {
		res = append(res, Render(v, v.El, v.Attrs...))
	}
	return strings.Join(res, "")
}

func getAttributes(attributes Attr) string {
	var res []string
	var keys []string
	for k := range attributes {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		v := attributes[k]
		pair := fmt.Sprintf("%v='%v'", escape(k), escape(v))
		res = append(res, pair)
	}
	prefix := ""
	if len(res) > 0 {
		prefix = " "
	}
	return prefix + strings.Join(res, " ")
}
