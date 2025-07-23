package final4

import (
	"fmt"
	"net/http"
)

/*
Создайте HTTP-сервер, который обрабатывает два маршрута:

/hello — возвращает JSON-ответ {"message": "Hello, World!"}.
/about — возвращает текстовый ответ "This is the about page".
Убедитесь, что сервер запускается на порту :8000 и выводит
сообщение "Сервер запущен" перед стартом.
*/

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "{\"message\": \"Hello, World!\"}")
}
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the about page")
}

func Ex4() {

	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/about", AboutHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
	fmt.Println("Сервер запущен на :8080")
}
