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

	studentsList, _, err := ParseText(textToParse)

	if err != nil {
		t.Error(err.Error())
	}

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

func TestParseTable(t *testing.T) {
	headers := []string{"Nome", "País", "Instituição", "Email", "Inscrição"}
	var tableToParse []map[string]string

	line := make(map[string]string, 3)
	line[headers[0]] = "Dari Araujo"
	line[headers[1]] = "Brasil"
	line[headers[2]] = "FUCAPI - FUNDAÇÃO CENTRO DE ANÁLISE, PESQUIS" +
		"A E INOVAÇÃO TECNOLÓGICA"
	line[headers[3]] = "produto+dari@even3.com.br"
	line[headers[4]] = "Pendente"
	tableToParse = append(tableToParse, line)

	line = make(map[string]string, 3)
	line[headers[0]] = "David E. Resende Almeida"
	line[headers[1]] = "Brasil"
	line[headers[2]] = "CENTRO UNIVERSITÁRIO UNA"
	line[headers[3]] = "produto+david@even3.com.br"
	line[headers[4]] = "Pendente"
	tableToParse = append(tableToParse, line)

	line = make(map[string]string, 3)
	line[headers[0]] = "Dulcilene Saraiva Reis"
	line[headers[1]] = "Brasil"
	line[headers[2]] = ""
	line[headers[3]] = "produto+dulcilene@even3.com.br"
	line[headers[4]] = "Pendente"
	tableToParse = append(tableToParse, line)

	textToParse := "Meu nome é {[Nome]}, moro no {[País]} e minha instituição " +
		"é {[Instituição]}. Meu endereço de email é {[Email]} e minha inscrição " +
		"está {[Inscrição]}."

	expectedTexts := []string{
		"Meu nome é Dari Araujo, moro no Brasil e minha instituição é FUCAPI - F" +
			"UNDAÇÃO CENTRO DE ANÁLISE, PESQUISA E INOVAÇÃO TECNOLÓGICA. Meu ender" +
			"eço de email é produto+dari@even3.com.br e minha inscrição está Pende" +
			"nte.",

		"Meu nome é David E. Resende Almeida, moro no Brasil e minha instituição" +
			" é CENTRO UNIVERSITÁRIO UNA. Meu endereço de email é produto+david@ev" +
			"en3.com.br e minha inscrição está Pendente.",

		"Meu nome é Dulcilene Saraiva Reis, moro no Brasil e minha instituição é" +
			" . Meu endereço de email é produto+dulcilene@even3.com.br e minha ins" +
			"crição está Pendente.",
	}

	certificateTexts, err := ParseTable(tableToParse, headers, textToParse)

	if err != nil {
		t.Error(err)
	}

	if len(certificateTexts) != 3 {
		t.Error("list has different size than expected")
	}

	//t.Log(certificateTexts)

	for i := 0; i < len(certificateTexts); i++ {
		if certificateTexts[i] != expectedTexts[i] {
			t.Errorf("text not expected - text nº %d", i+1)
		}
	}

}
