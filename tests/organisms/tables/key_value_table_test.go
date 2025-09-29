package tables

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/components/organisms/tables"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"testing"
)

func TestKeyValueTableKeyValueTable(t *testing.T) {
	kit := kitcla.New()
	component := kit.Organisms.Tables.KeyValueTable

	h := component.KeyValueTable(make([]*tables.KeyValueItem, 0), "", "")
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/key_value_table_key_value_table_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/key_value_table_key_value_table_1.html")
}
