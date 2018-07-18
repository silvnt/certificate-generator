package parser

import (
	"errors"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// ParseText parse to generate a list of students in a array of dict
func ParseText(txt string) ([]map[string]string, []string, error) {
	lines := strings.Split(txt, "\n")
	headers := strings.Split(lines[0], "	")
	lines = lines[1:]

	var h int
	for h = 0; h < len(headers); h++ {
		if headers[h] == "Email" {
			break
		}
	}
	if h == len(headers)-1 {
		return nil, nil, errors.New("requeried field: Email")
	}

	var studentsList []map[string]string

	for i := 0; i < len(lines); i++ {
		data := strings.Split(lines[i], "	")
		if len(data) != len(headers) {
			if data[0] == "" || data[0] == "\n" || data[0] == "\r" {
				continue
			} else {
				return nil, nil, errors.New("linha mal preenchida - linha nº " +
					strconv.Itoa(i+2))
			}
		}

		m := make(map[string]string, len(headers))

		for j := 0; j < len(headers); j++ {
			m[headers[j]] = data[j]
		}

		studentsList = append(studentsList, m)
	}

	return studentsList, headers, nil
}

// ParseTable generates a list of certificate texts & corresponding emails
func ParseTable(table []map[string]string, headers []string,
	textContent string) ([]map[string]string, error) {
	var certifTexts []map[string]string

	for i := 0; i < len(table); i++ {
		certif := make(map[string]string, 2)
		certif["text"] = textContent
		certif["email"] = table[i]["Email"]
		for j := 0; j < len(headers); j++ {
			regex := regexp.MustCompile(`\{\[` + headers[j] + `\]\}`)
			certif["text"] = regex.ReplaceAllString(certif["text"],
				table[i][headers[j]])
		}

		if regexp.MustCompile(`\{\[.\]\}`).MatchString(certif["text"]) {
			return nil, errors.New("há tags que não correpondem - verifique os con" +
				"teúdos")
		}

		log.Println(certif)

		certifTexts = append(certifTexts, certif)
	}

	return certifTexts, nil
}
