package tmpl

import (
	"log"
	"text/template"
)

var Timestamp = loadFromFile("tmpl/timestamp.tmpl")

func loadFromFile(fileName string) *template.Template {
	tmpl, err := template.ParseFiles(fileName)
	if err != nil {
		log.Panicf("Could not load template file \"%s\"\n%v", fileName, err)
	}
	return tmpl
}
