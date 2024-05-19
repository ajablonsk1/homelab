package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ajablonsk1/homelab/notes/config"
	"github.com/ajablonsk1/homelab/notes/note"
)

const (
	noteName      = "TODO.md"
	templateName  = "todo.md"
	doneMarkerIdx = 3
)

var templatePath = config.Get().TemplatesPath + "/" + templateName

func main() {
	args := os.Args
	if len(args) >= 2 {
		handleArguments(args)
	} else {
		printTasks()
	}
}

func handleArguments(args []string) {
	cmd := args[1]
	switch cmd {
	case "done":
		ids := getTaskIds(cmd, args)
		markTasksDone(ids)
	case "undone":
		ids := getTaskIds(cmd, args)
		markTasksUndone(ids)
	case "update":
		if len(args) != 2 {
			log.Fatal("Wrong arguments! Usage: todo update")
		}
		updateTasks()
	default:
		addTask(args)
	}
}

func getTaskIds(cmd string, args []string) []int {
	if len(args) != 3 {
		log.Fatalf("Wrong arguments! Usage: todo %s <id>,<id>,[...]", cmd)
	}

	idsStr := strings.Split(args[2], ",")
	ids := make([]int, len(idsStr))
	for idx, idStr := range idsStr {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Fatal("Id must be integer!")
		}
		ids[idx] = id
	}
	return ids
}

func markTasksDone(ids []int) {
	markTasks(ids, "x")
}

func markTasksUndone(ids []int) {
	markTasks(ids, " ")
}

func markTasks(ids []int, marker string) {
	tasks := strings.Split(note.Read(noteName), "\n")

	for _, id := range ids {
		index := id - 1
		if index < 0 || index > len(tasks)-1 {
			log.Fatal("Please provide valid id!")
		}
		task := tasks[index]

		doneTask := task[:doneMarkerIdx] + marker + task[doneMarkerIdx+1:]
		tasks[index] = doneTask
	}

	note.Replace(noteName, strings.Join(tasks, "\n"))
}

func updateTasks() {
	tasks := strings.Split(note.Read(noteName), "\n")
	res := make([]string, 0)

	for _, task := range tasks {
		if strings.Index(task, "x") != doneMarkerIdx {
			res = append(res, task)
		}
	}

	note.Replace(noteName, strings.Join(res, "\n"))
}

func addTask(args []string) {
	message := strings.Join(args[1:], " ")
	timestamp := time.Now().Format("2006-01-02 15:02")

	data := struct {
		Data      string
		Timestamp string
	}{
		Data:      message,
		Timestamp: timestamp,
	}
	note.Append(noteName, templatePath, data)
}

func printTasks() {
	tasks := strings.Split(note.Read(noteName), "\n")
	lastIdx := len(tasks) - 1
	for idx, task := range tasks[:lastIdx] {
		fmt.Printf("%d %s\n", idx+1, task)
	}
}
