package day21

import (
	"testing"
)

func TestDeterministic(t *testing.T) {
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

func TestDirac(t *testing.T) {
	tables := []struct {
		path     string
		expected int64
	}{
	{"testinputs/1.txt", 444356092776315},
	}

	for _, table := range tables {
		res := PlayDirac(table.path)

		if res != table.expected {
			t.Errorf("Result of %s was incorrect:\nGot: %d\nWant: %d.\n", table.path, res, table.expected)
		}
	}
}