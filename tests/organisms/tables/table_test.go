package tables

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/components/organisms/tables"
	"github.com/blue-lila/kitcla/dat"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"github.com/blue-lila/kitcla/sup/placeholder"
	"testing"
)

func TestTableCheckboxAction(t *testing.T) {
	kit := kitcla.New()
	component := kit.Organisms.Tables.Table

	h := component.CheckboxAction("", "", nil)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/table_checkbox_action_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/table_checkbox_action_1.html")
	sup.AddPage("CheckboxAction", html)
}

func TestTableEmptyCellRow(t *testing.T) {
	kit := kitcla.New()
	component := kit.Organisms.Tables.Table

	h := component.EmptyCellRow()
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/table_empty_cell_row_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/table_empty_cell_row_1.html")
	sup.AddPage("EmptyCellRow", html)
}

func TestTableSubTable(t *testing.T) {
	kit := kitcla.New()
	component := kit.Organisms.Tables.Table

	mod := component.Mod()

	h := component.SubTable(mod)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/table_sub_table_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/table_sub_table_1.html")
	sup.AddPage("SubTable", html)
}

func TestTableTable(t *testing.T) {
	kit := kitcla.New()
	component := kit.Organisms.Tables.Table

	columns := placeholder.TableColumns()
	items := placeholder.MakeItems(10)

	pagination := &dat.Pagination{
		PerPage:     10,
		CurrentPage: 1,
		ItemsCount:  15,
		//ActiveFilters: bridge.ActiveFilters(fetcherPagination.Modulation),
	}
	mod := &tables.TableMod{
		Columns:    columns,
		Items:      items,
		Pagination: pagination,
		BaseUrl:    "/app",
		RowCell:    placeholder.RowCell,
		//AllowedFilters: bridge.AllowedFilters(fetcherPagination.Modulation),
	}

	h := component.Table(mod)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/table_table_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/table_table_1.html")
	sup.AddPage("Table", html)
}
