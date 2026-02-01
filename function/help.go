package function

import "fmt"

func Printhelp() {
	fmt.Println("Использование:")
	fmt.Println("go run ./main <command> [arg]")
	fmt.Println("Команды:")
	fmt.Println("add описание - добавить новую задачу")
	fmt.Println("list  - показать задачи")
	fmt.Println("update id \"описание\" - обновить описание задачи")
	fmt.Println("delete id - Удалить задачу")
	fmt.Println("mark-done id - отметить задачу как выполненную")
	fmt.Println("mark-in-progress id - отметить задачу как в процессе")
}
