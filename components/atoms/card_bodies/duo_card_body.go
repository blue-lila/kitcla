package card_bodies

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/utils"
)

type DuoCardBody struct {
	Component *components.Component
}

type DuoCardBodyMod struct {
	LeftSide  *Side
	RightSide *Side
	BodyAttr  goc.Attr
}

type Side struct {
	Ratio   string
	Content goc.HTML
}

func (this *DuoCardBody) Side(ratio string, content goc.HTML) *Side {
	return &Side{
		Ratio:   ratio,
		Content: content,
	}
}

func (this *DuoCardBody) DuoCardBody(leftSide *Side, rightSide *Side, mod *DuoCardBodyMod) goc.HTML {
	mod = this.modDefaulting(mod)
	mod.RightSide = rightSide
	mod.LeftSide = leftSide
	return this.H(mod)
}

func (this *DuoCardBody) modDefaulting(mod *DuoCardBodyMod) *DuoCardBodyMod {
	if mod != nil {
		return mod
	}
	return &DuoCardBodyMod{}
}

func (this *DuoCardBody) H(mod *DuoCardBodyMod) goc.HTML {
	if mod.LeftSide == nil || mod.RightSide == nil {
		return this.Component.Cv("div", "Invalid params")
	}

	return this.Component.Dcs("flex flex-col bg-white border shadow-sm rounded-xl",
		this.body(mod),
	)
}

func (this *DuoCardBody) body(mod *DuoCardBodyMod) goc.HTML {
	attr := mod.BodyAttr
	attr = utils.MergeCssAttr(attr, "flex flex-row divide-x")
	return this.Component.Das(attr, this.left(mod), this.right(mod))
}

func (this *DuoCardBody) right(mod *DuoCardBodyMod) goc.HTML {
	return this.Component.Dcs(mod.RightSide.Ratio, mod.RightSide.Content)
}

func (this *DuoCardBody) left(mod *DuoCardBodyMod) goc.HTML {
	return this.Component.Dcs(mod.LeftSide.Ratio, mod.LeftSide.Content)
}
