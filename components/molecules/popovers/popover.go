package popovers

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/components/atoms/buttons"
	"github.com/blue-lila/kitcla/goc"
)

type Popover struct {
	ButtonAlp *buttons.ButtonAlp
	Component *components.Component
}

type PopoverMod struct {
	Button  goc.HTML
	Content goc.HTML
	Width   string
}

func (this *Popover) Popover(button goc.HTML, content goc.HTML) goc.HTML {
	mod := &PopoverMod{Content: content}
	return this.H(mod)
}

func (this *Popover) TextPopover(text string, content goc.HTML) goc.HTML {
	mod := &PopoverMod{Content: content, Button: this.textButton(text)}
	return this.H(mod)
}

func (this *Popover) IconPopover(icon string, content goc.HTML) goc.HTML {
	mod := &PopoverMod{Content: content, Button: this.iconButton(icon)}
	return this.H(mod)
}

func (this *Popover) IconPopoverWithFixedWidth(icon string, content goc.HTML, width string) goc.HTML {
	mod := &PopoverMod{Content: content, Button: this.iconButton(icon), Width: width}
	return this.H(mod)
}

func (this *Popover) GhostlyIconPopover(icon string, content goc.HTML) goc.HTML {
	mod := &PopoverMod{Content: content, Button: this.ghostlyIconButton(icon)}
	return this.H(mod)
}

func (this *Popover) H(mod *PopoverMod) goc.HTML {
	return this.Component.Das(goc.Attr{"x-data": "{ open: false }", "class": "relative"},
		mod.Button,
		this.outerWrapper(mod),
	)
}

func (this *Popover) textButton(text string) goc.HTML {
	return this.ButtonAlp.SecondaryLink(text, "open = !open", nil)
}

func (this *Popover) iconButton(icon string) goc.HTML {
	return this.ButtonAlp.SecondaryIconLink(icon, "open = !open", nil)
}

func (this *Popover) ghostlyIconButton(icon string) goc.HTML {
	return this.ButtonAlp.QuaternaryIconLink(icon, "open = !open", nil)
}

func (this *Popover) outerWrapper(mod *PopoverMod) goc.HTML {
	return this.Component.Das(
		goc.Attr{
			"x-show":                   "open",
			"@click.away":              "open = false",
			"x-transition:enter":       "transition ease-out duration-200",
			"x-transition:enter-start": "opacity-0 transform scale-95",
			"x-transition:enter-end":   "opacity-100 transform scale-100",
			"x-transition:leave":       "transition ease-in duration-150",
			"x-transition:leave-start": "opacity-100 transform scale-100",
			"x-transition:leave-end":   "opacity-0 transform scale-95",
			"class":                    "absolute top-0 right-0 mt-8 bg-white border border-gray-200 rounded-lg shadow-lg z-10",
		},
		this.innerWrapper(mod),
	)
}

func (this *Popover) innerWrapper(mod *PopoverMod) goc.HTML {
	a := goc.Attr{"class": "p-4"}
	if mod.Width != "" {
		a["style"] = "width: " + mod.Width + "px;"
	}
	return this.Component.Das(a, mod.Content)
}
