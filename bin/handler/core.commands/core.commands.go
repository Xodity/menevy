package core_commands

import (
	"flag"
	"fmt"
	"github.com/Xodity/menevy/bin/commands"
	"github.com/Xodity/menevy/bin/handler"
	"github.com/Xodity/menevy/config"
	"os"
)

func CoreHandler() {
	if len(os.Args) < 2 {
		config.PrintUsage()
		os.Exit(1)
	}
	flag.Parse()
	switch os.Args[1] {
	case "install":
		installCmd, withFlag, withPath := config.RouteFlag()
		installCmd.Parse(os.Args[2:])

		if *withFlag != "" {
			var path string
			if *withPath != "" {
				path = *withPath
			} else {
				repoName := commands.GetRepoName(*withFlag)
				path = "./" + repoName
			}
			err := commands.InstallHandlerWithPath(*withFlag, path)
			if err != nil {
				fmt.Println("Error installing GitHub package:", err)
				os.Exit(1)
			}
			commands.ComposerInstallWithPath(path)
			commands.CleanUselessFolders(path)
		} else {
			commands.ComposerInstall()
		}
	case "update":
		err := handler.RunCommand("composer", "update")
		if err != nil {
			fmt.Println("Error running 'composer update':", err)
			os.Exit(1)
		}
		fmt.Println("Composer update finished successfully.")
	case "run-dev":
		err := handler.RunCommand("go", "-S", "localhost:4000", "-t", "public")
		if err != nil {
			fmt.Println("Error running 'go localhost server':", err)
			os.Exit(1)
		}
		fmt.Println("Composer update finished successfully.")
	case "create-controller":
		if len(os.Args) < 3 {
			fmt.Println("Usage: menevy create-controller [controller_name]")
			os.Exit(1)
		}
		controllerName := os.Args[2]
		destinationPath := "./app/Controllers/"
		if len(os.Args) >= 4 {
			destinationPath = os.Args[3]
		}

		err := commands.CreateController(controllerName, destinationPath)
		if err != nil {
			fmt.Println("Error creating controller:", err)
			os.Exit(1)
		}
	case "create-model":
		if len(os.Args) < 3 {
			fmt.Println("Usage: menevy create-model [model_name]")
			os.Exit(1)
		}
		controllerName := os.Args[2]
		destinationPath := "./app/Models/"
		if len(os.Args) >= 4 {
			destinationPath = os.Args[3]
		}

		err := commands.CreateModel(controllerName, destinationPath)
		if err != nil {
			fmt.Println("Error creating controller:", err)
			os.Exit(1)
		}

	default:
		config.PrintUsage()
		os.Exit(1)
	}
}
