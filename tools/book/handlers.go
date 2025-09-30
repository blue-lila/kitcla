package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// =============================================================================
// HTTP Handlers
// =============================================================================

func (s *BookServer) handleIndex(w http.ResponseWriter, r *http.Request) {
	currentPage := r.URL.Query().Get("page")

	data := TemplateData{
		Menu:           s.organizeComponents(),
		CurrentPage:    currentPage,
		PageTitle:      "KitCLA Component Book",
		PageBreadcrumb: "Explore component overviews or browse the sidebar for individual components",
		IframeSrc:      "",
		ShowWelcome:    currentPage == "",
		IndexPages:     s.getTopLevelIndexPages(),
	}

	if currentPage != "" {
		if component := s.findComponentByPath(currentPage); component != nil {
			data.PageTitle = component.Name
			data.PageBreadcrumb = fmt.Sprintf("%s > %s > %s",
				s.capitalize(component.Kind),
				s.capitalize(component.SubKind),
				s.capitalize(component.ComponentName))

			// Check if this is a top-level overview page (e.g., atoms/index, molecules/index)
			// These should be rendered directly, not in iframe
			isOverviewPage := strings.HasSuffix(currentPage, "/index") && component.SubKind == ""

			if isOverviewPage {
				data.IframeSrc = ""
				data.ShowWelcome = false
				data.DirectHTML = s.getComponentHTML(currentPage)
				data.ShowDirectHTML = true
			} else {
				data.IframeSrc = "/component/" + currentPage
				data.ShowWelcome = false
				data.ShowDirectHTML = false
			}
		}
	}

	if err := s.tmpl.Execute(w, data); err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *BookServer) handleStatic(w http.ResponseWriter, r *http.Request) {
	file := strings.TrimPrefix(r.URL.Path, "/static/")

	// Map public names to actual asset file paths
	switch file {
	case "tailwind.min.css":
		file = "libs/tailwind.min.css"
	case "alpine.3.8.1.min.js":
		file = "libs/alpine.3.8.1.min.js"
	}

	filePath := filepath.Join(s.assetsPath, file)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.NotFound(w, r)
		return
	}

	// Set content type based on file extension
	if strings.HasSuffix(filePath, ".css") {
		w.Header().Set("Content-Type", "text/css")
	} else if strings.HasSuffix(filePath, ".js") {
		w.Header().Set("Content-Type", "application/javascript")
	}

	http.ServeFile(w, r, filePath)
}

func (s *BookServer) handleComponent(w http.ResponseWriter, r *http.Request) {
	componentPath := strings.TrimPrefix(r.URL.Path, "/component/")

	// Check the request type based on suffix
	var requestType string
	if strings.HasSuffix(componentPath, "/html") {
		requestType = "html"
		componentPath = strings.TrimSuffix(componentPath, "/html")
	} else if strings.HasSuffix(componentPath, "/arguments") {
		requestType = "arguments"
		componentPath = strings.TrimSuffix(componentPath, "/arguments")
	} else {
		requestType = "preview"
	}

	w.Header().Set("Content-Type", "text/html")

	if requestType == "arguments" {
		// Find component data to get test file path
		component := s.findComponentByPath(componentPath)
		if component == nil {
			http.NotFound(w, r)
			return
		}

		// Check if TestFilePath is empty
		if component.TestFilePath == "" {
			http.Error(w, "Test file not found - this page was not generated from a test", http.StatusNotFound)
			return
		}

		// Read test file content
		testPath := strings.TrimPrefix(component.TestFilePath, "/")
		if _, err := os.Stat(testPath); os.IsNotExist(err) {
			http.Error(w, "Test file not found", http.StatusNotFound)
			return
		}

		testContent, err := os.ReadFile(testPath)
		if err != nil {
			http.Error(w, "Error reading test file", http.StatusInternalServerError)
			return
		}

		s.serveArgumentsView(w, testContent)
		return
	}

	// For html and preview requests, read the HTML file
	htmlPath := filepath.Join(s.docsPath, componentPath+".html")
	if _, err := os.Stat(htmlPath); os.IsNotExist(err) {
		http.NotFound(w, r)
		return
	}

	content, err := os.ReadFile(htmlPath)
	if err != nil {
		http.Error(w, "Error reading HTML file", http.StatusInternalServerError)
		return
	}

	if requestType == "html" {
		s.serveHtmlView(w, content)
	} else {
		s.servePreview(w, content)
	}
}

func (s *BookServer) handleEmpty(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	indexPages := s.getTopLevelIndexPages()
	_, _ = w.Write([]byte(s.getEmptyPageHTML(indexPages)))
}
