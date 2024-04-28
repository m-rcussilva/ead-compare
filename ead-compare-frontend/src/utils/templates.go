package utils

import (
	"net/http"
	"text/template"
)

var templts *template.Template

func LoadTemplates() {
	templts = template.Must(template.ParseGlob("views/*.html"))
}

func ExecTemplate(w http.ResponseWriter, templ string, datas interface{}) {
	templts.ExecuteTemplate(w, templ, datas)
}
