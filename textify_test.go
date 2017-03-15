package textify

import "testing"

const (
	pdfFilename = "testing/test.pdf"
)

func TestPDF(t *testing.T) {
	text, err := PDF(pdfFilename, "\n")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(text)
}
