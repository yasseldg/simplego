package sStr

import (
	"strings"
)

// FindPatterns, if len(patterns) == 0, or found pattern, return true
func FindPatterns(path string, patterns ...string) bool {

	if len(patterns) == 0 {
		return true
	}

	for _, p := range patterns {
		if strings.Contains(path, p) {
			return true
		}
	}

	return false
}

// FindSuffix, if len(patterns) == 0, or found pattern, return true
func FindSuffix(path string, patterns ...string) bool {

	if len(patterns) == 0 {
		return true
	}

	for _, p := range patterns {
		if strings.HasSuffix(path, p) {
			return true
		}
	}

	return false
}

func SplitString(s, sep string) (res []string) {

	res = strings.Split(s, sep)
	for k, r := range res {
		res[k] = strings.TrimSpace(r)
	}
	return
}

// Abbreviation, devuelve en minusculas las primeras "n" letras de una palabra
func Abbreviation(s string, n int) string {
	return strings.ToLower(s[0:n])
}

func Lower(s string) string {
	return strings.ToLower(s)
}

func Upper(s string) string {
	return strings.ToUpper(s)
}
