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
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"./services"
)

func main() {

	var dryRun bool

	flag.BoolVar(&dryRun, "dry-run", false, "Dry Run - Do not rename files")

	flag.Usage = usage

	flag.Parse()

	args := flag.Args()

	if len(args) != 1 {
		fmt.Println("Please check options")
		flag.Usage()
		os.Exit(1)
	}

	path := args[0]

	fmt.Println("Input : ", path)
	fmt.Println("Dry Run : ", dryRun)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("File not found")
		os.Exit(1)
	}

	result := SanitizeFile(path, dryRun)
	fmt.Println(result)

}

func SanitizeFile(path string, dryRun bool) string {

	filename := filepath.Base(path)
	dir := filepath.Dir(path)

	sanitizer := services.Sanitizer{}
	sanitizedFilename := sanitizer.Sanitize(filename)

	sanitizedPath := filepath.Join(dir, sanitizedFilename)

	if !dryRun {
		os.Rename(path, sanitizedPath)
	}

	return sanitizedFilename
}

var usageStr = `
                               
Usage: sanitizeme [options]
sanitizeme -f filetorename.ext

Required Options:
    -i, --input        file or direcory

Other Options:
    --dry-run          dry run, for testing
`

func usage() {
	fmt.Printf("%s\n", usageStr)
	os.Exit(0)
}
