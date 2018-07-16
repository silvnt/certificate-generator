package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/silvnt/certifier-go/model/parser"
)

func main() {
	home := template.Must(template.ParseFiles("templates/home.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			home.Execute(w, nil)
			return
		}

		certifText := r.FormValue("input-editor")
		table := parser.ParseText(r.FormValue("input-table"))

		fmt.Println(certifText)
		fmt.Println(table)

		home.Execute(w, struct{ Success bool }{true})
	})

	http.Handle("/static/", http.FileServer(http.Dir("assets/")))

	http.ListenAndServe(":3000", nil)
}
