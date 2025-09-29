package tables

import (
	"encoding/hex"

	"github.com/blue-lila/kitcla/aria"
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/components/atoms/buttons"
	"github.com/blue-lila/kitcla/components/atoms/dropdowns"
	"github.com/blue-lila/kitcla/components/atoms/headers"
	"github.com/blue-lila/kitcla/components/atoms/icons"
	"github.com/blue-lila/kitcla/components/atoms/inputs"
	"github.com/blue-lila/kitcla/components/atoms/paginations"
	"github.com/blue-lila/kitcla/dat"
	"github.com/blue-lila/kitcla/goc"

	"math/rand"
	"strconv"
	"strings"
)

type Table struct {
	Button         *buttons.Button
	TextInput      *inputs.TextInput
	BooleanInput   *inputs.BooleanInput
	SelectInput    *inputs.SelectInput
	SelectInputAlp *inputs.SelectInputAlp
	IntegerInput   *inputs.IntegerInput
	Dropdown       *dropdowns.Dropdown
	Header         *headers.Header
	Component      *components.Component
	Pagination     *paginations.Pagination
}

type Item struct {
	Item dat.Entity
}

type CellsRow struct {
	Cells []goc.HTML
	Attr  goc.Attr
}

type TableMod struct {
	TableId           string
	Columns           []*Column
	Rows              []goc.HTML
	RowsViaCells      bool
	CellsRows         []*CellsRow
	Items             []dat.Entity
	Pagination        *dat.Pagination
	BaseUrl           string
	AllowedFilters    []*dat.Filter
	CheckboxActions   []goc.HTML
	HideTableHeader   bool
	HideActionsHeader bool
	HideFiltering     bool
	HidePagination    bool
	RowCell           func(item dat.Entity, column *Column) goc.HTML
	Title             string
	FullTableUrl      string
	NewEntityUrl      string
}

type CheckboxActionMod struct {
	Confirmation bool
}

const ColumnKindData = "data"
const ColumnKindId = "id"
const ColumnKindAction = "action"

type Column struct {
	Key    string
	Label  string
	Kind   string
	Width  string
	Hidden bool
}

func (this *Table) Mod() *TableMod {
	return &TableMod{}
}

func (this *Table) Columns(setToMerge []*Column) []*Column {
	var columns []*Column

	columns = append(columns, &Column{
		Key:  "_id",
		Kind: ColumnKindId,
	})
	for _, column := range setToMerge {
		columns = append(columns, column)
	}
	columns = append(columns, &Column{
		Key:   "_actions",
		Kind:  ColumnKindAction,
		Label: "Actions",
	})

	return columns
}

func (this *Table) CheckboxAction(label string, link string, mod *CheckboxActionMod) goc.HTML {
	modS := this.convertModToJsObject(mod)
	return this.Component.Cav("div", goc.Attr{"x-on:click": "submitCheckedIds($el, " + modS + ")", "data-href": link, "class": "cursor-pointer"}, label)
}

func (this *Table) convertModToJsObject(mod *CheckboxActionMod) string {
	var m []string
	if mod == nil {
		return "{}"
	}

	if mod.Confirmation {
		m = append(m, "\"confirmation\": true")
	}
	if len(m) < 1 {
		return "{}"
	}
	return "{" + strings.Join(m, ",") + "}"
}

