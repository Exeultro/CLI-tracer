package function

import (
	"fmt"
	"time"
)

func UpdateStatus(id int, newStatus string) {
	tasks := LoadTask()
	found := false

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = newStatus
			tasks[i].UpdateAt = time.Now()
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Задача с id: ", id, "Не найдена!")
		return
	}
	SaveTasks(tasks)
	fmt.Println("Статус задачи с ID: ", id, "Изменен на: ", newStatus)
}
