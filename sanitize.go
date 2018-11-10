package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("File is required")
		os.Exit(1)
	}

	filename := os.Args[1]

	// if _, err := os.Stat(filename); os.IsNotExist(err) {
	// 	fmt.Println("File not found")
	// 	os.Exit(1)
	// }

	sanitizedFilename := Sanitize(filename)

	// os.Rename(filename, sanitizedFilename)

	fmt.Println(sanitizedFilename)
}

func Sanitize(filename string) string {

	replaceString := "_"
	s := filename

	// Remove .
	re := regexp.MustCompile(`^\.+`)
	s = re.ReplaceAllString(s, replaceString)

	// Remove trailing spaces
	re = regexp.MustCompile(`\s+|^\s+|\s+$`)
	s = re.ReplaceAllString(s, replaceString)

	// Illegal chars
	re = regexp.MustCompile(`[\/\?<>\\:\*\|"():]`)
	s = re.ReplaceAllString(s, replaceString)

	// Remove trailing replace char
	re = regexp.MustCompile(`_+|^_+|_+$`)
	s = re.ReplaceAllString(s, replaceString)

	// Remove Accents
	s = RemoveAccents(s)

	return strings.ToUpper(s)
}

func RemoveAccents(filename string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	s, _, _ := transform.String(t, filename)

	return s
}
