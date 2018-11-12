package services

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type Sanitizer struct {
}

func (ss *Sanitizer) Sanitize(filename string) string {

	replaceString := "_"
	s := filename

	// Remove .
	re := regexp.MustCompile(`^\.+`)
	s = re.ReplaceAllString(s, replaceString)

	// Remove trailing spaces
	re = regexp.MustCompile(`\s+|^\s+|\s+$`)
	s = re.ReplaceAllString(s, replaceString)

	// Illegal chars
	re = regexp.MustCompile(`[\/\?<>\\:\*\|"():&',!]`)
	s = re.ReplaceAllString(s, replaceString)

	// Replace _-_
	re = regexp.MustCompile(`_-_+`)
	s = re.ReplaceAllString(s, replaceString)

	// Remove trailing replace char
	re = regexp.MustCompile(`_+|^_+|_+$`)
	s = re.ReplaceAllString(s, replaceString)

	// Replace trailing .
	re = regexp.MustCompile(`\.+`)
	s = re.ReplaceAllString(s, ".")

	// Remove Accents
	s = RemoveAccents(s)

	return strings.ToUpper(s)
}

func RemoveAccents(filename string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	s, _, _ := transform.String(t, filename)

	return s
}
