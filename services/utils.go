package services

import (
	"log"
	"os"
	"path/filepath"
)

func IsDir(path string) bool {

	/*
		folder := `C:\Temp\`

		isDir, err := services.IsDir(folder)

		if err != nil {
			fmt.Println("File not found")

		}
		fmt.Println(isDir)
		os.Exit(0)
	*/

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
