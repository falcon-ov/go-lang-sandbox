package problem15

import (
	"fmt"
	"net/http"
	"time"
)

/*
 * 15. Сетевое программирование - Задание 2
 * Создайте HTTP-сервер с двумя маршрутами.
 * /greet должен возвращать JSON с приветствием.
 * /time должен возвращать текущую дату и время.
 */
func greetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "{ greeting: \"Hello world!\"}")
}
func timeHandler(w http.ResponseWriter, r *http.Request) {
	year, month, day := time.Now().Date()
	hours, minutes, seconds := time.Now().Clock()
	result := fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", year, int(month), day, hours, minutes, seconds)
	fmt.Fprintln(w, result)
}

func StartServer2() {
	http.HandleFunc("/greet", greetHandler)
	http.HandleFunc("/time", timeHandler)

	fmt.Println("Сервер запущен на http://localhost:8080")
	// запускаем сервер на порту 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
