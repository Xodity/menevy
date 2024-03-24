package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func CreateController(controllerName, destinationPath string) error {
	fullPath := filepath.Join(destinationPath, controllerName+".php")

	err := os.MkdirAll(filepath.Dir(fullPath), os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating destination folder: %v", err)
	}

	sourcePath := "deviter/go/controller.go.php"
	content, err := ioutil.ReadFile(sourcePath)
	if err != nil {
		return fmt.Errorf("error reading source file: %v", err)
	}

	searchText := "Base"
	replaceText := controllerName

	newContent := strings.ReplaceAll(string(content), searchText, replaceText)

	file, err := os.Create(fullPath)
	if err != nil {
		return fmt.Errorf("error creating controller file: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(newContent)
	if err != nil {
		return fmt.Errorf("error writing controller file: %v", err)
	}

	fmt.Println("Controller file created successfully:", fullPath)
	return nil
}

func CreateModel(ModelName, destinationPath string) error {
	fullPath := filepath.Join(destinationPath, ModelName+".php")

	err := os.MkdirAll(filepath.Dir(fullPath), os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating destination folder: %v", err)
	}

	sourcePath := "deviter/go/model.go.php"
	content, err := ioutil.ReadFile(sourcePath)
	if err != nil {
		return fmt.Errorf("error reading source file: %v", err)
	}

	searchText := "base"
	replaceText := ModelName

	newContent := strings.ReplaceAll(string(content), searchText, replaceText)

	file, err := os.Create(fullPath)
	if err != nil {
		return fmt.Errorf("error creating controller file: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(newContent)
	if err != nil {
		return fmt.Errorf("error writing controller file: %v", err)
	}

	fmt.Println("Controller file created successfully:", fullPath)
	return nil
}
