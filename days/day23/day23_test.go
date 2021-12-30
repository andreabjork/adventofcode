package day23

import (
	"testing"
)

func TestAmphipods(t *testing.T) {
	tables := []struct {
		path     string
		expected int
	}{
		{"testinputs/1.txt", 12521},
	}

	for _, table := range tables {
		best := PlayAmphipods(table.path)

		if best != table.expected {
			t.Errorf("Result of %s was incorrect:\nGot: %d\nWant: %d.\n", table.path, best, table.expected)
		}
	}
}