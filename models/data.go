package models

import (
	"os"
)

// Rsl represents a *.rsl file.
type Rsl struct {
	Structs []StructDef
}

// StructDef represents the rsl file.
type StructDef struct {
	Name   string  // name of the struct
	Fields []Field // fields of the struct
}

type Field struct {
	Name     string // name of the field
	Type     string // type of the field
	Required bool   // is the field required

	Tags []string
}

// NewRsl returns a struct representing the StructDef file.
func NewRsl(source *os.File) Rsl {
	lines := FileToLines(source)
	structs := LinesToStructs(lines)

	return Rsl{Structs: structs}
}

// LinesToStructs builds a slice of Struct definitions
// from the lines we filtered earlier.
func LinesToStructs(lines []Line) []StructDef {
	var list []StructDef
	var def StructDef
	var foundFirst bool

	for _, line := range lines {

		// We encountered a new struct which
		// looks something like this: [Person]
		if line.IsStartOfStruct() {
			switch foundFirst {
			case true:
				// If this is the first struct we encountered, then simply start
				// the struct definition.
				def = StructDef{Name: line.Name}
				continue
			default:
				// If this is not the first struct we encountered, then we need
				// to append the previous struct to the list and start a new one.
				foundFirst = true

				def = StructDef{Name: line.Name}
				list = append(list, def)
				continue
			}
		}

		// At this point, we have a field definition.
		def.Fields = append(def.Fields, line.Field)
	}

	list = append(list, def)
	return list
}
