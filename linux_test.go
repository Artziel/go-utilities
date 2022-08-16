package GoUtilities

import (
	"runtime"
	"strings"
	"testing"
)

func TestGetLinuxInfo(t *testing.T) {
	osType := strings.ToLower(runtime.GOOS)

	_, err := GetSystemInfo()
	switch osType {
	case "linux":
		if err != nil {
			t.Errorf("Test GetSystemInfo unexpected error:\ngot  %s\nwant No Error", err)
		}
	default:
		t.Errorf("Test GetSystemInfo unexpected system:\ngot  unknown\nwant linux")
	}

}

func TestIsRoot(t *testing.T) {
	_, err := IsRoot()

	if err != nil {
		t.Errorf("Test IsRoot unexpected error:\ngot  %s\nwant No Error", err)
	}
}
