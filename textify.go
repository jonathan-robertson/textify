// Package textify extracts useable text from PDF files.
package textify

import (
	"bytes"

	"github.com/puddingfactory/pdf" // fork of rsc.io/pdf
)

// PDF reads a pdf file and returns its text
func PDF(filename string, newlineDelimiter string) (text string, err error) {
	r, err := pdf.Open(filename)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	for p := 1; p <= r.NumPage(); p++ {
		var prev, cur pdf.Text
		for _, letter := range r.Page(p).Content().Text {
			prev = cur
			cur = letter

			if prev.S != "" && prev.Y != cur.Y {
				buf.WriteString(newlineDelimiter)
			}
			buf.WriteString(letter.S)
		}
	}

	return buf.String(), nil
}
