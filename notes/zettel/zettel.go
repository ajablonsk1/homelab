package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ajablonsk1/homelab/notes/config"
	"github.com/ajablonsk1/homelab/notes/note"
)

const (
	zettelDir    = "4 - Zettelkasten"
	templateName = "zettel.md"
)

var templatePath = config.Get().TemplatesPath + "/" + templateName

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Wrong arguments! Usage: zettel <title>")
		os.Exit(1)
	}

	title := strings.Join(args[1:], " ")
	timestamp := time.Now().Format("2006-01-02 15:02")
	noteName := title + ".md"
	relativeNotePath := zettelDir + "/" + noteName

	data := struct {
		Title     string
		Timestamp string
	}{
		Title:     title,
		Timestamp: timestamp,
	}
	note.Create(relativeNotePath, templatePath, data)
	note.Open(relativeNotePath)
}
