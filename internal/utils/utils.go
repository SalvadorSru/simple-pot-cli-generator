package utils

import (
	"fmt"
	"os"
	"strings"
)

func ParsePath(path string) string {
	return strings.ReplaceAll(path, "\\", "/")
}

func CreateFolders(path string) error {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return fmt.Errorf("error creating folders at path %s: %v", path, err)
	}
	return nil
}

func PrintHelp() {
	fmt.Println("Usage:")
	fmt.Println("  -i, --input <path>     Path to the folder to scan (default is current directory)")
	fmt.Println("  -o, --output <file>    Name of the output POT file (default is 'to-translate.pot')")
	fmt.Println("  -h, --help             Display help information")
	fmt.Println()
	fmt.Println("Example:")
	fmt.Println("  ./po-generator -i ./src -o output.pot")
	fmt.Println("  ./po-generator")
}
