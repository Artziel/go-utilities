package GoUtilities

import (
	"os"
	"runtime"
	"testing"
)

func TestYAMLFile(t *testing.T) {

	fileName := "./sample.yaml"

	type Data struct {
		String  string            `yaml:"string"`
		Boolean bool              `yaml:"boolean"`
		Number  int               `yaml:"number"`
		Map     map[string]string `yaml:"map,omitempty"`
		List    []string          `yaml:"list"`
	}

	d := Data{
		String:  "Sample",
		Boolean: true,
		Number:  1234,
		Map: map[string]string{
			"item_1": "Item 1",
			"item_2": "Item 2",
			"item_3": "Item 3",
		},
		List: []string{
			"Item 1",
			"Item 2",
			"Item 3",
		},
	}

	err := WriteYAML(fileName, d)
	if err != nil {
		t.Errorf("Test Write YAML fail, unexpected error: %v\n", err.Error())
	} else {
		dd := Data{}
		err := ReadYAML(fileName, dd)
		if err != nil {
			t.Errorf("Test Read YAML File fail, unexpected error: %v\n", err.Error())
		}
	}

	os.Remove(fileName)
}

func TestPWD(t *testing.T) {
	dir, err := PWD()
	if err != nil {
		t.Errorf("Test PWD Function return unexpected error: %v\n", err.Error())
	}
	if dir == "" {
		t.Errorf("Test PWD Function return unexpected value: Empty String\n")
	}
}

func TestToFileName(t *testing.T) {
	type test struct {
		Input    string
		Expected string
	}

	names := []test{
		{Input: "no way miguel.yaml rrr", Expected: "no-way-miguel.yaml-rrr"},
		{Input: "años de data en un excel.xlsx", Expected: "años-de-data-en-un-excel.xlsx"},
		{Input: "una canción de amor.doc", Expected: "una-canción-de-amor.doc"},
		{Input: "correo de sample@sample.com.mail", Expected: "correo-de-sample@sample.com.mail"},
		{Input: `c@o#n·\ c&a%r/a(c)t=?/eres[ ]especiales`, Expected: "c@o-n---c-a-r-a(c)t-eres[-]especiales"},
	}

	for i, test := range names {
		result := ToFileName(test.Input)

		if result != test.Expected {
			t.Errorf("Test %v return unexpected result file name: got %v want %v", i, result, test.Expected)
		}
	}
}

func TestFolderExists(t *testing.T) {
	type test struct {
		Input    string
		Expected bool
	}

	dir, err := PWD()
	if err != nil {
		t.Fatal(err)
	}

	tests := []test{
		{Input: dir, Expected: true},
		{Input: "./not-exists", Expected: false},
	}

	for i, test := range tests {
		result := FolderExists(test.Input)
		if test.Expected != result {
			t.Errorf("Test %v return unexpected result: got %v want %v", i, result, test.Expected)
		}
	}
}

func TestCreatePaths(t *testing.T) {
	type test struct {
		Input    string
		Expected error
	}

	isWindows := false
	if runtime.GOOS == "windows" {
		isWindows = true
	}

	dir, err := PWD()
	if err != nil {
		t.Fatal(err)
	}

	var tests []test
	if isWindows {
		tests = []test{
			{Input: dir + "\\test-folder", Expected: nil},
			{Input: "zr:\\test-folder", Expected: ErrPathsCantBeCreate},
		}
	} else {
		tests = []test{
			{Input: dir + "/test-folder", Expected: nil},
			{Input: "/test-folder", Expected: ErrPathsCantBeCreate},
		}
	}

	for i, test := range tests {
		result := CreatePaths([]string{test.Input})
		if test.Expected == nil && result != nil {
			t.Errorf("Test %v return unexpected result: got ErrPathsCantBeCreate want No Error", i)
		}
		if test.Expected != nil && result == nil {
			t.Errorf("Test %v return unexpected result: got No Error want ErrPathsCantBeCreate", i)
		}
	}
}
