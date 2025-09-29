package inputs

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/components/atoms/resources"
	"github.com/blue-lila/kitcla/goc"
)

type RichTextInput struct {
	Resource  *resources.Resource
	Component *components.Component
}

type RichTextInputMod struct {
	Name  string
	Value string
}

func (this *RichTextInput) RichTextInput(name string, value string) goc.HTML {
	return this.H(&RichTextInputMod{Name: name, Value: value})
}

func (this *RichTextInput) Deps() goc.HTML {
	js := "https://cdnjs.cloudflare.com/ajax/libs/trix/1.3.1/trix.js"
	css := "https://cdnjs.cloudflare.com/ajax/libs/trix/1.3.1/trix.css"
	return this.Resource.ResourceJsCss(js, css)
}

func (this *RichTextInput) H(mod *RichTextInputMod) goc.HTML {
	return this.Component.Wrap("", this.input(mod), this.trixEditor(mod), this.removeAttachOptionJs())
}

func (this *RichTextInput) removeAttachOptionJs() goc.HTML {
	return this.Component.H("script",
		goc.Attr{"type": "text/javascript"},
		goc.UnsafeContent("document.querySelector(\".trix-button-group--file-tools\").remove();"),
	)
}

func (this *RichTextInput) trixEditor(mod *RichTextInputMod) goc.HTML {
	return this.Component.Cav("trix-editor",
		goc.Attr{"input": "trix-field-" + mod.Name, "class": "focus:outline-none focus:border focus:border-blue-500"},
	)
}

func (this *RichTextInput) input(mod *RichTextInputMod) goc.HTML {
	return this.Component.Cav("input",
		goc.Attr{"id": "trix-field-" + mod.Name, "type": "hidden", "name": mod.Name, "value": mod.Value},
	)
}
