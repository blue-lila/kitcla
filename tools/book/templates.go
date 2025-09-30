package main

import (
	"bytes"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// =============================================================================
// HTML Template Generators
// =============================================================================

type EmptyPageData struct {
	IndexPages []ComponentData
}

func (s *BookServer) serveHtmlView(w http.ResponseWriter, content []byte) {
	html, err := s.renderTemplate("html_view.html", string(content))
	if err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = w.Write([]byte(html))
}

func (s *BookServer) serveArgumentsView(w http.ResponseWriter, content []byte) {
	html, err := s.renderTemplate("arguments_view.html", string(content))
	if err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = w.Write([]byte(html))
}

func (s *BookServer) servePreview(w http.ResponseWriter, content []byte) {
	html, err := s.renderTemplate("preview.html", template.HTML(content))
	if err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = w.Write([]byte(html))
}

func (s *BookServer) getEmptyPageHTML(indexPages []ComponentData) string {
	data := EmptyPageData{
		IndexPages: indexPages,
	}

	html, err := s.renderTemplate("empty_page.html", data)
	if err != nil {
		return "Template error: " + err.Error()
	}
	return html
}

func (s *BookServer) renderTemplate(templateName string, data interface{}) (string, error) {
	templatePath := filepath.Join("tools/book/templates", templateName)

	templateContent, err := os.ReadFile(templatePath)
	if err != nil {
		return "", err
	}

	funcMap := template.FuncMap{
		"capitalize":  s.capitalize,
		"stripPrefix": func(prefix, str string) string { return strings.TrimPrefix(str, prefix) },
		"stripSuffix": func(suffix, str string) string { return strings.TrimSuffix(str, suffix) },
		"toLower":     strings.ToLower,
		"safeHTML":    func(str string) template.HTML { return template.HTML(str) },
	}

	tmpl, err := template.New(templateName).Funcs(funcMap).Parse(string(templateContent))
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
