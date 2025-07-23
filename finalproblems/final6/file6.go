package final6

import (
	"fmt"
	"net/http"
)

//  Создайте HTTP-сервер с маршрутом /user, который принимает два query-параметра:
//  id и name. Сервер должен вернуть JSON-ответ вида
//  {"id": "<значение>", "name": "<значение>"}.
//  Если параметры отсутствуют, верните JSON с сообщением об ошибке:
//  {"error": "id and name are required"}. Сервер должен работать на порту :8080.

func NameHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type:", "application/json")
	id := r.URL.Query().Get("id")
	name := r.URL.Query().Get("name")
	if name == "" || id == "" {
		fmt.Fprintln(w, "Ошибка, параметры {id, name} отсутствуют")
		return
	}
	fmt.Fprintf(w, "{\"id\": \"%v\", \"name\": \"%v\"}", id, name)
}

func Ex6() {
	http.HandleFunc("/user", NameHandler)

	fmt.Println("Запуск сервера на порте :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Ошибка запуска сервера")
	}
}
