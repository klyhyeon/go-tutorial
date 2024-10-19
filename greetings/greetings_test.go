package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "yuhyeon"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("yuhyeon")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(
			// %q: Formats a value as a double-quoted string.
			// If the value is a string, it will be enclosed in quotes,
			// and any special characters will be escaped (e.g., newline characters or tabs).

			// %#q: Similar to %q, but it prints a string using a Go-syntax representation.
			// It includes backticks for multiline strings or single quotes for byte slices.
			//This is useful when you want to print values in a format that can be copied and pasted directly into Go code.
			`Hello("yuhyeon") = %q, %v, want match for %#q, nil`,
			msg, err, want,
		)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	want := err == nil && msg != ""
	if !want {
		t.Fatalf(`Hello("") = %q, err = %v`, msg, err)
	}
}
