package controller

import (
	"html/template"
	"net/http"
	"os"
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

// HandleHome and execute requested operations
func HandleHome(homeTemplateName string) {

	home := template.Must(template.ParseFiles(homeTemplateName))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			home.Execute(w, nil)
			return
		}

		bgExt, bgStatus, err := GetBackground(r, "bg", "assets/static/")
		if err != nil {
			Redirect(w, home, err.Error()+". Retorne a página anterior.")
		} else if bgStatus == true {
			defer os.Remove("assets/static/bg" + bgExt)
		}

		table, tableHeaders, err := GetPeopleData(r)
		if err != nil {
			Redirect(w, home, err.Error()+". Retorne a página anterior.")
		}

		certificateTexts, err := GetCertificateTexts(r, table, tableHeaders)
		if err != nil {
			Redirect(w, home, err.Error()+". Retorne a página anterior.")
		}

		pdfBytes, err := GetPDFs(r, certificateTexts, "static/bg."+bgExt, bgStatus)
		if err != nil {
			Redirect(w, home, err.Error()+". Retorne a página anterior.")
		}

		err = SendEmails(table, "Prof. Dedeco", "dedeco@even3.com.br",
			"Certificado", "Obrigado pela participação no evento. Segue certificado"+
				" em anexo.<br><br>(Gerado por Certifier)", "certificado.pdf",
			pdfBytes)
		if err != nil {
			Redirect(w, home, err.Error()+". Retorne a página anterior.")
		}

		Redirect(w, home, "Certificados enviados com sucesso! Retorne a página an"+
			"terior.")
	})
}
