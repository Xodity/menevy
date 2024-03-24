package config

import (
	"fmt"
	"os"
)

func PrintUsage() {
	fmt.Println("Welcome to Menevy CLI")
	fmt.Println("Version Update:", version)
	fmt.Println("Current Path:", getCurrentPath())
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  - install")
	fmt.Println("  - update")
	fmt.Println("  - run-dev")
	fmt.Println("  - create-controller")
	fmt.Println("  - create-model")
}

func getCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return ""
	}
	return dir
}

const (
	version = "1.3"
)
