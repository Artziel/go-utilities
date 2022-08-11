package GoUtilities

import (
	"errors"
	"os/user"
	"runtime"
	"strings"
)

var ErrUnableToGetUser error = errors.New("unable to get current user")
var ErrInputValueUnexpected error = errors.New("unexpected input value")
var ErrUnsupportedSystem error = errors.New("the current OS is not supported")

type OS struct {
	Type         string `json:"Type"`
	Name         string `json:"Name"`
	Distro       string `json:"Distro"`
	Version      string `json:"Version"`
	VersionID    string `json:"VersionID"`
	CodeName     string `json:"CodeName"`
	Architecture string `json:"Architecture"`
}

type System struct {
	OS OS `json:"OS"`
}

func IsRoot() (bool, error) {
	currentUser, err := user.Current()
	if err != nil {
		return false, ErrUnableToGetUser
	}
	return currentUser.Username == "root", nil
}

func GetSystemInfo() (System, error) {
	var err error
	osType := strings.ToLower(runtime.GOOS)
	sys := System{OS: OS{Type: osType}}

	switch sys.OS.Type {
	case "linux":
		getLinuxInfo(&sys)
	default:
		err = ErrUnsupportedSystem
	}

	return sys, err
}
