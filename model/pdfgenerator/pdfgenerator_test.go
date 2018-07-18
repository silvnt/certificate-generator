package pdfgenerator

import (
	"io/ioutil"
	"testing"
)

func TestGenerate(t *testing.T) {
	content := []string{
		"<div align='center'><p><br></p><p><br></p><p><br></p><p><br></p><p><fon" +
			"t face='Arial, Helvetica, sans-serif'><b><font size='6'>CERTIFICADO</" +
			"font></b><br></font></p></div><div align='center'><font size='5' face" +
			"='Arial, Helvetica, sans-serif'><br></font></div><div align='center'>" +
			"<font size='5' face='Arial, Helvetica, sans-serif'><br></font></div><" +
			"div align='center'><font size='5' face='Arial, Helvetica, sans-serif'" +
			"><br></font></div><div align='center'><font size='5' face='Arial, Hel" +
			"vetica, sans-serif'>Certifico que Fulano participou do I Congresso No" +
			"rte Americano de Letras na cidade de Recife, no período de 10/04/2018" +
			" a 12/04/2018 totalizando uma carga horária de 20 horas.</font><br></" +
			"div>",
	}

	pdfbytes, err := Generate(content)

	if err != nil {
		t.Error(err)
	}

	err = ioutil.WriteFile("testfile.pdf", pdfbytes[0], 0700)

	if err != nil {
		t.Error(err)
	}

}
