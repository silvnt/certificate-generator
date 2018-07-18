package pdfgenerator

import (
	"strings"

	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

// Generate build PDFs in byte arrays
func Generate(content []string) ([][]byte, error) {
	html := []string{
		"<html><head><meta charset='UTF-8'></head><body style='margin: 0px 0px 0" +
			"px 0px'><div id='editor-box' style=width: 877px; height: 620px; paddi" +
			"ng: 0px 0px 0px 0px; border: 1px solid #888; z-index: 0;'><img id='ed" +
			"itor-bg' src='#' style='position: absolute; width: 877px; height: 620" +
			"px; margin: 0px 0px 0px 0px; padding: 0px 0px 0px 0px; z-index: -1; v" +
			"isibility: hidden;'><div id='editor' style='padding: 75px 75px 75px 7" +
			"5px; height: 468px; overflow: auto; resize: none; border: none;'>",
		"</div></div></body></html>",
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
		pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
		pdfg.MarginTop.Set(0)
		pdfg.MarginBottom.Set(0)
		pdfg.MarginLeft.Set(0)
		pdfg.MarginRight.Set(0)

		pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(html[0] + content[i] + html[1])))

		err = pdfg.Create()
		if err != nil {
			return nil, err
		}

		files = append(files, pdfg.Bytes())
	}

	return files, nil

}
