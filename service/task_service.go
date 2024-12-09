package service

import (
	"Task-Tracker-CLI/entity"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type TaskService struct{}

func NewTaskService() TaskService {
	return TaskService{}
}

func (t *TaskService) SaveTask(data []byte, databaseName, taskDescription string) {
	task := entity.Task{
		Type:        entity.TO_DO,
		Description: taskDescription,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	newData, errCreateTask := task.Save(data)
	if errCreateTask != nil {
		log.Fatalf("error creating task, error: %v", errCreateTask)
	}

	jsonData, errMarshall := json.Marshal(newData)

	if errMarshall != nil {
		log.Fatalf("error saving database, error: %v", errMarshall)
	}

	errSaveDatabase := os.WriteFile(databaseName, jsonData, 0644)

	if errSaveDatabase != nil {
		log.Fatalf("error saving database, error: %v", errSaveDatabase)
	}

	log.Printf("Task added successfully (ID: %d)", task.ID)
}

func (t *TaskService) ChangeTaskStatus(data []byte, databaseName, taskStatus string, taskID int) {
	var currentData entity.Database
	errUnmarshallDB := json.Unmarshal(data, &currentData)

	if len(data) > 0 && errUnmarshallDB != nil {
		log.Fatalf(fmt.Sprintf("unable to get data from database, error: %v", errUnmarshallDB))
	}

	for i, task := range currentData.Task {
		if task.ID == taskID {
			currentData.Task[i].Type = taskStatus
			currentData.Task[i].UpdatedAt = time.Now()
			break
		}
	}

	jsonData, errMarshall := json.Marshal(currentData)

	if errMarshall != nil {
		log.Fatalf("error saving database, error: %v", errMarshall)
	}

	errSaveDatabase := os.WriteFile(databaseName, jsonData, 0644)

	if errSaveDatabase != nil {
		log.Fatalf("error saving database, error: %v", errSaveDatabase)
	}

	log.Printf("Update Task status successfully (ID: %d)", taskID)
}

func (t *TaskService) GetTasks(data []byte, taskStatus string) {
	var currentData entity.Database
	errUnmarshallDB := json.Unmarshal(data, &currentData)

	if len(data) > 0 && errUnmarshallDB != nil {
		log.Fatalf(fmt.Sprintf("unable to get data from database, error: %v", errUnmarshallDB))
	}

	var taskResult []entity.Task

	for _, task := range currentData.Task {
		if len(taskStatus) > 0 && task.Type != taskStatus {
			continue
		}
		taskResult = append(taskResult, task)

	}
	// If want to improve then can be improved log
	log.Printf("List %s Tasks: \n", taskStatus)

	for i, task := range taskResult {
		if i == len(taskResult)-1 {
			log.Printf("%v", task)
		} else {
			log.Printf("%v\n", task)
		}
	}

}
