package day21

import (
	"testing"
)
func TestReduction(t *testing.T) {
	tables := []struct {
		path     string
		expected int
	}{
		{"testinputs/1.txt", 739785},
	}

	for _, table := range tables {
		res := PlayDeterministic(table.path)

		if res != table.expected {
			t.Errorf("Result of %s was incorrect:\nGot: %d\nWant: %d.\n", table.path, res, table.expected)
		}
	}
}