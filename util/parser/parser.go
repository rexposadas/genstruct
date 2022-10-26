package parser

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"os"
	"strings"

	"github.com/rexposadas/genstruct/models"
	"text/template"
)

func Execute(source *os.File) {

	// Build the data that will be executed against the template.
	data := models.NewRsl(source)

	// Generate the code we will write to a file.
	code, err := generate(data)
	if err != nil {
		panic(err)
	}

	// Write generated code to a file.
	if err := writeToFile(source.Name(), code); err != nil {
		panic(err)
	}
}

func writeToFile(filename string, data []byte) error {
	parts := strings.Split(filename, ".")
	name := fmt.Sprintf("%s.go", strings.Join(parts[:len(parts)-1], ""))

	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func generate(data models.Rsl) ([]byte, error) {

	structList := data.Structs

	tmpl, err := template.ParseFiles("../templates/struct.tmpl")
	if err != nil {
		log.Fatalf("Could not parse struct template: %v\n", err)
	}
	var processed bytes.Buffer
	err = tmpl.Execute(&processed, structList)
	if err != nil {
		log.Fatalf("unable to parse data into template: %v\n", err)
	}
	formatted, err := format.Source(processed.Bytes())
	if err != nil {
		log.Fatalf("Could not format processed template: %v\n", err)
	}

	return formatted, nil
}
