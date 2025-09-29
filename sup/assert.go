package sup

import (
	"os"
	"testing"
)

func AssertEqualHtmlFromDataFile(t *testing.T, html string, filepath string) {
	b, err := os.ReadFile(filepath)
	Poe(err)
	if string(b) != html {
		println("Html not equal to data from file " + filepath)
		t.FailNow()
	}
}
