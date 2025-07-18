package problem12

import "os"

/*
 * 12. Работа с файлами и вводом/выводом - Задание 1
 * Напишите программу, которая создаёт файл output.txt.
 * Запишите в него строку "Hello, Go!".
 */

func CreateFile(fileName string) {
	file, err := os.Create(fileName)
	defer file.Close()
	if err != nil {
		panic(err)
	}
}
func ReWriteFile(fileName, text string) {
	byteText := []byte(text)
	err := os.WriteFile(fileName, byteText, 0644)
	if err != nil {
		panic(err)
	}
}
func ReadFile(fileName string) string {
	content, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(content)
}

func AppendtoFile(fileName, text string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.WriteString(text)
	if err != nil {
		panic(err)
	}

}
func RemoveFile(fileName string) {
	err := os.Remove(fileName)
	if err != nil {
		panic(err)
	}
}
