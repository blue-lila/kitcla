package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const basePath = "/home/<user>/go/src/kitcla/tests"
const jsonFilePath = "/home/<user>go/src/gen-x/docs/kit/spec.json"

type Argument struct {
	Name    string `json:"name"`
	Kind    string `json:"kind"`
	Default string `json:"default"`
}

type Call struct {
	Name string     `json:"name"`
	Args []Argument `json:"args"`
}

type Component struct {
	Name    string `json:"name"`
	Kind    string `json:"kind"`
	SubKind string `json:"sub_kind"`
	Calls   []Call `json:"calls"`
}

type DataStructure struct {
	Components []Component `json:"components"`
}

func parseJSON(jsonData string) (*DataStructure, error) {
	var parsedData DataStructure
	if err := json.Unmarshal([]byte(jsonData), &parsedData); err != nil {
		return nil, err
	}
	return &parsedData, nil
}

func generateTestFunctions(component Component) string {
	var testFunctions strings.Builder

	for _, call := range component.Calls {
		testFunctions.WriteString(generateTestFunction(component, call))
	}
	return testFunctions.String()
}

func generateTestFunction(component Component, call Call) string {
	funcName := toPascalCase(call.Name)
	dataFile := fmt.Sprintf("./data/%s_%s_1.html", component.Name, call.Name)

	argsList := generateArgsList(call.Args)

	functionTemplate := `
func Test__TEST_FUNC_NAME__(t *testing.T) {
	kit := kitcla.New()
	component := kit.__KIT_KIND__.__KIT_SUB_KIND__.__COMPONENT_NAME__

	h := component.__CALL_NAME__(__ARGS__)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "__DATA_FILE__")
	sup.AssertEqualHtmlFromDataFile(t, html, "__DATA_FILE__")
}
`

	replacements := map[string]string{
		"__TEST_FUNC_NAME__": toPascalCase(component.Name) + funcName,
		"__KIT_KIND__":       toPascalCase(component.Kind),
		"__KIT_SUB_KIND__":   toPascalCase(component.SubKind),
		"__COMPONENT_NAME__": toPascalCase(component.Name),
		"__CALL_NAME__":      funcName,
		"__ARGS__":           argsList,
		"__DATA_FILE__":      dataFile,
	}

	return replacePlaceholders(functionTemplate, replacements)
}

func generateArgsList(args []Argument) string {
	var argsList []string
	for _, arg := range args {
		argsList = append(argsList, arg.Default)
	}
	return strings.Join(argsList, ", ")
}

func generateTestContent(component Component) string {
	baseTemplate := `package __PACKAGE__

import (
	"testing"
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
)

__TEST_FUNCTIONS__
`
	testFunctions := generateTestFunctions(component)
	replacements := map[string]string{
		"__PACKAGE__":        component.SubKind,
		"__TEST_FUNCTIONS__": testFunctions,
	}

	return replacePlaceholders(baseTemplate, replacements)
}

func replacePlaceholders(template string, replacements map[string]string) string {
	for placeholder, value := range replacements {
		template = strings.ReplaceAll(template, placeholder, value)
	}
	return template
}

func createTestFile(component Component) error {
	dirPath := filepath.Join(basePath, component.Kind, component.SubKind)
	filePath := filepath.Join(dirPath, component.Name+"_test.go")

	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return fmt.Errorf("erreur lors de la création du dossier %s: %w", dirPath, err)
	}

	testContent := generateTestContent(component)

	if err := os.WriteFile(filePath, []byte(testContent), 0644); err != nil {
		return fmt.Errorf("erreur lors de l'écriture du fichier %s: %w", filePath, err)
	}

	fmt.Println("Fichier créé :", filePath)
	return nil
}

func toPascalCase(name string) string {
	words := strings.Split(name, "_")
	for i := range words {
		words[i] = strings.Title(words[i])
	}
	return strings.Join(words, "")
}

func parseJSONFromFile(filePath string) (*DataStructure, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la lecture du fichier JSON: %w", err)
	}

	var parsedData DataStructure
	if err := json.Unmarshal(data, &parsedData); err != nil {
		return nil, fmt.Errorf("erreur lors du parsing JSON: %w", err)
	}
	return &parsedData, nil
}

func main() {
	parsedData, err := parseJSONFromFile(jsonFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err != nil {
		fmt.Println("Erreur lors du parsing JSON:", err)
		return
	}

	for _, component := range parsedData.Components {
		if err := createTestFile(component); err != nil {
			fmt.Println(err)
		}
	}
}
