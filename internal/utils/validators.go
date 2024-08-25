package utils

import "strings"

func ValidatePathExtension(path string, extensions_pattern string) bool {
	ok := false
	extensions := strings.Split(extensions_pattern, "/")

	for _, extension := range extensions {
		if ok {
			return true
		}
		ok = path[len(path)-len(extension):] == extension
	}

	return ok
}
