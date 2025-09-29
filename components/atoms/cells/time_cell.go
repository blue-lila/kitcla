package cells

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/goc"

	"time"
)

type TimeCell struct {
	Component *components.Component
}

type TimeCellMod struct {
	Value  time.Time
	Format string
}

const TimeCellModLayoutDMYHM = "02/01/2006 15:04"
const TimeCellModLayoutDMY = "02/01/2006"

func (this *TimeCell) DayMonthYear(value time.Time) goc.HTML {
	return this.H(&TimeCellMod{Value: value, Format: TimeCellModLayoutDMY})
}

func (this *TimeCell) TimeCellFormated(value time.Time, format string) goc.HTML {
	return this.H(&TimeCellMod{Value: value, Format: format})
}

func (this *TimeCell) TimeCell(value time.Time) goc.HTML {
	return this.H(&TimeCellMod{Value: value, Format: TimeCellModLayoutDMYHM})
}

func (this *TimeCell) H(mod *TimeCellMod) goc.HTML {
	return this.Component.Cv("span", mod.Value.Format(mod.Format))
}
