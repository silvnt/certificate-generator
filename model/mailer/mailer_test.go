package mailer

import (
	"io/ioutil"
	"testing"
)

func TestSendEmail(t *testing.T) {
	s := Sender{
		Name:               "Prof. Dedeco",
		Address:            "dedeco@even3.com.br",
		DefaultSubject:     "Congresso de Letras em Fortaleza - Certificado",
		DefaultTextContent: "Segue em anexo. Agradeço a participação.",
		DefaultFileName:    "oi",
	}

	file, err := ioutil.ReadFile("testfile.pdf")

	if err != nil {
		t.Error(err)
	}

	err = s.SendEmail("Silvano", "silvano.neto@upe.br", file)

	if err != nil {
		t.Error(err)
	}
}
