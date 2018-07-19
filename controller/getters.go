package controller

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/silvnt/certifier-go/model/parser"
	"github.com/silvnt/certifier-go/model/pdfgenerator"
)

// GetBackground of page
func GetBackground(r *http.Request, filename string, saveOn string) (string, bool, error) {
	bg, bgHeader, err := r.FormFile("input-bg")
	if err != nil {
		log.Println(err)
		return "", false, err
	}

	if bgHeader.Filename != "" {
		bgBuf := bytes.NewBuffer(nil)
		if _, err := io.Copy(bgBuf, bg); err != nil {
			log.Println(err)
			return "", true, err
		}

		f := strings.Split(bgHeader.Filename, ".")
		fileExtension := f[len(f)-1]
		if err = ioutil.WriteFile(saveOn+filename+"."+fileExtension, bgBuf.Bytes(),
			0700); err != nil {
			log.Println(err)
			return "", true, err
		}

		return fileExtension, true, nil
	}

	return "", false, nil
}

// GetPeopleData of page
func GetPeopleData(r *http.Request) ([]map[string]string, []string, error) {
	table, tableHeaders, err := parser.ParseText(r.FormValue("input-table"))

	if err != nil {
		log.Println(err)
		return nil, nil, err
	}

	return table, tableHeaders, nil
}

// GetCertificateTexts of page and generate
func GetCertificateTexts(r *http.Request, table []map[string]string,
	tableHeaders []string) ([]string, error) {
	certifText := r.FormValue("input-editor")

	parsed, err := parser.ParseTable(table, tableHeaders, certifText)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return parsed, nil
}

// GetPDFs through of pdf generator
func GetPDFs(r *http.Request, texts []string, bgLocalization string,
	bgStatus bool) ([][]byte, error) {
	pdfs, err := pdfgenerator.Generate(texts, bgStatus, r.Host+"/"+bgLocalization)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return pdfs, nil
}
