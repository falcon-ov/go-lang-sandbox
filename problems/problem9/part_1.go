package problem9

import "fmt"

/*
 * 9. Интерфейсы - Задание 1
 * Создайте интерфейс Speaker с методом Speak() string.
 * Реализуйте его для структуры Person, возвращающей приветствие.
 */

type Speaker interface {
	Speak() string
}

type Person struct {
	Name string
	Age  int
}

func (p Person) Speak() string {
	return "Привет меня зовут " + p.Name
}

func TestInterfaces() {
	var speaker Speaker
	speaker = Person{"Dan", 2}
	fmt.Println(speaker.Speak())
}
