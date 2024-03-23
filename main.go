package main

import (
	"fmt"
	"menevy/config"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) != 2 {
		config.PrintUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "install":
		err := runCommand("composer", "install")
		if err != nil {
			fmt.Println("Error running 'composer install':", err)
			os.Exit(1)
		}
		fmt.Println("Composer install finished successfully.")
	case "update":
		err := runCommand("composer", "update")
		if err != nil {
			fmt.Println("Error running 'composer update':", err)
			os.Exit(1)
		}
		fmt.Println("Composer update finished successfully.")
	case "run-dev":
		err := runCommand("php", "-S", "localhost:4000", "-t", "public")
		if err != nil {
			fmt.Println("Error running 'php localhost server':", err)
			os.Exit(1)
		}
		fmt.Println("Composer update finished successfully.")
	default:
		config.PrintUsage()
		os.Exit(1)
	}
}

func runCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
