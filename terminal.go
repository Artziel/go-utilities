package GoUtilities

import "golang.org/x/term"

func GetTerminalSize() (int, int) {
	if !term.IsTerminal(0) {
		return 0, 0
	}
	width, height, err := term.GetSize(0)
	if err != nil {
		return 0, 0
	}
	return width, height
}
