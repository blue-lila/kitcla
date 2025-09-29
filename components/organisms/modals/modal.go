package modals

import (
	"github.com/blue-lila/kitcla/aria"
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/components/atoms/buttons"
	"github.com/blue-lila/kitcla/goc"
)

type Modal struct {
	Button    *buttons.Button
	Component *components.Component
}

type ModalMod struct {
}

func (this *Modal) SimpleModal() goc.HTML {
	return this.H(&ModalMod{})
}

func (this *Modal) H(mod *ModalMod) goc.HTML {
	return this.Component.Dcs("relative z-10", this.backgroundOverlay(mod), this.screenCapture(mod))
}

func (this *Modal) backgroundOverlay(mod *ModalMod) goc.HTML {
	return this.Component.Dc("fixed inset-0 bg-gray-500 bg-opacity-75")
}

func (this *Modal) screenCapture(mod *ModalMod) goc.HTML {
	return this.Component.Dcs("fixed inset-0 z-10 w-screen overflow-y-auto", this.container(mod))
}

func (this *Modal) container(mod *ModalMod) goc.HTML {
	return this.Component.Dcs("flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0", this.modal(mod))
}

func (this *Modal) modal(mod *ModalMod) goc.HTML {
	css := "relative overflow-hidden rounded-lg bg-scale-0 px-4 pb-4 pt-5 text-left shadow-xl sm:my-8 sm:w-full sm:max-w-sm sm:p-6"
	return this.Component.Cas("div", goc.Attr{"class": css, "role": aria.AriaRoleDialog, aria.AriaModal: aria.AriaModalTrue},
		this.modalBody(mod), this.modalActions(mod))
}

func (this *Modal) modalBody(mod *ModalMod) goc.HTML {
	return this.Component.Dcs("mt-3 text-center sm:mt-5 space-y-2", this.modalTitle(mod), this.modalText(mod))
}

func (this *Modal) modalActions(mod *ModalMod) goc.HTML {
	return this.Component.Dcs("mt-5 sm:mt-6 flex justify-center", this.Button.PrimarySubmit("Okay", nil))
}

func (this *Modal) modalTitle(mod *ModalMod) goc.HTML {
	return this.Component.Dcv("text-base font-semibold leading-6 text-scale-10", "Hello Title 2")
}

func (this *Modal) modalText(mod *ModalMod) goc.HTML {
	return this.Component.Dcv("text-sm text-scale-7", "Hello Text")
}
