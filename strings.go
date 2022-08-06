package GoUtilities

import "strings"

func StringTruncate(str string, length int) string {

	truncated := ""

	if length > 0 {
		count := 0
		for _, char := range str {
			truncated += string(char)
			count++
			if count >= length {
				break
			}
		}

	}
	return truncated
}

func SplitLines(str string) []string {
	lines := []string{}

	for _, l := range strings.Split(strings.ReplaceAll(str, "\r\n", "\n"), "\n") {
		if l != "" {
			lines = append(lines, l)
		}
	}
	return lines
}
