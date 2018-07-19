package parser

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// ParseText parse to generate a list of students in a array of dict
func ParseText(txt string) ([]map[string]string, []string, error) {
	lines := strings.Split(txt, "\r\n")
	headers := strings.Split(lines[0], "\t")
	lines = lines[1:]

	var h int
	for h = 0; h < len(headers); h++ {
		if strings.Compare(headers[h], "Nome") == 0 {
			break
		}
	}
	if h == len(headers)-1 {
		return nil, nil, errors.New("requeried field: Nome")
	}

	/*for h = 0; h < len(headers); h++ {
		if strings.Compare(headers[h], "Email") == 0 {
			break
		}
	}
	if h == len(headers)-1 {
		return nil, nil, errors.New("requeried field: Email")
	}*/

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
	textContent string) ([]string, error) {
	var certifTexts []string

	for i := 0; i < len(table); i++ {
		certif := textContent
		for j := 0; j < len(headers); j++ {
			regex := regexp.MustCompile(`\{\[` + headers[j] + `\]\}`)
			certif = regex.ReplaceAllString(certif,
				table[i][headers[j]])
		}

		if regexp.MustCompile(`\{\[.\]\}`).MatchString(certif) {
			return nil, errors.New("Há tags que não correpondem - Verifique os con" +
				"teúdos")
		}

		certifTexts = append(certifTexts, certif)
	}

	return certifTexts, nil
}
