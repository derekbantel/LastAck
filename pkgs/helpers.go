package pkgs

import (
	"html/template"
	"net/http"
	"path/filepath"
)

const templateDir = "./templates"
const templateLayout = "index.html"

func renderTemplates(w http.ResponseWriter, templateName string) {
	tmpl, err := template.ParseFiles(filepath.Join(templateDir, templateLayout), filepath.Join(templateDir, templateName))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, templateLayout, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
