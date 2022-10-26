package models

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNewLine(t *testing.T) {

	cases := []struct {
		Input  string
		Output Line
	}{
		{Input: "    Name:string:required|json:name,omitempty", Output: Line{
			Raw:  "Name:string:required|json:name,omitempty",
			Tags: []string{"json:name,omitempty"}}},
		{Input: "Age:int", Output: Line{
			Raw: "Age:int"}},
	}

	for _, c := range cases {
		result, err := NewLine(c.Input)
		assert.NoError(t, err)

		assert.Equal(t, c.Output.Raw, result.Raw, "stored wrong raw values. expected %v, got %v", c.Output.Raw, result.Raw)
		assert.Equal(t, c.Output.Tags, result.Tags, "stored wrong tags. expected %v, got %v", c.Output.Tags, result.Tags)
	}

}

func TestMakeBasic(t *testing.T) {

	cases := []struct {
		Whole  string
		Result Field
	}{
		{
			Whole: "Name:string:required",
			Result: Field{
				Name:     "Name",
				Type:     "string",
				Required: true,
			},
		},

		{
			Whole: "Age:int",
			Result: Field{
				Name: "Age",
				Type: "int",
			},
		},
	}

	for _, c := range cases {
		result := parseStructField(c.Whole)
		assert.True(t, reflect.DeepEqual(result, c.Result), "Expected %v, got %v", c.Result, result)
	}
}

//func TestMakeTags(t *testing.T) {
//
//	cases := []struct {
//		Parts  []string
//		Result []map[string]string
//	}{
//		{
//			Parts: []string{"json:name,omitempty", "xml:name"},
//			Result: []map[string]string{
//				{"json": "name,omitempty"},
//				{"xml": "name"},
//			}},
//	}
//
//	for _, c := range cases {
//		result := makeTags(c.Parts)
//		if len(result) != len(c.Result) {
//			t.Errorf("Expected %d, got %d", len(c.Result), len(result))
//		}
//
//		assert.True(t, reflect.DeepEqual(result, c.Result), "Expected %v, got %v", c.Result, result)
//	}
//}
