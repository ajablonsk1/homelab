package note

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/flosch/pongo2/v6"
)

const (
	home             = "/Users/ajablonsky"
	notesDir         = home + "/notes"
	notesProjectPath = home + "/repos/homelab/notes"
)

func Create(relativeNotePath, templatePath, timestamp, title string) {
	template := pongo2.Must(pongo2.FromFile(templatePath))
	context := pongo2.Context{"date": timestamp, "title": title}

	out, err := template.Execute(context)
	if err != nil {
		fmt.Printf("Encountered an error while formatting template: %s", err)
		os.Exit(1)
	}

	filePath := notesDir + "/" + relativeNotePath
	err = os.WriteFile(filePath, []byte(out), 0644)
	if err != nil {
		fmt.Printf("Encountered an error while writing to a file: %s", err)
		os.Exit(1)
	}
}

func Open(relativeNotePath string) {
	scriptFile := fmt.Sprintf("%s/scripts/open_nvim.sh", notesProjectPath)
	cmd := exec.Command(scriptFile, notesDir, relativeNotePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Encountered an error while running open nvim script: %s", err)
		os.Exit(1)
	}
}
