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
	filename = "     -    a-abb-"

	fmt.Println(sanitize(filename))
}

func sanitize(filename string) string {

	replaceString := "_"
	re := regexp.MustCompile(`\s+`)

	s := re.ReplaceAllString(filename, replaceString)

	return s
}
