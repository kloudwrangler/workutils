package utils

import (
	"fmt"
	"os"
)

// GeneratePrefix Takes a directory, find and enumerates all the directories, and returns the next available number.
func GeneratePrefix(directory string) (string, error) {
	dirItems, err := os.ReadDir(directory)
	if err != nil {
		fmt.Println("We had problems reading the home dir")
		return "", err
	}
	// How many dirs
	count := 0
	for _, file := range dirItems {
		//fmt.Println(file.Name())
		if file.IsDir() {
			//fmt.Println("Its a directory")
			count++
		}
	}
	fmt.Println("We have ", count, " Directories")
	prefix := fmt.Sprintf("%02d", count+1)
	return prefix, nil
}
