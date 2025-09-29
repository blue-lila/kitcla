package tables

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/goc"
)

type KeyValueTable struct {
	Component *components.Component
}

type KeyValueItem struct {
	Key   string
	Value goc.HTML
}

type KeyValueTableMod struct {
	Items      []*KeyValueItem
	KeyWidth   string
	ValueWidth string
}

func (this *KeyValueTable) NewItems() []*KeyValueItem {
	var items []*KeyValueItem
	return items
}

func (this *KeyValueTable) AppendItem(items []*KeyValueItem, key string, value goc.HTML) []*KeyValueItem {
	return append(items, &KeyValueItem{Key: key, Value: value})
}

func (this *KeyValueTable) KeyValueTable(items []*KeyValueItem, keyWidth string, valueWidth string) goc.HTML {
	return this.H(&KeyValueTableMod{Items: items, KeyWidth: keyWidth, ValueWidth: valueWidth})
}

func (this *KeyValueTable) H(mod *KeyValueTableMod) goc.HTML {
	return this.Component.Ccs("div", "w-full overflow-hidden rounded border-scale-3", this.table(mod))
}

func (this *KeyValueTable) table(mod *KeyValueTableMod) goc.HTML {
	return this.Component.Ccs("div",
		"bg-scale-0 min-w-full",
		this.body(mod),
	)
}

func (this *KeyValueTable) cellKey(mod *KeyValueTableMod, item *KeyValueItem) goc.HTML {
	return this.Component.Ccv("div", mod.KeyWidth+" text-scale-6", item.Key)
}

func (this *KeyValueTable) cellValue(mod *KeyValueTableMod, item *KeyValueItem) goc.HTML {
	return this.Component.Ccs("div", mod.ValueWidth, item.Value)
}

func (this *KeyValueTable) row(mod *KeyValueTableMod, item *KeyValueItem) goc.HTML {
	return this.Component.Ccs("div", "flex py-4", this.cellKey(mod, item), this.cellValue(mod, item))
}

func (this *KeyValueTable) body(mod *KeyValueTableMod) goc.HTML {
	var rows []goc.HTML
	for _, item := range mod.Items {
		rows = append(rows, this.row(mod, item))
	}
	return this.Component.Ccs("div", "divide-y divide-grey-100 divide-solid", rows...)
}
