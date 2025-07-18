package problem15

import (
	"fmt"
	"net/http"
)

/*
 * 15. Сетевое программирование - Задание 1
 * Создайте HTTP-сервер.
 * На запрос по адресу /hello сервер должен отвечать "Hello, World!".
 */
// обработчик запроса по адресу /hello

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!") // пишем ответ в браузер
}
func StartServer() {
	http.HandleFunc("/hello", helloHandler) // регистрируем путь и функцию

	fmt.Println("Сервер запущен на http://localhost:8080")

	// запускаем сервер на порту 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