func (this *Table) randomTableId() string {
	b := make([]byte, 4)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func (this *Table) SetWidthsWithBestEffort(columns []*Column) []*Column {
	count := len(columns)
	if count > 1 && count <= 6 {
		denominator := strconv.Itoa(count)
		// for the css pruner: w-1/2 w-1/3 w-1/4 w-1/5 w-1/6
		return this.assignWidth("1/"+denominator, columns)
	}
	return columns
}

func (this *Table) assignWidth(width string, columns []*Column) []*Column {
	for _, h := range columns {
		h.Width = "w-" + width
	}
	return columns
}

func (this *Table) TdCss() string {
	return "p-3 text-left"
}

func (this *Table) EmptyCellRow() goc.HTML {
	return this.Component.C("td")
}

func (this *Table) BodyOnlyViaCellsTableMod(columns []*Column, cellsRows []*CellsRow) *TableMod {
	return &TableMod{
		TableId:           "",
		Columns:           columns,
		RowsViaCells:      true,
		CellsRows:         cellsRows,
		Items:             nil,
		Pagination:        nil,
		BaseUrl:           "",
		AllowedFilters:    nil,
		CheckboxActions:   nil,
		HideActionsHeader: true,
		HideTableHeader:   true,
		HideFiltering:     true,
		HidePagination:    true,
	}
}

func (this *Table) H(mod *TableMod) goc.HTML {
	if mod.TableId == "" {
		mod.TableId = this.randomTableId()
	}

	var footer goc.HTML
	if mod.Pagination != nil {
		footer = this.footer(mod)
	}

	return this.Component.W("-m-1.5 overflow-x-auto min-w-full", this.Component.W("p-1.5 min-w-full inline-block align-middle", this.Component.Cas("div", goc.Attr{
		"x-data": "table('" + mod.TableId + "')",
		"class":  "w-full bg-white border border-gray-200 rounded-lg shadow-sm"},
		this.headerActionsRow(mod),
		this.table(mod),
		footer,
	)))
}

func (this *Table) Table(mod *TableMod) goc.HTML {
	return this.H(mod)
}

func (this *Table) SubTable(mod *TableMod) goc.HTML {
	if mod == nil {
		return goc.HTML{}
	}
	mod.HideFiltering = true
	return this.H(mod)
}

func (this *Table) table(mod *TableMod) goc.HTML {
	var subs []goc.HTML
	if mod.HideTableHeader != true {
		subs = append(subs, this.tableHeader(mod))
	}
	subs = append(subs, this.body(mod))

	return this.Component.Cas("table",
		goc.Attr{
			"class": "bg-scale-0 min-w-full divide-y table-fixed",
			"role":  aria.AriaRoleTable,
		},
		subs...,
	)
}

func (this *Table) newFilterButton(mod *TableMod) goc.HTML {
	return this.Component.Cas("span",
		goc.Attr{"x-on:click": "toggleFiltersChoice()", "class": "flex items-center"},
		this.Button.SecondaryIconLink(icons.IconFilter, "#", nil),
	)
}

func (this *Table) filterOperation(mod *TableMod, filter *dat.Filter) goc.HTML {
	selectMod := this.SelectInput.Mod()
	for _, operation := range filter.Operations {
		selectMod.AddOption(operation.Name, operation.Text)
	}
	selectMod.SelectAttr = goc.Attr{"x-on:change": "chooseOperation", "id": "_new_filter_" + filter.Name}
	showCondition := "filterKey === '" + filter.Name + "'"
	return this.Component.Cas("span", goc.Attr{"x-show": showCondition}, this.SelectInput.H(selectMod))
}

func (this *Table) filterOperations(mod *TableMod) goc.HTML {
	var operations []goc.HTML
	for _, filter := range mod.AllowedFilters {
		operation := this.filterOperation(mod, filter)
		operations = append(operations, operation)
	}
	return this.Component.Cs("span", operations...)
}

func (this *Table) filterWidget(mod *TableMod, filter *dat.Filter, operation *dat.FilterOperation) goc.HTML {
	condition := "operationKey === '" + operation.Name + "'"
	condition += " && filterKey === '" + filter.Name + "'"

	attr := goc.Attr{
		"x-show": condition,
	}
	if operation.Widget == "integer" {
		input := this.Component.Cav("input",
			goc.Attr{"type": "text", "x-model": "value", "class": "border border-scale-3 px-3 py-2 rounded-md"},
		)
		return this.Component.Cas("span", attr, input)
	}
	if operation.Widget == "string" {
		input := this.Component.Cav("input",
			goc.Attr{"type": "text", "x-model": "value", "class": "border border-scale-3 px-3 py-2 rounded-md"},
		)
		return this.Component.Cas("span", attr, input)
	}
	if operation.Widget == "boolean" {
		selectMod := this.SelectInputAlp.Mod()
		selectMod.AddOption("", "")
		selectMod.AddOption("true", "true")
		selectMod.AddOption("false", "false")
		input := this.SelectInputAlp.SelectInputAlp("value", selectMod)
		return this.Component.Cas("span", attr, input)
	}

	panic("Unknown widget: " + operation.Widget)
}

func (this *Table) filterWidgets(mod *TableMod) goc.HTML {
	var widgets []goc.HTML
	for _, filter := range mod.AllowedFilters {
		for _, operation := range filter.Operations {
			widget := this.filterWidget(mod, filter, operation)
			widgets = append(widgets, widget)
		}
	}
	return this.Component.Cs("span", widgets...)
}

func (this *Table) createFilterButton(mod *TableMod) goc.HTML {
	return this.Component.Cas("span",
		goc.Attr{"x-on:click": "createFilter()"},
		this.Button.SecondaryIconLink(icons.IconFilter, "#", nil),
	)
}

func (this *Table) filter(mod *TableMod) goc.HTML {
	return this.Component.Cas("div", goc.Attr{"class": "flex flex-row space-x-2 items-center"},
		this.filterChoice(mod),
		this.filterOperations(mod),
		this.filterWidgets(mod),
		this.createFilterButton(mod),
	)
}

func (this *Table) filterChoice(mod *TableMod) goc.HTML {
	selectMod := this.SelectInput.Mod()
	for _, filter := range mod.AllowedFilters {
		selectMod.AddOption(filter.Name, filter.Text)
	}
	selectMod.SelectAttr = goc.Attr{"x-on:change": "chooseFilter", "id": "_new_filter"}
	return this.Component.Cs("span", this.SelectInput.H(selectMod))
}

func (this *Table) newFilter(mod *TableMod) goc.HTML {
	return this.Component.Cas("div", goc.Attr{"x-show": "showFilter", "class": "px-4"}, this.filter(mod))
}

func (this *Table) activeFilter(filter *dat.ActiveFilter) goc.HTML {
	if filter.Kind == dat.ActiveFilterKindExactStringValueFilter {
		return this.Component.Cav("div", goc.Attr{
			"class":      "mx-2 px-2 rounded border",
			"x-on:click": "removeFilter('" + filter.FilterKey + "', '" + filter.OperationKey + "')"},
			filter.FilterKey+" equals "+filter.Value,
		)
	}
	if filter.Kind == dat.ActiveFilterKindContainsStringValueFilter {
		return this.Component.Cav("div", goc.Attr{
			"class":      "mx-2 px-2 rounded border",
			"x-on:click": "removeFilter('" + filter.FilterKey + "', '" + filter.OperationKey + "')"},
			filter.FilterKey+" contains "+filter.Value,
		)
	}
	if filter.Kind == dat.ActiveFilterKindExactIntegerValueFilter {
		return this.Component.Cav("div", goc.Attr{
			"class":      "mx-2 px-2 rounded border",
			"x-on:click": "removeFilter('" + filter.FilterKey + "', '" + filter.OperationKey + "')"},
			filter.FilterKey+" equals "+filter.Value,
		)
	}
	if filter.Kind == dat.ActiveFilterKindExactBooleanValueFilter {
		return this.Component.Cav("div", goc.Attr{
			"class":      "mx-2 px-2 rounded border",
			"x-on:click": "removeFilter('" + filter.FilterKey + "', '" + filter.OperationKey + "')"},
			filter.FilterKey+" equals "+filter.Value,
		)
	}
	panic("Unknown active filter " + filter.Kind)
}

func (this *Table) activeFilters(mod *TableMod) goc.HTML {
	var filters []goc.HTML

	if mod.Pagination == nil {
		return goc.HTML{}
	}

	for _, filter := range mod.Pagination.ActiveFilters {
		filters = append(filters, this.activeFilter(filter))
	}

	return this.Component.Cs("div", filters...)
}

func (this *Table) filtersContainer(mod *TableMod) goc.HTML {
	if mod.HideFiltering {
		return goc.HTML{}
	}
	return this.Component.Ccs("div", "flex flex-row",
		this.newFilterButton(mod),
		this.newFilter(mod),
		this.activeFilters(mod),
	)
}

func (this *Table) headerActionsRow(mod *TableMod) goc.HTML {
	if mod.HideActionsHeader == true {
		return goc.HTML{}
	}

	leftButtons := this.leftButtons(mod)
	rightButtons := this.rightButtons(mod)
	title := goc.HTML{}
	if mod.Title != "" {
		title = this.Header.H2(mod.Title)
	}
	return this.Component.Ccs("div", "px-3 py-4 flex justify-between items-center border-b border-gray-200",
		title, leftButtons, rightButtons)
}

func (this *Table) tableHeader(mod *TableMod) goc.HTML {
	return this.Component.Cas("thead", goc.Attr{"class": "bg-gray-50", "role": aria.AriaRoleRowgroup},
		this.headerTitlesRow(mod),
	)
}

func (this *Table) footer(mod *TableMod) goc.HTML {
	return this.Component.Ccs("div", "border-t border-gray-200", this.footerRow(mod))
}

func (this *Table) footerRow(mod *TableMod) goc.HTML {
	return this.Component.Cs("div", this.footerCell(mod))
}

func (this *Table) footerCell(mod *TableMod) goc.HTML {
	return goc.H("div",
		goc.Attr{"class": "p-3"},
		this.footerContainer(mod),
	)
}

func (this *Table) footerContainer(mod *TableMod) goc.HTML {
	return this.Component.Ccs("div",
		"flex justify-between",
		this.paginationCount(mod),
		this.pagination(mod),
	)
}

func (this *Table) paginationCount(mod *TableMod) goc.HTML {
	total := strconv.Itoa(mod.Pagination.ItemsCount)
	visible := strconv.Itoa(len(mod.Items))
	return this.Component.Ccv("div", "flex text-scale-6 items-center font-normal",
		visible+" of "+total+" results",
	)
}

func (this *Table) pagination(mod *TableMod) goc.HTML {
	if mod.HidePagination {
		return goc.HTML{}
	}
	mod2 := this.Pagination.Mod()
	mod2.Pagination = mod.Pagination
	mod2.BaseUrl = mod.BaseUrl
	return this.Pagination.H(mod2)
}

func (this *Table) headerTitlesRow(mod *TableMod) goc.HTML {
	var columns []goc.HTML

	for _, header := range mod.Columns {
		columns = append(columns, this.headerElement(mod, header))
	}

	return this.Component.Cas("tr", goc.Attr{"role": aria.AriaRoleRow}, columns...)
}

func (this *Table) headerElement(mod *TableMod, column *Column) goc.HTML {
	if column.Hidden == true {
		return goc.HTML{}
	}

	css := "px-3 py-3 text-left text-xs font-semibold uppercase tracking-wide text-gray-800 "
	if column.Key == "_id" {
		attr := goc.Attr{
			"data-table-id":   mod.TableId,
			"data-table-type": "checkbox_ids_controller",
			"x-on:click":      "checkAllIds()",
		}
		booleanInput := this.BooleanInput.BooleanInputWithAttr("all", false, attr)
		return this.Component.Cas("th", goc.Attr{"class": css + column.Width, "role": aria.AriaRoleColumnheader}, booleanInput)
	}

	return this.Component.Cav("th", goc.Attr{"class": css + column.Width, "role": aria.AriaRoleColumnheader}, column.Label)
}

func (this *Table) row(item dat.Entity, mod *TableMod) goc.HTML {
	var cells []goc.HTML
	if mod.RowCell == nil {
		panic("Row cell is not declared in mod")
	}
	for key, column := range mod.Columns {
		if column.Hidden == true {
			continue
		}
		if column.Key == "_id" {
			cell := this.checkboxId(mod, key, item.GetId())
			cells = append(cells, cell)
			continue
		}

		cell := mod.RowCell(item, column)
		cells = append(cells, cell)
	}
	return this.Component.Cas("tr", goc.Attr{"class": "bg-white hover:bg-gray-50", "role": aria.AriaRoleRow}, cells...)
}

func (this *Table) checkboxId(mod *TableMod, key int, id string) goc.HTML {
	name := "id_" + strconv.Itoa(key+1)
	attr := goc.Attr{"data-table-id": mod.TableId, "data-table-type": "checkbox_id", "data-id": id}
	booleanInput := this.BooleanInput.BooleanInputWithAttr(name, false, attr)
	cell := this.Component.Cas("td", goc.Attr{"class": "px-3", "role": aria.AriaRoleCell}, booleanInput)
	return cell
}

func (this *Table) rowViaCells(mod *TableMod, row *CellsRow) goc.HTML {
	attr := goc.Attr{}
	for k, v := range row.Attr {
		attr[k] = v
	}
	attr["css"] = "bg-white hover:bg-gray-50"
	return this.Component.Cas("tr", attr, row.Cells...)
}

func (this *Table) emptyTable(mod *TableMod) goc.HTML {
	td := this.Component.Cav("td",
		goc.Attr{
			"colspan": strconv.Itoa(len(mod.Columns)),
			"class":   "text-center py-4 text text-scale-8",
		},
		"No items to show")
	return this.Component.Cs("tr", td)
}

func (this *Table) body(mod *TableMod) goc.HTML {
	if mod.RowsViaCells == true {
		return this.bodyViaCells(mod)
	}
	return this.bodyViaProvidedFunction(mod)
}

func (this *Table) bodyViaCells(mod *TableMod) goc.HTML {
	if len(mod.CellsRows) == 0 {
		return this.Component.Cs("tbody", this.emptyTable(mod))
	}
	var rows []goc.HTML
	for _, cellsRow := range mod.CellsRows {
		rows = append(rows, this.rowViaCells(mod, cellsRow))
	}
	return this.Component.Ccs("tbody", "divide-y divide-gray-200", rows...)
}

func (this *Table) bodyViaProvidedFunction(mod *TableMod) goc.HTML {
	if len(mod.Items) == 0 {
		return this.Component.Cs("tbody", this.emptyTable(mod))
	}
	var rows []goc.HTML
	for _, item := range mod.Items {
		rows = append(rows, this.row(item, mod))
	}
	return this.Component.Ccs("tbody", "divide-y divide-gray-200", rows...)
}

func (this *Table) rightButtons(mod *TableMod) goc.HTML {
	var subs []goc.HTML

	if len(mod.CheckboxActions) > 1 {
		button := this.Button.SecondaryIconLink(icons.IconFasSquareCheck, "#", nil)
		subs = append(subs, this.Dropdown.Dropdown(button, mod.CheckboxActions))
	}
	if mod.FullTableUrl != "" {
		button := this.Button.SecondaryIconLink(icons.IconFasTable, mod.FullTableUrl, nil)
		subs = append(subs, button)
	}
	if mod.NewEntityUrl != "" {
		button := this.Button.SecondaryLink("New", mod.NewEntityUrl, nil)
		subs = append(subs, button)
	}

	return this.Component.Dcs("flex flex-row space-x-2", subs...)
}

func (this *Table) hasCheckboxActions(mod *TableMod) bool {
	return len(mod.CheckboxActions) > 1
}

func (this *Table) leftButtons(mod *TableMod) goc.HTML {
	var set []goc.HTML

	if mod.HideFiltering == false {
		set = append(set, this.filtersContainer(mod), this.selectColumnsContainer(mod))
	}

	return this.Component.Ccs("div", "flex flex-row space-x-2", set...)
}

func (this *Table) selectColumnsContainer(mod *TableMod) goc.HTML {
	return this.Component.Ccs("div", "flex flex-row space-x-2 items-center", this.selectColumnsButton(),
		this.Component.Ti("showSelectColumns", this.selectColumnsForm(mod)))
}

func (this *Table) columns(mod *TableMod) []goc.HTML {
	var set []goc.HTML
	for _, column := range mod.Columns {
		set = append(set, this.columnField(mod, column))
	}
	return set
}

func (this *Table) selectColumnsForm(mod *TableMod) goc.HTML {
	return this.Component.Dcs("flex flex-row space-x-2", this.columns(mod)...)
}

func (this *Table) selectColumnsButton() goc.HTML {
	return this.Component.Cas("span",
		goc.Attr{"x-on:click": "toggleSelectColumns()", "class": "flex items-center"},
		this.Button.SecondaryIconLink(icons.IconFasTableColumns, "#", nil),
	)
}

func (this *Table) columnField(mod *TableMod, column *Column) goc.HTML {
	if column.Kind != ColumnKindData {
		return goc.HTML{}
	}
	return this.Component.Ccs("div", "flex flex-row space-x-2 items-center",
		this.columnCheckbox(mod, column),
		this.columnLabel(mod, column),
	)
}

func (this *Table) columnCheckbox(mod *TableMod, column *Column) goc.HTML {
	return this.Component.Cas("div", goc.Attr{"@click": "selectColumn('" + column.Key + "')"}, this.BooleanInput.BooleanInput(column.Key, !column.Hidden))
}

func (this *Table) columnLabel(mod *TableMod, column *Column) goc.HTML {
	return this.Component.Cv("span", column.Label)
}
