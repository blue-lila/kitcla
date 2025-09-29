package placeholder

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/components/atoms/card_bodies"
	"github.com/blue-lila/kitcla/components/atoms/tabs"
	"github.com/blue-lila/kitcla/components/organisms/tables"
	"github.com/blue-lila/kitcla/dat"
	"github.com/blue-lila/kitcla/goc"
	u "github.com/blue-lila/kitcla/goc_utils"
	"time"
)

type Plant struct {
	Id          string
	Name        string
	Species     string
	Age         int
	Height      float64
	IsEvergreen bool
}

func (this *Plant) GetId() string {
	return this.Id
}

func MakeItems(n int) []dat.Entity {
	entities := MakeEntities(n)
	items := make([]dat.Entity, len(entities))
	for i, v := range entities {
		items[i] = v
	}
	return items
}

func MakeEntities(n int) []*Plant {
	var plants []*Plant

	plantData := []Plant{
		{"1", "Aloe Vera", "Aloe", 3, 25.0, true},
		{"2", "Bamboo", "Bambusoideae", 5, 150.0, true},
		{"3", "Cactus", "Cactaceae", 7, 40.0, true},
		{"4", "Daisy", "Bellis", 2, 15.0, false},
		{"5", "Eucalyptus", "Eucalyptus", 6, 200.0, true},
		{"6", "Fern", "Polypodiopsida", 4, 50.0, true},
		{"7", "Ginseng", "Panax", 10, 30.0, false},
		{"8", "Hibiscus", "Hibiscus", 3, 60.0, false},
		{"9", "Ivy", "Hedera", 8, 100.0, true},
		{"10", "Jasmine", "Jasminum", 6, 120.0, false},
		{"11", "Kalanchoe", "Kalanchoe", 2, 25.0, true},
		{"12", "Lavender", "Lavandula", 5, 45.0, false},
		{"13", "Mint", "Mentha", 1, 20.0, true},
		{"14", "Narcissus", "Narcissus", 3, 30.0, false},
		{"15", "Orchid", "Orchidaceae", 7, 35.0, false},
		{"16", "Palm", "Arecaceae", 12, 250.0, true},
		{"17", "Quince", "Cydonia", 9, 180.0, false},
		{"18", "Rose", "Rosa", 4, 50.0, false},
		{"19", "Sunflower", "Helianthus", 1, 180.0, false},
		{"20", "Tulip", "Tulipa", 2, 40.0, false},
		{"21", "Umbrella Plant", "Schefflera", 5, 130.0, true},
		{"22", "Violet", "Viola", 1, 10.0, false},
		{"23", "Willow", "Salix", 15, 500.0, false},
		{"24", "Xerophyte", "Succulent", 6, 35.0, true},
		{"25", "Yucca", "Yucca", 7, 160.0, true},
		{"26", "Zelkova", "Zelkova", 14, 220.0, true},
	}

	for _, data := range plantData {
		plants = append(plants, &Plant{
			Id:          data.Id,
			Name:        data.Name,
			Species:     data.Species,
			Age:         data.Age,
			Height:      data.Height,
			IsEvergreen: data.IsEvergreen,
		})
	}

	if n > len(plants) {
		n = len(plants)
	}
	return plants[:n]
}

func RowCell(item dat.Entity, column *tables.Column) goc.HTML {
	kit := kitcla.New()
	plant := item.(*Plant)
	if column.Key == "name" {
		return u.Ccs("td", kit.Organisms.Tables.Table.TdCss(), kit.Atoms.Cells.TextCell.TextCell(plant.Name))
	}
	if column.Key == "height" {
		return u.Ccs("td", kit.Organisms.Tables.Table.TdCss(), kit.Atoms.Cells.DecimalCell.DecimalCell(plant.Height))
	}
	if column.Key == "evergreen" {
		return u.Ccs("td", kit.Organisms.Tables.Table.TdCss(), kit.Atoms.Cells.BooleanCell.BooleanCell(plant.IsEvergreen))
	}
	return u.Ccs("td", kit.Organisms.Tables.Table.TdCss(), u.Dv(column.Label))
}

func TableColumns() []*tables.Column {
	var columns []*tables.Column

	columns = append(columns, &tables.Column{
		Key:   "name",
		Label: "Name",
	})
	columns = append(columns, &tables.Column{
		Key:   "height",
		Label: "Height",
	})
	columns = append(columns, &tables.Column{
		Key:   "evergreen",
		Label: "Evergreen",
	})
	columns = append(columns, &tables.Column{
		Key:   "actions",
		Label: "Actions",
	})
	return columns
}

func GocHtmlValue() goc.HTML {
	return goc.HTML{}
}

func CardTitle() string {
	return "Plant it"
}

func CardContent() goc.HTML {
	return u.Dcs("p-4 border-2 border-dashed", u.Dv("Hello World"))
}

func GocAttrValue() goc.Attr {
	return goc.Attr{}
}

func TimeValue() time.Time {
	return time.Time{}
}

func ButtonValue() goc.HTML {
	return u.Ccv("", "My button")
}

func Button2Value() goc.HTML {
	return u.Ccv("", "My button 2")
}

func ButtonsValue() []goc.HTML {
	return []goc.HTML{ButtonValue(), Button2Value()}
}

func LeftSideValue() *card_bodies.Side {
	return &card_bodies.Side{
		Ratio:   "2/3",
		Content: u.Cv("div", "Hello World"),
	}
}

func RightSideValue() *card_bodies.Side {
	return &card_bodies.Side{
		Ratio:   "1/3",
		Content: u.Cv("div", "Hello World 2"),
	}
}

func GetTabModValue() *tabs.TabMod {
	mod := &tabs.TabMod{}
	mod.Items = GetTabItemValue()
	return mod
}

func GetTabItemValue() []*tabs.TabItem {
	var set []*tabs.TabItem
	set = append(set, &tabs.TabItem{
		Text:   "Tab 1",
		Value:  "#",
		Active: true,
	})
	set = append(set, &tabs.TabItem{
		Text:   "Tab 2",
		Value:  "#",
		Active: false,
	})
	set = append(set, &tabs.TabItem{
		Text:   "Tab 3",
		Value:  "#",
		Active: false,
	})
	return set
}
