package problem16

import (
	"fmt"

	"github.com/google/uuid"
)

/*
 * 16. Основы работы с модулями - Задание 1
 * Создайте модуль.
 * Добавьте зависимость github.com/google/uuid.
 * Сгенерируйте и выведите UUID.
 */

func UuidTest() {
	id := uuid.New()
	fmt.Println("Сгенерированный UUID:", id.String())
}
