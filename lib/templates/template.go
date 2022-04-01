package templates

import (
	"html/template"
	"io"
)

const (
	TemplateFile = "lib/templates/page.html"
)

var TemplateCache *template.Template

func NewTemplateCache(file string) {
	TemplateCache = template.Must(template.ParseFiles(file))
}

func RenderTemplate(file string, templateData interface{}, writer io.Writer) error {
	err := TemplateCache.ExecuteTemplate(writer, file+".html", templateData)
	return err
}
