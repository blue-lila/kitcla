package duo_cards

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/components/atoms/card_bodies"
	"github.com/blue-lila/kitcla/components/organisms/tables"
	"github.com/blue-lila/kitcla/goc"

	"github.com/blue-lila/kitcla/utils"
	"strconv"
)

type PeekDuoCard struct {
	DuoCardBody *card_bodies.DuoCardBody
	Table       *tables.Table
	Component   *components.Component
}

type PeekDuoCardMod struct {
	Columns   []*tables.Column
	CellsRows []*tables.CellsRow
	ShowItems []goc.HTML
	TableMod  *tables.TableMod
}

func (this *PeekDuoCard) PeekDuoCard(columns []*tables.Column, cellsRows []*tables.CellsRow, showItems []goc.HTML) goc.HTML {
	mod := &PeekDuoCardMod{
		Columns:   columns,
		CellsRows: cellsRows,
		ShowItems: showItems,
	}
	return this.H(mod)
}

func (this *PeekDuoCard) PeekDuoCardViaTableMod(showItems []goc.HTML, tableMod *tables.TableMod) goc.HTML {
	mod := &PeekDuoCardMod{
		TableMod:  tableMod,
		ShowItems: showItems,
	}
	return this.H(mod)
}

func (this *PeekDuoCard) H(mod *PeekDuoCardMod) goc.HTML {
	cellsRows := mod.CellsRows
	if mod.TableMod != nil {
		cellsRows = mod.TableMod.CellsRows
	}
	for i, row := range cellsRows {
		n := strconv.Itoa(i)
		row.Attr = utils.MergeAttr(row.Attr, "@click", "tab = "+n)
	}
	var items []goc.HTML
	for i, item := range mod.ShowItems {
		n := strconv.Itoa(i)
		items = append(items, this.Component.Wa(goc.Attr{"x-show": "tab == " + n}, item))
	}
	mod.ShowItems = items

	tableMod := mod.TableMod
	if tableMod == nil {
		tableMod = this.Table.BodyOnlyViaCellsTableMod(mod.Columns, mod.CellsRows)
	}
	table := this.Table.H(tableMod)
	leftSide := this.DuoCardBody.Side("w-2/3", table)
	rightSide := this.DuoCardBody.Side("w-1/3", this.Component.Cs("div", mod.ShowItems...))
	mod2 := &card_bodies.DuoCardBodyMod{
		BodyAttr: goc.Attr{"x-data": "{tab: 0}"},
	}
	return this.DuoCardBody.DuoCardBody(leftSide, rightSide, mod2)
}
