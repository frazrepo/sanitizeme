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
	fmt.Println(sanitize(filename))
}

func sanitize(filename string) string {

	re := regexp.MustCompile(`ab*`)
	re.ReplaceAllString("-a-abb-", "T")

	return filename + "sanitized"
}
