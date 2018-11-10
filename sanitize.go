package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("File is required")
		os.Exit(1)
	}

	filename := os.Args[1]

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Println("File not found")
		os.Exit(1)
	}

	sanitizedFilename := sanitize(filename)
	os.Rename(filename, sanitizedFilename)

	fmt.Println(sanitizedFilename)
}

func sanitize(filename string) string {

	replaceString := "_"
	s := filename

	// Remove .
	re := regexp.MustCompile(`^\.+`)
	s = re.ReplaceAllString(s, replaceString)

	// Remove trailing spaces
	re = regexp.MustCompile(`\s+|^\s+|\s+$`)
	s = re.ReplaceAllString(s, replaceString)

	return s
}
