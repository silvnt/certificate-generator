package parser

import (
	"strings"
)

// ParseText parse to generate a list of students in a array of dict
func ParseText(txt string) []map[string]string {
	lines := strings.Split(txt, "\n")
	headers := strings.Split(lines[0], "	")
	lines = lines[1:]

	var studentsList []map[string]string

	for i := 0; i < len(lines); i++ {
		data := strings.Split(lines[i], "	")
		m := make(map[string]string, len(headers))

		for j := 0; j < len(headers); j++ {
			m[headers[j]] = data[j]
		}

		studentsList = append(studentsList, m)
	}

	return studentsList
}
