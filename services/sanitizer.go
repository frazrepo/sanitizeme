// Copyright (c) 2018, github.com/frazrepo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
