package function

import (
	"fmt"
	"time"
)

const dbFile = "tasks.json"

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"` // "todo", "in-progress", "done"
	CreatedAt   time.Time `json:"CreatedAt"`
	UpdateAt    time.Time `json:"UpdateAt"`
}

func AddTask(description string) {
	tasks := LoadTask()
	maxID := 0
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	newTask := Task{
		ID:          maxID + 1,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdateAt:    time.Now(),
	}

	tasks = append(tasks, newTask)
	SaveTasks(tasks)
	fmt.Println("Задача успешно добавлена! ID:", newTask.ID)
}
