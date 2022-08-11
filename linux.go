package GoUtilities

import (
	"errors"
	"os"
	"regexp"
	"runtime"
	"strings"
)

var ErrNotLinuxSystem error = errors.New("the program is not running on a linux system")

func getLinuxInfo(sys *System) error {
	var err error

	getVal := func(source string, name string) string {
		result := ""

		r := regexp.MustCompile(`(\A|\r|\n|\r\n|\s.?)` + name + `=(.*)\n`)
		match := r.FindStringSubmatch(source)
		if len(match) > 0 {
			result = strings.ReplaceAll(match[2], "\"", "")
		}

		return result
	}

	if strings.ToLower(runtime.GOOS) != "linux" {
		err = ErrNotLinuxSystem
	} else {
		if _, err := os.Stat("/lib64/ld-linux-x86-64.so.2"); err == nil {
			sys.OS.Architecture = "amd64"
		} else if _, err := os.Stat("/lib/ld-linux.so.2"); err == nil {
			sys.OS.Architecture = "i386"
		}

		content := ""
		b, e := os.ReadFile("/etc/os-release")
		if e == nil {
			content = string(b)
			sys.OS.Name = getVal(content, "PRETTY_NAME")
			sys.OS.Distro = getVal(content, "ID")
			sys.OS.Version = getVal(content, "VERSION")
			sys.OS.VersionID = getVal(content, "VERSION_ID")
			sys.OS.CodeName = getVal(content, "VERSION_CODENAME")
		}
	}
	return err
}
