package main_test

import (
	"testing"
)

func TestIsTextMatch(t *testing.T) {
	t.Parallel()
	tests := []struct {
		subject  string
		argument string
		result   bool
	}{
		{
			subject:  "matched argument",
			argument: "https://test/queries",
			result:   true,
		},
		{
			subject:  "does not match argument",
			argument: "https://www.google.co.jp",
			result:   false,
		},
	}

	for _, te := range tests {
		t.Run(te.subject, func(t *testing.T) {
			if isTextMatch(te.argument) != te.result {
				t.Errorf("returns value is invalid[argument: %v, result: %v]\n", te.argument, te.result)
			}
		})
	}
}
