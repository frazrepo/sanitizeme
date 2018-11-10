package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Filename is required")
		os.Exit(1)
	}

	filename := os.Args[1]
	filename = ".     -    a-ab   b-"

	fmt.Println(sanitize(filename))
}

func sanitize(filename string) string {

	replaceString := "_"
	s := filename

	// Remove .
	re := regexp.MustCompile(`^\.+`)
	s = re.ReplaceAllString(s, replaceString)

	// Remove trailing space
	re = regexp.MustCompile(`\s+|^\s+|\s+$`)
	s = re.ReplaceAllString(s, replaceString)

	return s
}
