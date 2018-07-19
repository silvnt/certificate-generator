package controller

import (
	"log"

	"github.com/silvnt/certifier-go/model/mailer"
)

// SendEmails send bulk emails
func SendEmails(table []map[string]string,
	senderName string, senderAddress, subject string, content string,
	filename string, pdfs [][]byte) error {
	sender := mailer.Sender{
		Name:               senderName,
		Address:            senderAddress,
		DefaultSubject:     subject,
		DefaultTextContent: content,
		DefaultFileName:    filename,
	}

	for i := 0; i < len(table); i++ {
		err := sender.SendEmail(table[i]["Nome"], table[i]["Email"], pdfs[i])
		if err != nil {
			log.Println(err)
			return err
		}

		log.Println("Enviado para: " + table[i]["Nome"] + " " + table[i]["Email"])
	}

	return nil
}
