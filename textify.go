// Package textify extracts useable text from PDF files.
// github.com/puddingfactory/pdf is a fork of rsc.io/pdf
package textify

import (
	"bytes"

	"github.com/puddingfactory/pdf"
)

// PDF reads a pdf file and returns its text
func PDF(filename string, newlineDelimiter string) (text string, err error) {
	r, err := pdf.Open(filename)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	var prev, cur pdf.Text
	for p := 1; p <= r.NumPage(); p++ {
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
