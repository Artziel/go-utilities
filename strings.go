package GoUtilities

import (
	"fmt"
	"math"
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
	if bytes < 1024 {
		return fmt.Sprintf("%dB", bytes)
	}
	if bytes < 1048576 {
		v := float64(bytes) / 1024.0
		if math.Mod(v, 1) != 0 {
			return fmt.Sprintf("%.2fKB", v)
		} else {
			return fmt.Sprintf("%.0fKB", v)
		}
	}
	if bytes < 1073741824 {
		v := float64(bytes) / 1048576.0
		if math.Mod(v, 1) != 0 {
			return fmt.Sprintf("%.2fMB", v)
		} else {
			return fmt.Sprintf("%.0fMB", v)
		}
	}

	v := float64(bytes) / 1073741824.0
	if math.Mod(v, 1) != 0 {
		return fmt.Sprintf("%.2fGB", v)
	} else {
		return fmt.Sprintf("%.0fGB", v)
	}
}

func HumanReadDuration(ms int64) string {
	if ms < 1000 {
		return fmt.Sprintf("%dms", ms)
	} else if ms < 60000 {
		v := float64(ms) / 1000.0
		if math.Mod(v, 1) != 0 {
			return fmt.Sprintf("%.2fsec", v)
		} else {
			return fmt.Sprintf("%.0fsec", v)
		}
	}

	v := float64(ms) / 60000.0
	if math.Mod(v, 1) != 0 {
		return fmt.Sprintf("%.2fmin", v)
	} else {
		return fmt.Sprintf("%.0fmin", v)
	}
}
