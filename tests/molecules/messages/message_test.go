package messages

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/dat"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"testing"
)

func TestMessageMessage(t *testing.T) {
	kit := kitcla.New()
	component := kit.Molecules.Messages.Message

	h := component.Message(&dat.Message{
		Title: "Hello",
		Body:  "Test",
		Kind:  "warning",
	})
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/message_message_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/message_message_1.html")
}
