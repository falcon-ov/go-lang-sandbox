package problem8

import "fmt"

/*
 * 8. Структуры и методы - Задание 2
 * Добавьте к структуре Person_ метод introduce().
 * Метод должен возвращать строку "Меня зовут [name], мне [age] лет".
 * Вызовите этот метод.
 */

type Person_ struct {
	Name string
	Age  int
}

func (p Person_) Introduce() {
	fmt.Printf("Меня зовут %v, мне %v лет \n", p.Name, p.Age)
}

func TestStructMethods() {
	Dan := Person_{"Dan", 20}
	Dan.Introduce()
}
