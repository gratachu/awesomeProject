package greeting

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Aki")

	got := buffer.String()
	want := "Hello Aki"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
