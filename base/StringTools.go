package base

import (
	"fmt"
	"strings"
	"text/template"
)

func StringCut(s string, fs string) string {
	ts := strings.Split(s, fs)
	return ts[0]
}

func StringFmtOneLine(s string) string {
	return fmt.Sprintf("%-20s", s)
}

// This function will collect template funcs of myself for template.Funcs()
func TempFunc() template.FuncMap {
	return template.FuncMap{
		"cut": StringCut,
		"fmt1line": StringFmtOneLine,
	}
}