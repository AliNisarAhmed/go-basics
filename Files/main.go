package main

import (
	"ali.com/go/files/fileutils"
	"fmt"
	"os"
)

func main() {
	rootPath, _ := os.Getwd()

	c, err := fileutils.ReadTextFile(rootPath + "/data/file.txt")
	if err != nil {
		panic("File not found")
	}
	newContent := fmt.Sprintf("Original: %v\nDouble: %v%v\n", c, c, c)
	fileutils.WriteToFile(rootPath+"/data/output.txt", newContent)
}
