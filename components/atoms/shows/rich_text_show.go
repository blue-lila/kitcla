package shows

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/goc"
)

type RichTextShow struct {
	Component *components.Component
}

type RichTextShowMod struct {
	Name  string
	Value string
}

func (this *RichTextShow) RichTextShow(name string, value string) goc.HTML {
	return this.H(&RichTextShowMod{Name: name, Value: value})
}

// H mod.value must be sanitized
// You can use bluemonday for that
//
//	Import via: "github.com/microcosm-cc/bluemonday"
//	Use it:
//		p := bluemonday.UGCPolicy()
//		s := p.Sanitize(s)
func (this *RichTextShow) H(mod *RichTextShowMod) goc.HTML {
	return this.Component.H("div", goc.UnsafeContent(mod.Value))
}
