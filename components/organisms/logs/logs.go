package logs

import (
	"github.com/blue-lila/kitcla/goc"
)

type Logs struct {
}

type LogsMod struct {
}

func (this *Logs) H(mod *LogsMod) goc.HTML {
	return goc.HTML{}
}

func (this *Logs) Mod() *LogsMod {
	return &LogsMod{}
}
