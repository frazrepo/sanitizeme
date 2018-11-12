package main

import (
	"fmt"
	"os"
	"path/filepath"

	"./services"
)

func main() {

	sanitizer := services.Sanitizer{}

	//Debug
	// filename1 := os.Args[1]
	// fmt.Println(sanitizer.Sanitize(filename1))
	// os.Exit(0)
	// End Debug

	if len(os.Args) < 2 {
		fmt.Println("File is required")
		os.Exit(1)
	}

	path := os.Args[1]

	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("File not found")
		os.Exit(1)
	}

	filename := filepath.Base(path)
	dir := filepath.Dir(path)

	sanitizedFilename := sanitizer.Sanitize(filename)

	sanitizedPath := filepath.Join(dir, sanitizedFilename)

	os.Rename(path, sanitizedPath)

	fmt.Println(sanitizedFilename)
}
