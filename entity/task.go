package entity

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

const (
	DONE        = "DONE"
	IN_PROGRESS = "IN_PROGRESS"
	TO_DO       = "TO_DO"
)

type Task struct {
	ID          int
	Type        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (t *Task) Save(data []byte) (*Database, error) {
	var currentData Database
	errUnmarshallDB := json.Unmarshal(data, &currentData)

	if len(data) > 0 && errUnmarshallDB != nil {
		return nil, errors.New(fmt.Sprintf("unable to get data from database, error: %v", errUnmarshallDB))
	}

	t.ID = currentData.NextIndex + 1

	currentData.Task = append(currentData.Task, *t)

	currentData.NextIndex += 1
	currentData.TotalCount += 1

	return &currentData, nil
}
