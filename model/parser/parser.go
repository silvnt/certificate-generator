package parser

import (
	"strings"
)

// Student is a type of data structure
type Student struct {
	Name               string
	Country            string
	Institution        string
	Email              string
	RegistrationStatus string
}

// ParseText parse to generate a list of students in a structure list
func ParseText(txt string) []Student {
	var list []Student
	lines := strings.Split(txt, "\n")

	for i := 0; i < len(lines); i++ {
		data := strings.Split(lines[i], "	")
		list = append(list, Student{
			Name:               data[0],
			Country:            data[1],
			Institution:        data[2],
			Email:              data[3],
			RegistrationStatus: data[4],
		})
	}

	return list
}
