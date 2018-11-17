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

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"./services"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("File is required")
		os.Exit(1)
	}

	path := os.Args[1]

	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("File not found")
		os.Exit(1)
	}

	result := SanitizeFile(path)
	fmt.Println(result)

}

func SanitizeFile(path string) string {

	filename := filepath.Base(path)
	dir := filepath.Dir(path)

	sanitizer := services.Sanitizer{}
	sanitizedFilename := sanitizer.Sanitize(filename)

	sanitizedPath := filepath.Join(dir, sanitizedFilename)
	os.Rename(path, sanitizedPath)

	return sanitizedFilename
}
