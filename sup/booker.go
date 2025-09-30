package sup

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

func AddIndexPage(name string, html string) {
	filePath2 := testFilePath()

	ss := strings.Split(filePath2, "/")
	// Find "tests" in the path and extract the next element as kind
	// Path: /path/to/tests/{kind}/{subkind}/{component}_test.go
	var kind string
	for i, part := range ss {
		if part == "tests" && i+1 < len(ss) {
			kind = ss[i+1]
			break
		}
	}

	rootDir, err := getProjectRoot()
	Poe(err)

	// Create index files at the subKind level (e.g., atoms/buttons/index.html)
	htmlFilePath := filepath.Join(rootDir, "docs", "pages", kind, name, "index.html")
	jsonFilePath := filepath.Join(rootDir, "docs", "pages", kind, name, "index.json")
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
		Name:          strings.Title(name) + " Overview",
		ComponentName: name,
		Kind:          kind,
		SubKind:       name,
		TestFilePath:  strings.TrimPrefix(testFilePath(), rootDir),
		HtmlFilePath:  strings.TrimPrefix(htmlFilePath, rootDir),
	}
	b, err := json.Marshal(s)
	Poe(err)
	err = os.WriteFile(jsonFilePath, b, 0644)
	Poe(err)
}

func AddCustomPage(html string, htmlFilePath string, jsonFilePath string) {
	rootDir, err := getProjectRoot()
	Poe(err)

	// Handle relative paths
	if !filepath.IsAbs(htmlFilePath) {
		htmlFilePath = filepath.Join(rootDir, htmlFilePath)
	}
	if !filepath.IsAbs(jsonFilePath) {
		jsonFilePath = filepath.Join(rootDir, jsonFilePath)
	}

	basePath := filepath.Dir(htmlFilePath)
	err = os.MkdirAll(basePath, 0777)
	Poe(err)
	err = os.WriteFile(htmlFilePath, []byte(html), 0644)
	Poe(err)

	// Extract metadata from file paths
	relHtmlPath := strings.TrimPrefix(htmlFilePath, rootDir+"/")
	parts := strings.Split(strings.TrimPrefix(relHtmlPath, "docs/pages/"), "/")

	var kind, subKind, componentName, name string
	if len(parts) >= 1 {
		kind = parts[0]
	}
	if len(parts) >= 2 && parts[1] != "index.html" {
		subKind = parts[1]
	}
	if len(parts) >= 3 && parts[2] != "index.html" {
		componentName = parts[2]
	}

	// Generate name from the last meaningful part
	if componentName != "" {
		name = strings.Title(componentName)
	} else if subKind != "" {
		name = strings.Title(subKind) + " Overview"
	} else {
		name = strings.Title(kind) + " Overview"
	}

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
		HtmlFilePath:  "/" + relHtmlPath,
	}
	b, err := json.Marshal(s)
	Poe(err)
	err = os.WriteFile(jsonFilePath, b, 0644)
	Poe(err)
}

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
	Poe(err)
	err = os.WriteFile(jsonFilePath, b, 0644)
	Poe(err)
}
