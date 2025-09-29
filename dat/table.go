package dat

type Filter struct {
	Name       string
	Text       string
	Operations []*FilterOperation
}

type FilterOperation struct {
	Name   string
	Text   string
	Widget string
}

type Pagination struct {
	ItemsCount    int
	PerPage       int
	CurrentPage   int
	ActiveFilters []*ActiveFilter
}

const ActiveFilterKindExactStringValueFilter = "ExactStringValueFilter"
const ActiveFilterKindContainsStringValueFilter = "ContainsStringValueFilter"
const ActiveFilterKindExactIntegerValueFilter = "ExactIntegerValueFilter"
const ActiveFilterKindExactBooleanValueFilter = "ExactBooleanValueFilter"

type ActiveFilter struct {
	Kind         string
	FilterKey    string
	OperationKey string
	Value        string
}
