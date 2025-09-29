package trees

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/goc"
)

type Tree struct {
	Component *components.Component
}

type TreeMod struct {
	TreeId       string
	RootItem     *TreeItem
	TextRenderer TextRendererFunc
}

type TextRendererFunc func(mod *TreeMod, item *TreeItem) goc.HTML

type TreeItem struct {
	Id       string
	Value    string
	Text     string
	Leaf     bool
	Children []*TreeItem
}

func (this *Tree) Tree(treeId string, rootItem *TreeItem) goc.HTML {
	return this.H(&TreeMod{TreeId: treeId, RootItem: rootItem})
}

func (this *Tree) TreeWithTextRenderer(treeId string, rootItem *TreeItem, textRendererFunc TextRendererFunc) goc.HTML {
	return this.H(&TreeMod{TreeId: treeId, RootItem: rootItem, TextRenderer: textRendererFunc})
}

func (this *Tree) H(mod *TreeMod) goc.HTML {
	return this.Component.Ccs("ul", "list-disc ml-4", this.item(mod, mod.RootItem))
}

func (this *Tree) item(mod *TreeMod, item *TreeItem) goc.HTML {
	if mod.TextRenderer != nil {
		return this.Component.Cs("li", mod.TextRenderer(mod, item), this.children(mod, item))
	}
	return this.Component.Cs("li", this.text(mod, item), this.children(mod, item))
}

func (this *Tree) text(mod *TreeMod, item *TreeItem) goc.HTML {
	return this.Component.Cv("div", item.Text)
}

func (this *Tree) children(mod *TreeMod, item *TreeItem) goc.HTML {
	if len(item.Children) < 1 {
		return goc.HTML{}
	}
	var children []goc.HTML
	for _, child := range item.Children {
		children = append(children, this.item(mod, child))
	}
	return this.Component.Ccs("ul", "list-disc ml-4", children...)
}
