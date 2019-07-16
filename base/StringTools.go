package base

import "strings"

func StringCut(s string, fs string) string {
	ts := strings.Split(s, fs)
	return ts[0]
}

