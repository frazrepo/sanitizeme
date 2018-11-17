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
	"log"
	"os"
	"path/filepath"
)

func IsDir(path string) bool {

	fi, err := os.Stat(path)

	if err != nil {
		panic(err)
	}

	if fi.IsDir() {
		return true
	} else {
		return false
	}
}

func Visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		*files = append(*files, path)
		return nil
	}
}

func ReadDir(dir string) []string {
	var files []string
	err := filepath.Walk(dir, Visit(&files))
	if err != nil {
		panic(err)
	}
	return files
}
