package grids

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/components/atoms/icons"
	"github.com/blue-lila/kitcla/components/atoms/links"
	"github.com/blue-lila/kitcla/goc"

	"strconv"
)

type IconCardGrid struct {
	Icon      *icons.Icon
	Link      *links.Link
	Component *components.Component
}

type IconCardGridMod struct {
	Cards []*IconCardGridCard
}

type IconCardGridCard struct {
	Name  string
	Url   string
	Count int
	Icon  string
}

func (this *IconCardGrid) H(mod *IconCardGridMod) goc.HTML {
	var cards []goc.HTML

	for _, card := range mod.Cards {
		cards = append(cards, this.gridCard(card.Name, card.Icon, card.Url, card.Count))
	}

	return goc.H("div", goc.Attr{"class": "grid grid-cols-1 md:grid-cols-3 lg:grid-cols-4 gap-4"}, cards)
}

func (this *IconCardGrid) IconCardGrid(mod *IconCardGridMod) goc.HTML {
	return this.H(mod)
}

func (this *IconCardGrid) gridCard(name string, icon string, url string, count int) goc.HTML {
	return goc.H("div", goc.Attr{"class": "bg-scale-0 rounded-lg shadow flex flex-col"},
		this.body(name, icon, count, url),
	)
}

func (this *IconCardGrid) cardTitle(name string, url string) goc.HTML {
	return goc.H("div", goc.Attr{"class": "text-xl"},
		this.Link.Link(name, url),
	)
}

func (this *IconCardGrid) cardCount(count int) goc.HTML {
	countS := strconv.Itoa(count)
	return goc.H("div", goc.Attr{"class": "text-scale-6 text font-medium"},
		countS,
	)
}

func (this *IconCardGrid) cardIcon(icon string) goc.HTML {
	return this.Component.Wrap("bg-primary-1 text-scale-0 w-9 h-9 flex items-center justify-center rounded", this.Icon.IconWithSize(icon, "6"))
}

func (this *IconCardGrid) text(name string, count int, url string) goc.HTML {
	return goc.H("div", goc.Attr{"class": "flex flex-row items-baseline space-x-2"},
		this.cardTitle(name, url),
		this.cardCount(count),
	)
}

func (this *IconCardGrid) body(name string, icon string, count int, url string) interface{} {
	return goc.H("div", goc.Attr{"class": "my-5 mx-6  flex flex-row space-x-4 items-center"},
		this.cardIcon(icon),
		this.text(name, count, url),
	)
}
