package pdfgenerator

import (
	"io/ioutil"
	"testing"
)

func TestGenerate(t *testing.T) {
	content := []string{
		"<div align='center'><font face='Arial, Helvetica, sans-serif'><b><font " +
			"size='6'><br></font></b></font></div><div align='center'><font face='" +
			"Arial, Helvetica, sans-serif'><b><font size='6'>CERTIFICADO</font></b" +
			"><br></font></p></div><div align='center'><font size='5' face='Arial," +
			" Helvetica, sans-serif'><br></font></div><div align='center'><font si" +
			"ze='5' face='Arial, Helvetica, sans-serif'><br></font></div><div alig" +
			"n='center'><font size='5' face='Arial, Helvetica, sans-serif'>Certifi" +
			"co que Fulano participou do I Congresso Norte Americano de Letras na " +
			"cidade de Recife, no período de 10/04/2018 a 12/04/2018 totalizando u" +
			"ma carga horária de 20 horas.</font><br></div></div>",
	}

	var str string
	pdfbytes, err := Generate(content, false, str)

	if err != nil {
		t.Error(err)
	}

	err = ioutil.WriteFile("testfile.pdf", pdfbytes[0], 0700)

	if err != nil {
		t.Error(err)
	}

}
