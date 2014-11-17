package tokens

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	input := "Hello world, this is nyxtom@gmail.com http://google.com and such @nyxtom"
	results := Split(input, URLIndex, EmailIndex, EmoticonWordPunctIndex)
	expected := []string{"Hello", "world", ",", "this", "is", "nyxtom@gmail.com", "http://google.com", "and", "such", "@nyxtom"}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("Split(\"%s\") \nresulted in:\n%s, \nexpected:\n%s",
			input, fmt.Sprintf("[]string{\"%s\"}", strings.Join(results, "\",\"")),
			fmt.Sprintf("[]string{\"%s\"}", strings.Join(expected, "\",\"")))
	}
}
