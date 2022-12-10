package day18

import (
	"strings"
	"testing"
)

func TestReduction(t *testing.T) {
	tables := []struct {
		path     string
		expected string
	}{
		{"testinputs/day23.txt", "[[[[1,1],[2,2]],[3,3]],[4,4]]"},
		{"testinputs/2.txt", "[[[[3,0],[5,3]],[4,4]],[5,5]]"},
		{"testinputs/3.txt", "[[[[5,0],[7,4]],[5,5]],[6,6]]"},
		{"testinputs/4.txt", "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]"},
		{"testinputs/12.txt", "[[[[0,9],2],3],4]"},
		{"testinputs/13.txt", "[7,[6,[5,[7,0]]]]"},
		{"testinputs/14.txt", "[[6,[5,[7,0]]],3]"},
		{"testinputs/15.txt", "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"},
	}

	for _, table := range tables {
		reduction, _ := FishmathA(table.path)

		if strings.Compare(reduction, table.expected) != 0 {
			t.Errorf("Reduction of %s was incorrect:\nGot: %s\nWant: %s.\n", table.path, reduction, table.expected)
		}
	}
}

func TestMagnitude(t *testing.T) {
	tables := []struct {
		path     string
		expected int
	}{
		{"testinputs/5.txt", 143},
		{"testinputs/6.txt", 1384},
		{"testinputs/7.txt", 445},
		{"testinputs/8.txt", 791},
		{"testinputs/9.txt", 1137},
		{"testinputs/10.txt", 3488},
		{"testinputs/11.txt", 4140},
	}

	for _, table := range tables {
		_, magnitude := FishmathA(table.path)
		if magnitude != table.expected {
			t.Errorf("Magnitude of %s was incorrect, got: %d, want: %d.", table.path, magnitude, table.expected)
		}
	}
}

func TestAllSums(t *testing.T) {
	tables := []struct {
		path     string
		expStr 	 string
		expMag   int
	}{
		{"testinputs/16.txt", "[[[[7,8],[6,6]],[[6,0],[7,7]]],[[[7,8],[8,8]],[[7,9],[0,6]]]]", 3993},
	}

	for _, table := range tables {
		str, magnitude := FishmathB(table.path)
		if str != table.expStr {
			t.Errorf("Reduction of %s was incorrect:\nGot: %s\nWant: %s.\n", table.path, str, table.expStr)
		}
		if magnitude != table.expMag {
			t.Errorf("Magnitude of %s was incorrect, got: %d, want: %d.", table.path, magnitude, table.expMag)
		}
	}
}
