package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ajablonsk1/homelab/notes/note"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	materialDir  = "1 - Source Material"
	templatePath = "/Users/ajablonsky/repos/homelab/notes/templates/zettel.md"
)

func checkMaterialType(materialType string) error {
	switch materialType {
	case "Article", "Book", "Podcast", "Video":
		return nil
	default:
		return fmt.Errorf("Material type must be: Article | Book | Podcast | Video")
	}
}

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Println("Wrong arguments! Usage: material <material_type> <title>")
		os.Exit(1)
	}

	caser := cases.Title(language.English)
	materialType := caser.String(args[1])
	err := checkMaterialType(materialType)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	materialType = materialType + "s"
	title := strings.Join(args[2:], " ")

	timestamp := time.Now().Format("2006-01-02 15:02")
	noteName := title + ".md"
	relativeNotePath := fmt.Sprintf("%s/%s/%s", materialDir, materialType, noteName)

	note.Create(relativeNotePath, templatePath, timestamp, title)
	note.Open(relativeNotePath)
}
