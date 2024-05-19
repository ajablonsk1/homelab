package note

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"os/exec"

	"github.com/ajablonsk1/homelab/notes/config"
)

var (
	notesDir         = config.Get().NotesPath
	templatePath     = config.Get().TemplatesPath
	notesProjectPath = templatePath + "/.."
)

func Create(relativeNotePath, templatePath string, data interface{}) {
	out, err := formatTemplate(templatePath, data)
	if err != nil {
		log.Fatal(err)
	}

	filePath := notesDir + "/" + relativeNotePath
	err = os.WriteFile(filePath, []byte(out), 0644)
	if err != nil {
		log.Fatalf("Encountered an error while writing to a file: %s", err)
	}
}

func Open(relativeNotePath string) {
	scriptFile := fmt.Sprintf("%s/scripts/open_nvim.sh", notesProjectPath)
	cmd := exec.Command(scriptFile, notesDir, relativeNotePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Encountered an error while running open nvim script: %s", err)
	}
}

func Append(relativeNotePath, templatePath string, data interface{}) {
	out, err := formatTemplate(templatePath, data)
	if err != nil {
		log.Fatal(err)
	}

	filePath := notesDir + "/" + relativeNotePath
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Error opening a file: %s", err)
	}
	defer f.Close()

	if _, err := f.WriteString(out); err != nil {
		log.Fatalf("Error writing to a file: %s", err)
	}
}

func Read(relativeNotePath string) string {
	filePath := notesDir + "/" + relativeNotePath
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error opening a file: %s", err)
	}

	return string(data)
}

func Replace(relativeNotePath, content string) {
	filePath := notesDir + "/" + relativeNotePath
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Error opening a file: %s", err)
	}
	defer f.Close()

	if _, err := f.WriteString(content); err != nil {
		log.Fatalf("Error writing to a file: %s", err)
	}
}

func formatTemplate(templatePath string, data interface{}) (string, error) {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", fmt.Errorf("encountered an error while parsing template: %s", err)
	}

	var parsedContent bytes.Buffer
	err = tmpl.Execute(&parsedContent, data)
	if err != nil {
		return "", fmt.Errorf("encountered an error while formatting template: %s", err)
	}

	return parsedContent.String(), nil
}
