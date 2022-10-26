package main

import (
	"log"
	"os"
)

func main() {
	source, err := os.Open("person.rx")
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}

	Execute(source)
}
