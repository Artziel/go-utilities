package GoUtilities

import (
	"testing"
)

func TestStringTruncate(t *testing.T) {
	type Test struct {
		Input    string
		Length   int
		Expected string
	}

	tests := []Test{
		{Input: "La casa es grande pero más grande es la plaza!", Length: 5, Expected: "La ca"},
		{Input: "La casa es grande pero más grande es la plaza!", Length: 10, Expected: "La casa es"},
		{Input: "La casa es grande pero más grande es la plaza!", Length: 26, Expected: "La casa es grande pero más"},
	}

	for i, test := range tests {
		result := StringTruncate(test.Input, test.Length)
		if test.Expected != result {
			t.Errorf("Test %v return unexpected value:\ngot  %v\nwant %v", i, result, test.Expected)
		}
	}
}

func TestSplitLinest(t *testing.T) {
	type Test struct {
		Input    string
		Expected int
	}

	tests := []Test{
		{Input: "First Line\nSecond Line\nThird Line\nFourth  Line", Expected: 4},
		{Input: "First Line\r\nSecond Line\r\nThird Line\r\nFourth  Line", Expected: 4},
		{Input: "First Line\r\nSecond Line\nThird Line\r\nFourth  Line", Expected: 4},
	}

	for i, test := range tests {
		result := SplitLines(test.Input)
		if test.Expected != 4 {
			t.Errorf("Test %v return unexpected number of lines:\ngot  %v\nwant 4", i, result)
		}
	}
}
