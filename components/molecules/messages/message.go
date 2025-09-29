package messages

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/dat"
	"github.com/blue-lila/kitcla/goc"
)

type Message struct {
	Component *components.Component
}

type MessageMod struct {
	Message *dat.Message
}

func (this *Message) Message(message *dat.Message) goc.HTML {
	return this.H(&MessageMod{Message: message})
}

func (this *Message) H(mod *MessageMod) goc.HTML {
	css := ""
	if mod.Message.Kind == "success" {
		css = "text-green-700 bg-green-100 border-green-500"
	}
	if mod.Message.Kind == "warning" {
		css = "text-yellow-700 bg-yellow-100 border-yellow-500"
	}
	if mod.Message.Kind == "failure" {
		css = "text-red-700 bg-red-100 border-red-500"
	}

	return this.Component.Ccs("div", "flex flex-col font-bold py-4 px-5 border "+css, this.title(mod))
}
func (this *Message) title(mod *MessageMod) goc.HTML {
	return this.Component.Ccv("div", "text-bold", mod.Message.Title)
}
