package mailer

import "testing"

func TestSendEmail(t *testing.T) {
	s := Sender{
		Name:               "Prof. Dedeco",
		Address:            "dedeco@even3.com.br",
		DefaultSubject:     "Congresso de Letras em Fortaleza - Certificado",
		DefaultTextContent: "Segue em anexo. Agradeço a participação.",
	}

	err := s.SendEmail("Abdellatif Bouazza", "produto+abdellatif@even3.com.br", "testfile.pdf")

	if err != nil {
		t.Error(err)
	}
}
