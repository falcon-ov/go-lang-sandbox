package variablesProblems

import "fmt"

// Используя пример программы выше напишите программу,
// переводящую температуру из градусов Фаренгейта в градусы Цельсия.
// (C = (F - 32) * 5/9)

func FahrenToCels() {
	var input float32
	fmt.Print("Введите температуру в Фаренгетах: ")
	fmt.Scanf("%f", &input)
	celsiusValue := (input - 32) * 5 / 9
	fmt.Println("В градусах Цельсия ", celsiusValue)
}
