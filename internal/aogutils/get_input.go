package aogutils

import (
	"fmt"
	"os"
)

func GetInput() string {
	return GetFile("./input.txt")
}

func GetTest(num int) string {
	if num == 1 {
		return GetFile("./test.txt")
	}
	return GetFile(fmt.Sprintf("./test%v.txt", num))
}

func GetFile(file_path string) string {
	file, err := os.ReadFile(file_path)

	if err != nil {
		fmt.Printf("Error whilest parsing data")
		return ""
	}

	return string(file)
}
