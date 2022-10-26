package models

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

// Tests the file test_fixtures/simple.rcl
func Test_Simple(t *testing.T) {

	// Read file test_fixtures/simple.rcl
	source, err := os.Open("test_fixtures/simple.rcl")
	assert.NoError(t, err, "failed to open file")

	// Execute the parser
	rsl := NewRsl(source)
	structList := rsl.Structs

	// Get general expected content.
	assert.Equal(t, 2, len(structList), "Expected 2 structs")

	person := rsl.Structs[0]
	address := rsl.Structs[1]

	assert.Equal(t, "Person", person.Name, "Expected Person")
	assert.Equal(t, "Address", address.Name, "Expected Address")

}
