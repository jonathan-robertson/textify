package textify_test

import (
	"testing"

	"github.com/puddingfactory/textify"
)

const (
	pdfFilename = "testing/test.pdf"
)

func TestPDF(t *testing.T) {
	text, err := textify.PDF(pdfFilename, "\n")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(text)
}
