package function

import (
	"fmt"
	"time"
)

func UpdateTask(id int, newDescription string) {
	tasks := LoadTask()
	found := false

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Description = newDescription
			tasks[i].UpdateAt = time.Now()
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Задача с ID: ", id, "Не найдена!")
		return
	}

	SaveTasks(tasks)
	fmt.Println("Задача с ID:", id, "Обнавлена!")
}
