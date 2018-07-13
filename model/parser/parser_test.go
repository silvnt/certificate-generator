package parser

import (
	"testing"
)

func TestParseText(t *testing.T) {
	textToParse := "Nome	País	Instituição	Email	Inscrição\nDari Araujo	Bras" +
		"il	FUCAPI - FUNDAÇÃO CENTRO DE ANÁLISE, PESQUISA E INOVAÇÃO TECNOLÓGICA" +
		"	produto+dari@even3.com.br	Pendente\nDavid E. Resende Almeida	Brasil	" +
		"CENTRO UNIVERSITÁRIO UNA	produto+david@even3.com.br	Pendente\nDuarcide" +
		"s Ferreira Mariosa	Brasil	UAL - Universidade Autonoma de Lisboa	produt" +
		"o+duarcides@even3.com.br	Aprovado\nDulcilene Saraiva Reis	Brasil		pr" +
		"oduto+dulcilene@even3.com.br	Pendente\nEdilane Castelo Branco	Brasil	" +
		"UNIVERSIDADE FEDERAL DO AMAZONAS	produto+edilane@even3.com.br	Aprovado" +
		"\nEdilson Carneiro	Brasil	FACULDADE DE ARACAJU	produto+edilson@even3." +
		"com.br	Pendente"

	studentsList := ParseText(textToParse)

	if len(studentsList) != 6 {
		t.Error("list has different size than expected")
	}

	if studentsList[3]["Instituição"] != "" {
		t.Error("field Institution (student nº 4) is not null - not expected")
	}

	for i := 0; i < len(studentsList); i++ {
		if studentsList[i]["Nome"] == "" {
			t.Errorf("field Nome is null - student nº %d - not expected", i+1)
		}

		if studentsList[i]["País"] == "" {
			t.Errorf("field País is null - student nº %d - not expected", i+1)
		}

		if studentsList[i]["Email"] == "" {
			t.Errorf("field Email is null - student nº %d - not expected", i+1)
		}

		if studentsList[i]["Instituição"] == "" && i != 3 {
			t.Errorf("field Instituição is null - student nº %d - not expected", i+1)
		}

		if studentsList[i]["Inscrição"] == "" {
			t.Errorf("field Inscrição is null - student nº %d - not expected", i+1)
		}
	}
}
