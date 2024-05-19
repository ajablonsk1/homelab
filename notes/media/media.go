package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ajablonsk1/homelab/notes/config"
	"github.com/ajablonsk1/homelab/notes/note"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	mediaDir     = "1 - Media"
	templateName = "zettel.md"
)

var templatePath = config.Get().TemplatesPath + "/" + templateName

func checkMediaType(mediaType string) error {
	switch mediaType {
	case "Article", "Book", "Podcast", "Video":
		return nil
	default:
		return fmt.Errorf("media type must be: Article | Book | Podcast | Video")
	}
}

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Println("Wrong arguments! Usage: media <media_type> <title>")
		os.Exit(1)
	}

	caser := cases.Title(language.English)
	mediaType := caser.String(args[1])
	err := checkMediaType(mediaType)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	mediaType = mediaType + "s"
	title := strings.Join(args[2:], " ")

	timestamp := time.Now().Format("2006-01-02 15:02")
	noteName := title + ".md"
	relativeNotePath := fmt.Sprintf("%s/%s/%s", mediaDir, mediaType, noteName)

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
