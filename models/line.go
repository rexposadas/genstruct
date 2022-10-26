package models

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Line Type
const (
	StartOfStruct = iota
	FieldDefinition
)

// Line is a line in an rx file.
type Line struct {
	Raw  string // raw line
	Name string // StructName
	Type int    // Type of line

	Field Field // translated directly to go
	Tags  []string
}

func (l *Line) IsStartOfStruct() bool {
	return l.Type == StartOfStruct
}

// FileToLines builds a slice of lines. Each line object is
// a breakdown of a line in the file for easier processing later.
func FileToLines(source *os.File) []Line {
	s := bufio.NewScanner(source)
	var list []Line

	for s.Scan() {
		line, err := NewLine(s.Text())
		if err != nil { // noop on empty lines.
			continue
		}

		list = append(list, line)
	}

	return list
}

// NewLine return an object representing a line in an rsl file.
// Source can look like one of these:
// type 1: A struct start
//			[Person]
// type 2: Member variables
//		Name:string:required|json:name,omitempty
//
// Return an error if the line is empty.
func NewLine(source string) (Line, error) {

	source = strings.TrimSpace(source)
	// noop on empty lines.
	if source == "" {
		return Line{}, fmt.Errorf("empty line")
	}

	result := Line{Raw: source}

	// At this point, we either have (1) beginning of a struct
	// or (b) member variable.

	// Check for start of struct first.
	if IsStructStart(source) {
		s := strings.Trim(source, "[]")
		result.Name = s
		result.Type = StartOfStruct

		return result, nil
	}

	// This point we have a member variable.
	// Example: Name:string:required|json:name,omitempty
	result.Type = FieldDefinition
	result.parseField()

	return result, nil
}

func (l *Line) GetField() Field {
	return l.Field
}

func (l *Line) IsEmpty() bool {
	return strings.Trim(l.Raw, "") == ""
}

// parseField parses the member variable section.
// Example:
//  Name:string:required
func (l *Line) parseField() {

	parts := strings.Split(l.Raw, "|")

	// Step 1: parse the struct part.
	structPart := strings.TrimSpace(parts[0])
	l.Field = parseStructField(structPart)

	// Step 2: parse the tags (if present)
	if len(parts) == 1 {
		return
	}
	l.Tags = parts[1:]
}

func parseStructField(whole string) Field {
	// The struct part is <name>:<type>:<required>
	parts := strings.Split(whole, ":")

	def := Field{
		Name: parts[0],
		Type: parts[1],
	}

	if len(parts) == 3 {
		def.Required = parts[2] == "required"
	}

	return def
}

// IsStructStart returns true if the line is the start of a struct.
func IsStructStart(source string) bool {
	return strings.HasPrefix(source, "[") &&
		strings.HasSuffix(source, "]")
}

// StructName returns the name of the struct parsed from the line.
func (l *Line) StructName() string {
	return strings.Trim(l.Raw, "[]")
}
