package line

import "testing"

func TestColorWrapCodes(t *testing.T) {
	c := Color{Code: black}

	expected := "\x1b[30mtest\x1b[0m"
	received := c.Sprint("test")

	if expected != received {
		t.Errorf("Expected %q but received %q", expected, received)
	}
}
