package controller

import (
	"bytes"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/silvnt/certifier-go/model/mailer"
	"github.com/silvnt/certifier-go/model/parser"
	"github.com/silvnt/certifier-go/model/pdfgenerator"
)

// NewPage redirected
type NewPage struct {
	ThisIs bool
	Msg    string
}

// Redirect to a new a page
func Redirect(w http.ResponseWriter, template *template.Template, msg string) {
	template.Execute(w, NewPage{
		ThisIs: true,
		Msg:    msg,
	})
}

// StartServer and execute requested operations
func StartServer(address string, homeTemplateName string) {

	home := template.Must(template.ParseFiles("templates/" + homeTemplateName))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			home.Execute(w, nil)
			return
		}

		log.Println("Processo de geração e envio iniciado!")

		// Etapa background
		bg, bgHeader, err := r.FormFile("input-bg")
		if err != nil {
			log.Println(err)
			Redirect(w, home, err.Error()+". Retorne a página anterior.")
		}

		bgStatus := false
		var bgDir string

		bgBuf := bytes.NewBuffer(nil)
		if _, err := io.Copy(bgBuf, bg); err != nil {
			log.Println(err)
			Redirect(w, home, err.Error()+". Retorne a página anterior.")
		}

		if bgHeader.Filename != "" {
			filename := strings.Split(bgHeader.Filename, ".")
			bgDir = "static/bg." + filename[len(filename)-1]
			if err = ioutil.WriteFile("assets/"+bgDir, bgBuf.Bytes(),
				0700); err != nil {
				log.Println(err)
				Redirect(w, home, err.Error()+". Retorne a página anterior.")
			}

			bgStatus = true
			defer os.Remove("assets/" + bgDir)
		}

		// Etapa do Parser - Transformação do texto em tabela
		table, tableHeaders, err := parser.ParseText(r.FormValue("input-table"))
		if err != nil {
			log.Println(err)
			Redirect(w, home, err.Error()+". Retorne a página anterior.")
		}

		certifText := r.FormValue("input-editor")

		// Etapa do Parser - Substituição das tags por texto
		parsed, err := parser.ParseTable(table, tableHeaders, certifText)
		if err != nil {
			log.Println(err)
			Redirect(w, home, err.Error()+". Retorne a página anterior.")
		}

		// Etapa de geração dos certificados
		pdfs, err := pdfgenerator.Generate(parsed, bgStatus, address+"/"+bgDir)
		if err != nil {
			log.Println(err)
			Redirect(w, home, err.Error()+". Retorne a página anterior.")
		} else {
			log.Println("PDFs gerados com sucesso! Aguardando envio...")
		}

		// Etapa de envio dos certificados
		sender := mailer.Sender{
			Name:           "Prof. Dedeco",
			Address:        "dedeco@even3.com.br",
			DefaultSubject: "Certificado",
			DefaultTextContent: "Obrigado pela participação no evento. Segue certi" +
				"ficado em anexo.<br><br>(Gerado por Certifier)",
			DefaultFileName: "certificado.pdf",
		}

		for i := 0; i < len(table); i++ {
			err := sender.SendEmail(table[i]["Nome"], table[i]["Email"], pdfs[i])
			if err != nil {
				log.Println(err)
				Redirect(w, home, err.Error()+". Retorne a página anterior.")
			} else {
				log.Println("Enviado para: " + table[i]["Nome"] + " " +
					table[i]["Email"])
			}
		}

		Redirect(w, home, "Certificados enviados com sucesso! Retorne a página a"+
			"nterior.")
	})

	http.Handle("/static/", http.FileServer(http.Dir("assets/")))

	http.ListenAndServe(address, nil)
}
