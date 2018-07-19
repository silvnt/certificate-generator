package pdfgenerator

import (
	"strings"

	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

// Generate build PDFs in byte arrays
func Generate(content []string, bg bool, bgDir string) ([][]byte, error) {
	html := []string{
		"<html><head><meta charset='UTF-8'></head><body style='margin: 0px 0px 0" +
			"px 0px'><div style='width: 730px; height: 515px; padding: 0px 0px 0px" +
			" 0px; z-index: 0;'>",
		"<img style='position: absolute; width: 730px; height: 515px; margin: 0p" +
			"x 0px 0px 0px; padding: 0px 0px 0px 0px; z-index: -1;' src='http://" +
			bgDir + "'>",
		"<div style='padding: 75px 75px 75px 75px; height: 363px; overflow: auto" +
			"; resize: none; border: none;'>",
		"</div></body></html>",
	}

	var files [][]byte

	for i := 0; i < len(content); i++ {
		pdfg, err := wkhtmltopdf.NewPDFGenerator()
		if err != nil {
			return nil, err
		}

		pdfg.Dpi.Set(300)
		pdfg.Orientation.Set(wkhtmltopdf.OrientationLandscape)
		pdfg.Grayscale.Set(false)
		pdfg.PageSize.Set(wkhtmltopdf.PageSizeA6)
		pdfg.MarginTop.Set(0)
		pdfg.MarginBottom.Set(0)
		pdfg.MarginLeft.Set(0)
		pdfg.MarginRight.Set(0)

		if bg == true {
			pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(html[0] +
				html[1] + html[2] + content[i] + html[3])))
		} else {
			pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(html[0] +
				html[2] + content[i] + html[3])))
		}

		err = pdfg.Create()
		if err != nil {
			return nil, err
		}

		files = append(files, pdfg.Bytes())
	}

	return files, nil

}
