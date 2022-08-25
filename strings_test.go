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

func TestSplitLines(t *testing.T) {
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

func TestHummanReadDuration(t *testing.T) {
	type Test struct {
		Input    int64
		Expected string
	}

	tests := []Test{
		{Input: 500, Expected: "500ms"},
		{Input: 2000, Expected: "2sec"},
		{Input: 3050, Expected: "3.05sec"},
		{Input: 120000, Expected: "2min"},
	}

	for i, test := range tests {
		result := HumanReadDuration(test.Input)
		if test.Expected != result {
			t.Errorf("Test %v return unexpected HummanReadSize string:\ngot  %s\nwant %s", i, result, test.Expected)
		}
	}
}

func TestHummanReadSize(t *testing.T) {
	type Test struct {
		Input    int64
		Expected string
	}

	tests := []Test{
		{Input: 500, Expected: "500B"},
		{Input: 2048, Expected: "2KB"},
		{Input: 3048, Expected: "2.98KB"},
		{Input: 10737418, Expected: "10.24MB"},
		{Input: 14793741824, Expected: "13.78GB"},
	}

	for i, test := range tests {
		result := HummanReadSize(test.Input)
		if test.Expected != result {
			t.Errorf("Test %v return unexpected HummanReadSize string:\ngot  %s\nwant %s", i, result, test.Expected)
		}
	}
}
