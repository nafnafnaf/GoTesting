package main

import (
	"fmt"
	"os"
)

func main() {
	// Get the current working directory
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	// Read the directory entries
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	// Print the names of files and directories
	fmt.Println("Contents of", dir+":")
	for _, file := range files {
		fmt.Println("-", file.Name())
	}
}
