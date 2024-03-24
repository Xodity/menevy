package commands

import (
	"fmt"
	"github.com/Xodity/menevy/bin/handler"
	"os"
	"path/filepath"
	"strings"
)

func InstallHandlerNoPath(packageName string) error {
	var path string

	path = "https://github.com/" + packageName

	err := handler.RunCommand("git", "clone", path)
	if err != nil {
		return err
	}

	fmt.Println("Package ", packageName, " installed successfully.")
	return nil
}

func InstallHandlerWithPath(packageName string, pathName string) error {
	var path string

	if pathName != "" {
		path = "https://github.com/" + packageName
		err := handler.RunCommand("git", "clone", path, pathName)
		if err != nil {
			return err
		}

		fmt.Println("Package ", packageName, " installed successfully.")
		return nil
	} else {
		path = "https://github.com/" + packageName
		err := handler.RunCommand("git", "clone", path)
		if err != nil {
			return err
		}

		fmt.Println("Package ", packageName, " installed successfully.")
		return nil
	}
}

func ComposerInstallWithPath(path string) {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}
	err = os.Chdir(path)
	if err != nil {
		fmt.Println("Error changing directory:", err)
		return
	}
	err = handler.RunCommand("composer", "install")
	if err != nil {
		fmt.Println("Error running 'composer install':", err)
		return
	}
	fmt.Println("Composer install finished successfully.")
	err = os.Chdir(cwd)
	if err != nil {
		fmt.Println("Error changing directory back:", err)
	}
}

func ComposerInstall() {
	err := handler.RunCommand("composer", "install")
	if err != nil {
		fmt.Println("Error running 'composer install':", err)
		return
	}
	fmt.Println("Composer install finished successfully.")
}

func GetRepoName(packageName string) string {
	parts := strings.Split(packageName, "/")
	return parts[len(parts)-1]
}

func CleanUselessFolders(path string) {
	foldersToDelete := []string{".git", ".github"}

	for _, folder := range foldersToDelete {
		fullPath := filepath.Join(path, folder)
		err := os.RemoveAll(fullPath)
		if err != nil {
			fmt.Printf("Error deleting folder %s: %v\n", folder, err)
		} else {
			fmt.Printf("Cleaning your package %s successfully.\n", folder)
		}
	}
}
