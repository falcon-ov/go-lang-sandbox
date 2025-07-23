package final7

import (
	"fmt"
	"net/http"
)

// Создайте HTTP-сервер, который обрабатывает три маршрута:

// / — возвращает текстовый ответ "Welcome to the Home Page!".
// /info — возвращает JSON-ответ {"version": "1.0", "status": "running"}.
// /contact — возвращает текстовый ответ "Contact us at support@example.com".
// Сервер должен запускаться на порту :9000 и выводить в консоль сообщение
// "Сервер запущен на порту 9000" перед стартом.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the Home Page!")
}
func InfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	fmt.Fprint(w, "{\"version\": \"1.0\", \"status\": \"running\"}")
}
func ContactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Contact us at support@example.com")
}

func Ex7() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/contact", ContactHandler)
	http.HandleFunc("/info", InfoHandler)

	if err := http.ListenAndServe(":9000", nil); err != nil {
		fmt.Println("Произошла ошибка")
	}
}
