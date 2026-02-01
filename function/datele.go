package function

import "fmt"

func DeleteTask(id int) {
	tasks := LoadTask()
	var newTasks []Task
	found := false

	for _, task := range tasks {
		if task.ID == id {
			found = true
			continue
		}
		newTasks = append(newTasks, task)
	}

	if !found {
		fmt.Println("Задача с ID: ", id, "Не найдена!")
		return
	}

	SaveTasks(newTasks)
	fmt.Println("Задача  с ID: ", id, "Удалена!")
}
