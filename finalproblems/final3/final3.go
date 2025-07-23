package final3

import (
	"context"
	"fmt"
	"time"
)

/*
Напишите программу, которая передает через контекст значение "role"
(например, "admin") в горутину.
В горутине проверьте значение "role" и выведите сообщение: "Доступ разрешен",
 если роль "admin", или "Доступ запрещен" в противном случае.
*/

func Ex3() {
	ctx := context.WithValue(context.Background(), "role", "admin")
	go func(ctx context.Context) {
		role := ctx.Value("role").(string)
		if role == "admin" {
			fmt.Println("Доступ разрешен")
		} else {
			fmt.Println("Доступ запрещен")
		}
	}(ctx)
	time.Sleep(time.Second * 1)
}
