package cli

import (
	"fmt"
	"strconv"
	"task-cli/internal/storage"
	"task-cli/internal/tasks"
)

func Run(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: tracker-cli help")
		return
	}
	command := args[1]
	fileStorage := storage.FileStorage{Name: "C:\\Users\\user\\bin\\tracker-cli\\tasks.json"}
	tracker := tasks.NewTracker(&fileStorage)
	switch command {
	case "help":
		fmt.Println("add: tracker-cli add <task>. It adds your task for tracking")
		fmt.Println("update: tracker-cli update <task_id> <new_task>. It update your task by ID")
		fmt.Println("delete: tracker-cli delete <task_id>. It deletes your task by ID")
		fmt.Println("mark-in-progress: tracker-cli mark-in-progress <task_id>. It marks your task in-progress by ID")
		fmt.Println("mark-done: tracker-cli mark-done <task_id>. It marks your task done by ID")
		fmt.Println("list: tracker-cli list. It print your task list with them id")
		fmt.Println("list: tracker-cli list <done/todo/in-progress>. It print your task list with them id by category")

	case "add":
		if len(args) < 3 {
			fmt.Println("Usage: tracker-cli add <task>")
		}
		err := tracker.Add(args[2])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Task: ", args[2], "added")

	case "update":
		if len(args) < 4 {
			fmt.Println("Usage: tracker-cli update <task_id> <new_task>")
		}
		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("You should provide a valid task id")
		}
		err = tracker.Update(id, args[3])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Task: ", args[2], "updated")

	case "delete":
		if len(args) < 3 {
			fmt.Println("Usage: tracker-cli delete <task_id>")
		}
		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("You should provide a valid task id")
		}
		err = tracker.Delete(id)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Task ", args[2], "deleted")

	case "mark-in-progress":
		if len(args) < 3 {
			fmt.Println("Usage: tracker-cli mark-in-progress <task_id>")
		}
		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("You should provide a valid task id")
		}
		err = tracker.Update(id, "in-progress")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Task ", args[2], "marked in-progress")

	case "mark-done":
		if len(args) < 3 {
			fmt.Println("Usage: tracker-cli mark-done <task_id>")
		}
		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("You should provide a valid task id")
		}
		err = tracker.Update(id, "done")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Task ", args[2], "marked done")

	case "list":
		var result []tasks.Task
		var err error
		if len(args) < 3 {
			result, err = tracker.List("")
		} else if args[2] == "in-progress" || args[2] == "done" {
			result, err = tracker.List(args[2])
		} else {
			fmt.Println("You should provide a valid category")
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		for _, task := range result {
			fmt.Printf("%v. %v - %v, created at: %v, updated at: %v \n", task.ID, task.Description, task.Status,
				task.CreatedAt, task.UpdatedAt)
		}

	default:
		fmt.Println("Unknown command: ", command)
	}
}
