package main

import (
	"cli_task_tracker/function"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		function.Printhelp()
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Ошибка: Не указано описание")
			return
		}
		function.AddTask(os.Args[2])

	case "list":
		filter := "all"
		if len(os.Args) < 2 {
			filter = os.Args[3]
		}
		function.Tasklist(filter)

	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Ошибка: Вы не указали id")
			fmt.Println("Пример: task-cli update 1 \"Новое описание\"")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		function.UpdateTask(id, os.Args[3])

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Id должен быть числом")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		function.DeleteTask(id)

	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("Не указан ID задачи")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		function.UpdateStatus(id, "done")

	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("Не указан ID задачи")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		function.UpdateStatus(id, "in-progress")

	default:
		fmt.Println("Неизвестная команда", command)
		function.Printhelp()
	}

}

func algoritm() {
	var s string
	fmt.Scan(&s)

	ans := ' '
	anscnt := 0

	for i := 0; i < len(s); i++ {
		nowcnt := 0

		for j := 0; j < len(s); j++ {
			if s[i] == s[j] {
				nowcnt++
			}
		}
		if i == 0 || nowcnt < anscnt {
			ans = rune(s[i])
			anscnt = nowcnt
		}
	}
	fmt.Printf("%c", ans)

}
