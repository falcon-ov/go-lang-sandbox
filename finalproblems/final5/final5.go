package final5

import (
	"fmt"
	"io"
	"net/http"
)

/*
Напишите программу, которая отправляет GET-запрос к
публичному API https://jsonplaceholder.typicode.com/posts/1
и выводит полученный JSON-ответ в консоль.
Если запрос завершается с ошибкой, выведите сообщение об ошибке.
Убедитесь, что тело ответа закрывается после чтения.
*/

func Ex5() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Ошибка запроса:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения ответа:", err)
	}

	fmt.Println("Ответ сервера: ", string(body))
}
