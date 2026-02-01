package function

import (
	"encoding/json"
	"log"
	"os"
)

func SaveTasks(tasks []Task) {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		log.Fatal("Критическая ошибка кодирования в JSON:", err)
	}
	err = os.WriteFile(dbFile, data, 0644)
	if err != nil {
		log.Fatal("Ошибка записи в файл:", err)
	}
}
