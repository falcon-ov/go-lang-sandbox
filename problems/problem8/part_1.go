package problem8

import "fmt"

/*
 * 8. Структуры и методы - Задание 1
 * Определите структуру Person с полями name и age.
 * Создайте экземпляр и выведите его поля.
 */

type Person struct {
	Name string
	Age  int
}

func TestStructures() {
	var Jhon = Person{"John", 20}
	fmt.Println(Jhon.Name)
	fmt.Println(Jhon.Age)
}
