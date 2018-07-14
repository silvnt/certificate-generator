package main

import (
	"html/template"
	"net/http"

	"github.com/silvnt/certificate-generator/model/parser"
)

func main() {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		table := parser.ParseText(r.FormValue("students-list"))

		_ = table

		tmpl.Execute(w, struct{ Success bool }{true})
	})

	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", fs)

	http.ListenAndServe(":3000", nil)
}
