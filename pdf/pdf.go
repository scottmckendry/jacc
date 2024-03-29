package pdf

import (
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func GeneratePdf(url string, filePath string) (string, error) {
	pdfGenerator, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return "", err
	}

	pdfGenerator.MarginTop.Set(0)
	pdfGenerator.MarginBottom.Set(0)
	pdfGenerator.MarginLeft.Set(0)
	pdfGenerator.MarginRight.Set(0)

	page := wkhtmltopdf.NewPage(url)

	pdfGenerator.AddPage(page)

	err = pdfGenerator.Create()
	if err != nil {
		return "", err
	}

	err = pdfGenerator.WriteFile(filePath)
	if err != nil {
		return "", err
	}

	return filePath, nil
}
