package pdf_test

import (
	"testing"

	"jacc/pdf"
)

func TestGeneratePdf(t *testing.T) {
	url := "https://www.google.com"

	pdfFile, err := pdf.GeneratePdf(url, "test.pdf")
	if err != nil {
		t.Error(err)
	}

	if pdfFile == "" {
		t.Error("Expected a PDF file to be returned.")
	}
}

func TestGeneratePdf_InvalidUrl(t *testing.T) {
	url := "foo"

	pdfFile, err := pdf.GeneratePdf(url, "test.pdf")
	if err == nil {
		t.Error("Expected an error to be returned.")
	}

	if pdfFile != "" {
		t.Error("Expected no PDF file to be returned.")
	}
}
