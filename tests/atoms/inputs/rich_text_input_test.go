package inputs

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"testing"
)

func TestRichTextInputDeps(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Inputs.RichTextInput

	h := component.Deps()
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/rich_text_input_deps_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/rich_text_input_deps_1.html")
}

func TestRichTextInputRichTextInput(t *testing.T) {
	kit := kitcla.New()
	component := kit.Atoms.Inputs.RichTextInput

	h := component.RichTextInput("", "")
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/rich_text_input_rich_text_input_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/rich_text_input_rich_text_input_1.html")
}
