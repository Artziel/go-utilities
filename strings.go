package GoUtilities

import (
	"fmt"
	"strings"
)

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

func HummanReadSize(bytes int64) string {
	val := ""
	if bytes < 1024 {
		val = fmt.Sprintf("%dB", bytes)
	} else if bytes < 1048576 {
		val = fmt.Sprintf("%.2fKB", float64(bytes)/1048576.0)
	} else if bytes < 1073741824 {
		val = fmt.Sprintf("%.2fMB", float64(bytes)/1073741824.0)
	} else {
		val = fmt.Sprintf("%.2fGB", float64(bytes)/1073741824.0)
	}
	return val
}

func HumanReadDuration(ms int64) string {
	val := ""
	if ms < 1000 {
		val = fmt.Sprintf("%dms", ms)
	} else if ms < 60000 {
		val = fmt.Sprintf("%.2fsec", float64(ms)/1000.0)
	} else {
		val = fmt.Sprintf("%.2fmin", float64(ms)/60000.0)
	}

	return val
}
