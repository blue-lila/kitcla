package utils

import "github.com/blue-lila/kitcla/goc"

func MergeAttr(attr goc.Attr, key string, value string) goc.Attr {
	if attr == nil {
		attr = goc.Attr{}
	}
	attr[key] = value
	return attr
}

func MergeCssAttr(attr goc.Attr, value string) goc.Attr {
	if attr == nil {
		attr = goc.Attr{}
	}
	attr["class"] = value
	return attr
}
