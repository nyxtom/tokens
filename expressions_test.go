package tokens

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
    "runtime"
)

func nameOf(f interface{}) string {
    v := reflect.ValueOf(f)
    if v.Kind() == reflect.Func {
        if rf := runtime.FuncForPC(v.Pointer()); rf != nil {
            return rf.Name()
        }
    }
    return v.String()
}

// AssertTokenizeFunc will use the given input string, tokenize function and expected output to test
func AssertTokenizeFunc(t *testing.T, input string, f func(string) []string, expected []string) {
	results := f(input)
	name := nameOf(f)
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("%s(\"%s\") \nresulted in:\n%s, \nexpected:\n%s",
			name, input,
			fmt.Sprintf("[]string{\"%s\"}", strings.Join(results, "\",\"")),
			fmt.Sprintf("[]string{\"%s\"}", strings.Join(expected, "\",\"")))
	}
}

func TestEmail(t *testing.T) {
	input := "this is nyxtom@gmail.com and g@google.com"
	expected := []string{"nyxtom@gmail.com", "g@google.com"}
	AssertTokenizeFunc(t, input, Email, expected)
}

func TestURL(t *testing.T) {
	input := "this is http://www.google.com http://salient.io nyxtom@gmail.com and g@google.com"
	expected := []string{"http://www.google.com", "http://salient.io"}
	AssertTokenizeFunc(t, input, URL, expected)
}

func TestMention(t *testing.T) {
	input := "this is http://www.google.com @nyxtom #test http://salient.io and "
	expected := []string{"@nyxtom"}
	AssertTokenizeFunc(t, input, Mention, expected)
}

func TestHashTag(t *testing.T) {
	input := "this is http://www.google.com #test http://salient.io nyxtom@gmail.com and g@google.com"
	expected := []string{"#test"}
	AssertTokenizeFunc(t, input, HashTag, expected)
}

func TestCashTag(t *testing.T) {
	input := "this is $GOOG http://www.google.com $APPL #test http://salient.io nyxtom@gmail.com and g@google.com"
	expected := []string{"$GOOG", "$APPL"}
	AssertTokenizeFunc(t, input, CashTag, expected)
}

func TestEmoticon(t *testing.T) {
	input := "this is $GOOG :) $APPL #test nyxtom@gmail.com and g@google.com"
	expected := []string{":)"}
	AssertTokenizeFunc(t, input, Emoticon, expected)
}
