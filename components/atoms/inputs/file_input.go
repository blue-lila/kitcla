package inputs

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/components/atoms/fields"
	"github.com/blue-lila/kitcla/goc"
)

type FileInput struct {
	fields.Field
	Component *components.Component
}

type FileInputMod struct {
}

func (this *FileInput) H(mod *FileInputMod) goc.HTML {
	return goc.HTML{}
}

func (this *FileInput) FileInput(name string) goc.HTML {
	return this.Component.Ca("input", goc.Attr{"type": "file", "name": "file[" + name + "]"})
}
