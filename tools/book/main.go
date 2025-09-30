package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// =============================================================================
// Data Structures
// =============================================================================

type ComponentData struct {
	Name          string `json:"Name"`
	ComponentName string `json:"ComponentName"`
	Kind          string `json:"Kind"`
	SubKind       string `json:"SubKind"`
	TestFilePath  string `json:"TestFilePath"`
	HtmlFilePath  string `json:"HtmlFilePath"`
}

type TemplateData struct {
	Menu           map[string]map[string]map[string][]ComponentData
	CurrentPage    string
	PageTitle      string
	PageBreadcrumb string
	IframeSrc      string
	ShowWelcome    bool
	IndexPages     []ComponentData
	DirectHTML     template.HTML
	ShowDirectHTML bool
}

type BookServer struct {
	components []ComponentData
	docsPath   string
	assetsPath string
	tmpl       *template.Template
}

// =============================================================================
// Constructor
// =============================================================================

func NewBookServer() *BookServer {
	return &BookServer{
		docsPath:   "./docs/pages",
		assetsPath: "./assets",
	}
}

// =============================================================================
// Component Data Management
// =============================================================================

func (s *BookServer) loadComponents() error {
	s.components = []ComponentData{}

	err := filepath.Walk(s.docsPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.HasSuffix(path, ".json") {
			data, err := os.ReadFile(path)
			if err != nil {
				log.Printf("Error reading file %s: %v", path, err)
				return nil
			}

			var component ComponentData
			if err := json.Unmarshal(data, &component); err != nil {
				log.Printf("Error parsing JSON from %s: %v", path, err)
				return nil
			}

			s.components = append(s.components, component)
		}

		return nil
	})

	return err
}

func (s *BookServer) organizeComponents() map[string]map[string]map[string][]ComponentData {
	organized := make(map[string]map[string]map[string][]ComponentData)

	for _, component := range s.components {
		kind := component.Kind
		subKind := component.SubKind
		componentName := component.ComponentName

		// Skip category-level overviews (atoms/index.html) - they are linked from category summary
		if subKind == "" && componentName == "" {
			continue
		}

		// For subcategory-level overviews (atoms/buttons/index.html where componentName == subKind),
		// we keep them but mark them specially so they can be used as the subcategory overview
		// We'll store them under a special empty componentName key
		if subKind != "" && componentName == subKind {
			if organized[kind] == nil {
				organized[kind] = make(map[string]map[string][]ComponentData)
			}
			if organized[kind][subKind] == nil {
				organized[kind][subKind] = make(map[string][]ComponentData)
			}
			// Store overview under empty componentName so it doesn't show as a regular item
			organized[kind][subKind][""] = append(organized[kind][subKind][""], component)
			continue
		}

		if organized[kind] == nil {
			organized[kind] = make(map[string]map[string][]ComponentData)
		}
		if organized[kind][subKind] == nil {
			organized[kind][subKind] = make(map[string][]ComponentData)
		}
		organized[kind][subKind][componentName] = append(organized[kind][subKind][componentName], component)
	}

	return organized
}

func (s *BookServer) findComponentByPath(pagePath string) *ComponentData {
	for _, component := range s.components {
		componentPath := strings.TrimPrefix(component.HtmlFilePath, "/docs/pages/")
		componentPath = strings.TrimSuffix(componentPath, ".html")
		if componentPath == pagePath {
			return &component
		}
	}
	return nil
}

// =============================================================================
// Utilities
// =============================================================================

func (s *BookServer) getTopLevelIndexPages() []ComponentData {
	var topLevelPages []ComponentData
	for _, component := range s.components {
		// Only include top-level category overviews like /atoms/index.html, /molecules/index.html
		// Path format: /docs/pages/{category}/index.html
		path := component.HtmlFilePath
		if strings.HasSuffix(path, "/index.html") {
			// Count the slashes to determine depth
			// /docs/pages/atoms/index.html has 4 slashes (top-level)
			// /docs/pages/atoms/buttons/index.html has 5 slashes (subcategory)
			slashCount := strings.Count(path, "/")
			if slashCount == 4 {
				topLevelPages = append(topLevelPages, component)
			}
		}
	}
	return topLevelPages
}

func (s *BookServer) getComponentHTML(componentPath string) template.HTML {
	htmlPath := filepath.Join(s.docsPath, componentPath+".html")
	content, err := os.ReadFile(htmlPath)
	if err != nil {
		return ""
	}
	return template.HTML(content)
}

func (s *BookServer) capitalize(str string) string {
	if str == "" {
		return str
	}
	return strings.ToUpper(string(str[0])) + strings.ReplaceAll(str[1:], "_", " ")
}

func (s *BookServer) setupTemplate() error {
	funcMap := template.FuncMap{
		"capitalize":  s.capitalize,
		"stripPrefix": func(prefix, str string) string { return strings.TrimPrefix(str, prefix) },
		"stripSuffix": func(suffix, str string) string { return strings.TrimSuffix(str, suffix) },
		"toLower":     strings.ToLower,
	}

	var err error
	s.tmpl, err = template.New("index.html").Funcs(funcMap).ParseFiles("tools/book/assets/index.html")
	if err != nil {
		return fmt.Errorf("failed to load template: %v", err)
	}

	return nil
}

func (s *BookServer) setupRoutes() {
	http.HandleFunc("/", s.handleIndex)
	http.HandleFunc("/static/", s.handleStatic)
	http.HandleFunc("/component/", s.handleComponent)
	http.HandleFunc("/empty", s.handleEmpty)
}

// =============================================================================
// Server Management
// =============================================================================

func (s *BookServer) Start(port int) error {
	// Initialize template system
	if err := s.setupTemplate(); err != nil {
		return err
	}

	// Load component data
	if err := s.loadComponents(); err != nil {
		return fmt.Errorf("failed to load components: %v", err)
	}

	log.Printf("Loaded %d components", len(s.components))

	// Setup HTTP routes
	s.setupRoutes()

	// Start server
	addr := fmt.Sprintf(":%d", port)
	log.Printf("Starting KitCLA Book server on http://localhost%s", addr)

	return http.ListenAndServe(addr, nil)
}

// =============================================================================
// Main Entry Point
// =============================================================================

func main() {
	server := NewBookServer()

	// Parse port from environment variable or use default
	port := 7000
	if portEnv := os.Getenv("PORT"); portEnv != "" {
		if p, err := strconv.Atoi(portEnv); err == nil {
			port = p
		} else {
			log.Printf("Invalid PORT environment variable '%s', using default port %d", portEnv, port)
		}
	}

	if err := server.Start(port); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
