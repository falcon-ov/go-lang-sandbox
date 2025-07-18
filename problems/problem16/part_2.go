package problem16

import "github.com/fatih/color"

/*
 * 16. Основы работы с модулями - Задание 2
 * Напишите программу, использующую библиотеку github.com/fatih/color.
 * Выведите цветной текст в консоль.
 */

func ColorTest() {
	color.Red("Это красный текст")
	color.Green("Это зелёный текст")
	color.Blue("Это синий текст")

	// кастомный стиль (например: жёлтый жирный)
	boldYellow := color.New(color.FgYellow, color.Bold)
	boldYellow.Println("Это жёлтый жирный текст")
}
