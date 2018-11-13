package services

import (
	"os"
)

func IsDirectory(path string) (bool, error) {

	/*
		folder := `C:\Temp\`

		isDir, err := services.IsDirectory(folder)

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
		return true, nil
	} else {
		return false, nil
	}
}
