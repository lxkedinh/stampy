package tmpl

import (
	"log"
	"text/template"
)

func LoadFromFile(fileName string) *template.Template {
	tmpl, err := template.ParseFiles(fileName)
	if err != nil {
		log.Panicf("Could not load template file \"%s\"", fileName)
	}
	return tmpl
}
