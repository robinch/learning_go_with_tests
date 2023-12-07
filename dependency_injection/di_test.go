package dependency_injection

import (
	"bytes"
	"testing"
)

func TestGreeting(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Rob")

	got := buffer.String()
	want := "Hello, Rob\n"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
