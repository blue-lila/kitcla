package goc_debug

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/blue-lila/kitcla/goc"
	"runtime"
)

type Debugging struct {
	Items []goc.Debug
}

func ExportDebug(root goc.HTML) *Debugging {
	debugging := &Debugging{}
	exportDebugNode(root, debugging)
	return debugging
}

func exportDebugNode(h goc.HTML, debugging *Debugging) {
	for _, v := range h.Attrs {
		switch v := v.(type) {
		case goc.Debug:
			debugging.Items = append(debugging.Items, v)
		case []goc.HTML:
			for _, sub := range v {
				exportDebugNode(sub, debugging)
			}
		case goc.HTML:
			exportDebugNode(v, debugging)
		}
	}
}

func DebugInfo() goc.Debug {
	return goc.Debug{
		Id:   randomIdentifier(),
		Path: getCallerInfo(),
	}
}

func getCallerInfo() string {
	pc, file, line, ok := runtime.Caller(3)
	if !ok {
		return ""
	}
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return ""
	}
	return fmt.Sprintf("%s:%d (%s)", file, line, fn.Name())
}

func randomIdentifier() string {
	bytes := make([]byte, 4)
	_, err := rand.Read(bytes)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(bytes)
}
