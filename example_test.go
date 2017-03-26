package textify_test

import (
	"log"

	"github.com/puddingfactory/textify"
)

func ExamplePDF() {
	text, err := textify.PDF("testing/test.pdf", "\n")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(text)
}
