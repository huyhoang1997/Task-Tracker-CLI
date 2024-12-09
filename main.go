package main

import (
	"Task-Tracker-CLI/entity"
	"Task-Tracker-CLI/service"
	"errors"
	"log"
	"os"
	"strconv"
)

func main() {
	fileName := "database.json"
	data, err := os.ReadFile(fileName)

	if err != nil && errors.Is(os.ErrNotExist, err) {
		file, errCreateFile := os.Create(fileName)
		if errCreateFile != nil {
			log.Fatalf("error creating database file, error: %v", errCreateFile)
		}
		defer file.Close()
		data, err = os.ReadFile(fileName)
	}

	taskService := service.NewTaskService()

	switch os.Args[1] {
	case "add":
		var taskDescription string
		if len(os.Args) >= 3 {
			taskDescription = os.Args[2]
		}
		taskService.SaveTask(data, fileName, taskDescription)
		break
	case "mark-in-progress":
		if len(os.Args) < 3 {
			log.Fatalf("you should input a task id (ex: 1)")
		}
		taskID, errGetTaskID := strconv.Atoi(os.Args[2])
		if errGetTaskID != nil {
			log.Fatalf("you should input task id as number (ex: 1)")
		}
		taskService.ChangeTaskStatus(data, fileName, entity.IN_PROGRESS, taskID)
		break
	case "mark-done":
		if len(os.Args) < 3 {
			log.Fatalf("you should input a task id (ex: 1)")
		}
		taskID, errGetTaskID := strconv.Atoi(os.Args[2])
		if errGetTaskID != nil {
			log.Fatalf("you should input task id as number (ex: 1)")
		}
		taskService.ChangeTaskStatus(data, fileName, entity.DONE, taskID)
		break
	case "list":
		var taskStatus string
		if len(os.Args) >= 3 {
			taskStatus = os.Args[2]
		}
		taskService.GetTasks(data, taskStatus)
		break
	case "help":
		log.Print("available commands: \n # Adding a new task\nTask-Tracker-CLI add \"Buy groceries\"\n# Output: Task added successfully (ID: 1)\n\n# Updating and deleting tasks\nTask-Tracker-CLI update 1 \"Buy groceries and cook dinner\"\nTask-Tracker-CLI delete 1\n\n# Marking a task as in progress or done\nTask-Tracker-CLI mark-in-progress 1\nTask-Tracker-CLI mark-done 1\n\n# Listing all tasks\nTask-Tracker-CLI list\n\n# Listing tasks by status\nTask-Tracker-CLI list done\nTask-Tracker-CLI list todo\nTask-Tracker-CLI list in-progress\n")
	}
}
