package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// we createe a renderring function in go  to render tmpl files
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parseTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parseTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}
