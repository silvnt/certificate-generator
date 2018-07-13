package parser

import (
	"testing"
)

func TestParseText(t *testing.T) {
	textToParse := "Dari Araujo	Brasil	FUCAPI - FUNDAÇÃO CENTRO DE ANÁLISE, P" +
		"ESQUISA E INOVAÇÃO TECNOLÓGICA	produto+dari@even3.com.br	Pendente\nDavi" +
		"d E. Resende Almeida	Brasil	CENTRO UNIVERSITÁRIO UNA	produto+david@ev" +
		"en3.com.br	Pendente\nDuarcides Ferreira Mariosa	Brasil	UAL - Universi" +
		"dade Autonoma de Lisboa	produto+duarcides@even3.com.br	Aprovado\nDulc" +
		"ilene Saraiva Reis	Brasil		produto+dulcilene@even3.com.br	Pendente\n" +
		"Edilane Castelo Branco	Brasil	UNIVERSIDADE FEDERAL DO AMAZONAS	produt" +
		"o+edilane@even3.com.br	Aprovado\nEdilson Carneiro	Brasil	FACULDADE DE" +
		" ARACAJU	produto+edilson@even3.com.br	Pendente"

	studentsList := ParseText(textToParse)

	if len(studentsList) != 6 {
		t.Error("list has different size than expected")
	}

	if studentsList[3].Institution != "" {
		t.Error("field Institution (student nº 4) is not null - not expected")
	}

	for i := 0; i < len(studentsList); i++ {
		if studentsList[i].Name == "" {
			t.Errorf("field Name is null - student nº %d - not expected", i+1)
		}

		if studentsList[i].Country == "" {
			t.Errorf("field Country is null - student nº %d - not expected", i+1)
		}

		if studentsList[i].Email == "" {
			t.Errorf("field Email is null - student nº %d - not expected", i+1)
		}

		if studentsList[i].Institution == "" && i != 3 {
			t.Errorf("field Institution is null - student nº %d - not expected", i+1)
		}

		if studentsList[i].RegistrationStatus == "" {
			t.Errorf("field RegistrationStatus is null - student nº %d - not expect"+
				"ed", i+1)
		}
	}
}
