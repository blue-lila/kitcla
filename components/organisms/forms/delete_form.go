package forms

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/components/atoms/buttons"
	"github.com/blue-lila/kitcla/components/atoms/links"
	"github.com/blue-lila/kitcla/goc"
)

type DeleteForm struct {
	Button    *buttons.Button
	Link      *links.Link
	Component *components.Component
}

type DeleteFormMod struct {
	ActionUrl string
}

func (this *DeleteForm) DeleteForm(actionUrl string) goc.HTML {
	return this.H(&DeleteFormMod{ActionUrl: actionUrl})
}

func (this *DeleteForm) H(mod *DeleteFormMod) goc.HTML {
	deleteButton := this.Link.SubmitLink("Delete")
	return this.Component.Cas("form", goc.Attr{"action": mod.ActionUrl, "method": "post"}, deleteButton)
}
