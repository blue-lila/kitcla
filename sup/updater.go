package sup

import (
	"os"
	"path/filepath"
	"testing"
)

func UpdateEqualHtmlFromDataFile(t *testing.T, html string, filePath string) {
	basePath := filepath.Dir(filePath)
	err := os.MkdirAll(basePath, 0777)
	Poe(err)
	err = os.WriteFile(filePath, []byte(html), 0644)
	Poe(err)
}
