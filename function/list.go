package function

import "fmt"

func Tasklist(filter string) {
	tasks := LoadTask()
	if len(tasks) == 0 {
		fmt.Println("Список задач пуст")
		return
	}

	fmt.Println("Список задач")
	tasksFound := false

	for _, task := range tasks {
		if filter == "all" || task.Status == filter {
			statusIcon := "[?]"
			if task.Status == "in-progress" {
				statusIcon = "[>]"
			} else if task.Status == "done" {
				statusIcon = "[x]"
			}
			fmt.Println("ID: ", statusIcon, task.ID, task.Description)
			tasksFound = true
		}
	}
	if !tasksFound {
		fmt.Println("Задача со статусом", filter, "не найдена!")
	}
}
