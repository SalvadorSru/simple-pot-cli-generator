package pot

import "strings"

func isQuote(character rune) bool {
	return character == '"' || character == '\''
}

func trimPath(path string) string {
	position := strings.Index(path, "wp-content")
	if position == -1 {
		return path
	}

	return path[position+len("wp-content")+1:]
}
