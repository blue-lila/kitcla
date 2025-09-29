package grids

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/goc"
)

type KeyValueGrid struct {
	Component *components.Component
}

type KeyValueItem struct {
	Key   string
	Value goc.HTML
}

type KeyValueGridMod struct {
	Items []*KeyValueItem
}

func (this *KeyValueGrid) NewItems() []*KeyValueItem {
	var items []*KeyValueItem
	return items
}

func (this *KeyValueGrid) AppendItem(items []*KeyValueItem, key string, value goc.HTML) []*KeyValueItem {
	return append(items, &KeyValueItem{Key: key, Value: value})
}

func (this *KeyValueGrid) KeyValueGrid(items []*KeyValueItem) goc.HTML {
	return this.H(&KeyValueGridMod{Items: items})
}

func (this *KeyValueGrid) H(mod *KeyValueGridMod) goc.HTML {
	return this.Component.Ccs("div", "w-full overflow-hidden rounded border-scale-3",
		this.grid(mod))
}

func (this *KeyValueGrid) grid(mod *KeyValueGridMod) goc.HTML {
	var cells []goc.HTML
	for _, item := range mod.Items {
		cells = append(cells, this.gridCell(mod, item))
	}

	return this.Component.Ccs("div",
		"grid grid-cols-6 lg:grid-cols-6 sm:grid-cols-2 gap-5",
		cells...,
	)
}

func (this *KeyValueGrid) cellKey(mod *KeyValueGridMod, item *KeyValueItem) goc.HTML {
	return this.Component.Ccv("div", "text-xs uppercase text-gray-500", item.Key)
}

func (this *KeyValueGrid) cellValue(mod *KeyValueGridMod, item *KeyValueItem) goc.HTML {
	return this.Component.Ccs("div", "", item.Value)
}

func (this *KeyValueGrid) gridCell(mod *KeyValueGridMod, item *KeyValueItem) goc.HTML {
	return this.Component.Ccs("div", "flex flex-col space-y-2",
		this.cellKey(mod, item),
		this.cellValue(mod, item),
	)
}
