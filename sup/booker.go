package sup

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

func AddPage(name string, html string) {
	filePath2 := testFilePath()
	println(filePath2)

	ss := strings.Split(filePath2, "/")
	componentName := strings.TrimSuffix(ss[len(ss)-1], "_test.go")
	subKind := ss[len(ss)-2]
	kind := ss[len(ss)-3]

	rootDir, err := getProjectRoot()
	Poe(err)
	fileName := ToSnakeCase(name)
	htmlFilePath := filepath.Join(rootDir, "docs", "pages", kind, subKind, componentName, fileName+".html")
	jsonFilePath := filepath.Join(rootDir, "docs", "pages", kind, subKind, componentName, fileName+".json")
	basePath := filepath.Dir(htmlFilePath)
	err = os.MkdirAll(basePath, 0777)
	Poe(err)
	err = os.WriteFile(htmlFilePath, []byte(html), 0644)
	Poe(err)
	s := struct {
		Name          string
		ComponentName string
		Kind          string
		SubKind       string
		TestFilePath  string
		HtmlFilePath  string
	}{
		Name:          name,
		ComponentName: componentName,
		Kind:          kind,
		SubKind:       subKind,
		TestFilePath:  strings.TrimPrefix(testFilePath(), rootDir),
		HtmlFilePath:  strings.TrimPrefix(htmlFilePath, rootDir),
	}
	b, err := json.Marshal(s)
	err = os.WriteFile(jsonFilePath, b, 0644)
	Poe(err)
}
