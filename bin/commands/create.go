package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
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

	lastIndex := strings.LastIndex(controllerName, "/")
	if lastIndex == -1 {
		return fmt.Errorf("invalid controller name format")
	}
	replaceText := controllerName[lastIndex+1:]
	re := regexp.MustCompile(`(?i)base`)
	newContent := re.ReplaceAllString(string(content), replaceText)

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

	lastIndex := strings.LastIndex(ModelName, "/")
	if lastIndex == -1 {
		return fmt.Errorf("invalid model name format")
	}
	replaceText := ModelName[lastIndex+1:]

	newContent := strings.ReplaceAll(string(content), "Base", replaceText)

	lowercaseReplaceText := strings.ToLower(replaceText)
	newContent = strings.ReplaceAll(newContent, "base", lowercaseReplaceText+"s")

	file, err := os.Create(fullPath)
	if err != nil {
		return fmt.Errorf("error creating controller file: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(newContent)
	if err != nil {
		return fmt.Errorf("error writing controller file: %v", err)
	}

	fmt.Println("Model file created successfully:", fullPath)
	return nil
}
