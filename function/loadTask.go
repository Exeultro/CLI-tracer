package function

import (
	"encoding/json"
	"log"
	"os"
)

func LoadTask() []Task {
	data, err := os.ReadFile(dbFile)
	if os.IsNotExist(err) {
		return []Task{}
	} else if err != nil {
		log.Fatal("Критическая ошибка чтения файла:", err)
	}

	if len(data) == 0 {
		return []Task{}
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		log.Fatal("Критическая ошибка парсинга JSON:", err)
	}
	return tasks
}
