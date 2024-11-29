package utils

import "strings"

func LimitStackTrace(trace string, depth int) string {
	lines := strings.Split(trace, "\n")
	if len(lines) > depth {
		lines = lines[:depth]
	}

	return strings.Join(lines, "\n")
}

func InStringSlice(target string, src []string) bool {
	for _, el := range src {
		if el == target {
			return true
		}
	}

	return false
}
