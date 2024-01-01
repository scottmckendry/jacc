package pdf

import (
	"fmt"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func GeneratePdf(url string, filename string) (string, error) {
	pdfGenerator, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return "", err
	}

	page := wkhtmltopdf.NewPage(url)

	pdfGenerator.AddPage(page)

	err = pdfGenerator.Create()
	if err != nil {
		return "", err
	}

	invoiceFile := fmt.Sprintf("./public/%s.pdf", filename)

	err = pdfGenerator.WriteFile(invoiceFile)
	if err != nil {
		return "", err
	}

	return invoiceFile, nil
}
