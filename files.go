package GoUtilities

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

var ErrPathsCantBeCreate error = errors.New("one or more paths can not be created, check your access permits")
var ErrYAMLFileCanNotBeParsed error = errors.New("the YAML config file can not be parsed")
var ErrYAMLFileNotFound error = errors.New("the YAML file can not be found or read")

func PWD() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	return dir, err
}

func FolderExists(path string) bool {
	exists := true

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		exists = false
	}

	return exists
}

func ToFileName(name string) string {
	r := regexp.MustCompile(`[^aA-zZ0-9ñÑáÁéÉíÍóÓúÚ@._()]+`)
	result := r.ReplaceAllString(name, "-")
	result = strings.Replace(result, "\\", "-", -1)
	return result
}

func CreatePaths(paths []string) error {
	fail := false

	for _, p := range paths {
		if !FolderExists(p) {
			os.Mkdir(p, os.ModePerm)
		}
		if !FolderExists(p) {
			fail = true
		}
	}

	if fail {
		return ErrPathsCantBeCreate
	}

	return nil
}

func ReadYAML(fileName string, data interface{}) error {
	file, err := os.ReadFile(fileName)

	if err == nil {
		err = yaml.Unmarshal(file, &data)
	}

	return err
}

func WriteYAML(fileName string, d interface{}) error {

	data, err := yaml.Marshal(&d)

	if err == nil {
		err = os.WriteFile(fileName, data, 0)
		if err == nil {
			os.Chmod(fileName, 0750)
		}
	}

	return err
}
