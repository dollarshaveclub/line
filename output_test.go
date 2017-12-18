package line

import (
	"testing"
)

func TestInsertPrefixes(t *testing.T) {
	output := Output{prefix: "~", suffix: "<"}

	tests := map[string]struct {
		input     string
		skipFirst bool
		skipLast  bool
		expected  string
	}{
		"noskips":         {"1 2 3\n4 5 6\n\n", false, false, "~1 2 3<\n~4 5 6<\n~<\n~"},
		"skipfirst":       {"1 2 3\n4 5 6\n\n", true, false, "1 2 3<\n~4 5 6<\n~<\n~"},
		"skiplast":        {"1 2 3\n4 5 6\n\n", false, true, "~1 2 3<\n~4 5 6<\n~<\n"},
		"skipboth":        {"1 2 3\n4 5 6\n\n", true, true, "1 2 3<\n~4 5 6<\n~<\n"},
		"nonewlinenoskip": {"1 2 3", false, false, "~1 2 3"},
		"nonewlineskip":   {"1 2 3", true, true, "1 2 3"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			received := output.insertFixes(test.skipFirst, test.skipLast, test.input)
			if received != test.expected {
				t.Errorf("expected %q but received %q", test.expected, received)
			}
		})
	}
}
