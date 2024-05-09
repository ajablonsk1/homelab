package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ajablonsk1/homelab/notes/note"
	"github.com/flosch/pongo2/v6"
)

const (
	inboxDir     = "0 - Inbox"
	templatePath = "/Users/ajablonsky/repos/homelab/notes/templates/inbox.md"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Wrong arguments! Usage: inbox <title>")
		os.Exit(1)
	}

	title := strings.Join(args[1:], " ")
	timestamp := time.Now().Format("2006-01-02 15:02")
	noteName := timestamp + " - " + title + ".md"
	relativeNotePath := inboxDir + "/" + noteName

	note.Create(relativeNotePath, templatePath, pongo2.Context{"title": title, "timestamp": timestamp})
	note.Open(relativeNotePath)
}
