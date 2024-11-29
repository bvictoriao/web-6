package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

// Переменная для счетчика
var counter int
var mu sync.Mutex // Для синхронизации доступа к счетчику

func main() {
	// Обработчик GET и POST запросов
	http.HandleFunc("/count", countHandler)

	// Запуск сервера на порту 3333
	fmt.Println("Server is running on port 3333...")
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func countHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Возвращаем текущее значение счетчика
		mu.Lock()
		fmt.Fprintf(w, "Counter: %d", counter)
		mu.Unlock()

	case "POST":
		// Читаем значение "count" из формы
		err := r.ParseForm()
		if err != nil {
			fmt.Println("Error parsing form: ", err)
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}

		if err != nil {
			fmt.Println("Invalid count value: ", err)
			http.Error(w, "это не число", http.StatusBadRequest)
			return
		}

		// Получаем значение из формы
		countStr := r.FormValue("count")

		// Конвертируем строку в число
		count, err := strconv.Atoi(countStr)
		if err != nil {
			http.Error(w, "это не число", http.StatusBadRequest)
			return
		}

		// Увеличиваем счетчик
		mu.Lock()
		counter += count
		mu.Unlock()

		// Возвращаем обновленное значение счетчика
		fmt.Fprintf(w, "Counter updated: %d", counter)

	default:
		// Если метод не поддерживается
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
