package main

import (
	"html/template"
	"log"
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
		log.Println(r.FormValue("input-table"))
		table, headers, err := parser.ParseText(r.FormValue("input-table"))

		if err != nil {
			log.Println(err)
		} else {
			parsed, err := parser.ParseTable(table, headers, certifText)

			if err != nil {
				log.Println(err)
			}

			_ = parsed
		}

		home.Execute(w, struct{ Success bool }{true})
	})

	http.Handle("/static/", http.FileServer(http.Dir("assets/")))

	http.ListenAndServe(":3000", nil)
}
