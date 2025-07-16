package main

import "fmt"

func main() {
	// defer fmt.Println("factorial: ", functionsProblems.Factorial(10))
	// xs := []float64{1, 3, 4}
	// fmt.Printf("%.2f", functionsProblems.Average(xs))

	// fmt.Println(functionsProblems.AddStrings("Hello", "world"))
	// fmt.Println(functionsProblems.AddFloats(xs...))

	// add := func(x, y int) int {
	// 	return x + y
	// }

	// fmt.Println(add(1, 2))

	// evenGenerator := functionsProblems.MakeEvenGenerator()
	// fmt.Println(evenGenerator())
	// fmt.Println(evenGenerator())

	// defer func() {
	// 	str := recover()
	// 	fmt.Println(str)
	// }()
	// panic("PANIC")

	// fmt.Println("Sum: ", functionsProblems.Sum([]int{1, 2, 3}))

	// fmt.Println("Max val: ", functionsProblems.MaxNum(1, 9, 2, 3, 4))

	// nextOdd := functionsProblems.MakeOddGenerator()
	// fmt.Println("Odd Generator: ", nextOdd())
	// fmt.Println("Odd Generator: ", nextOdd())

	// fmt.Println("Fib num: ", functionsProblems.Fib(6))

	// x := 91
	// fmt.Println(x)
	// xPtr := new(int)
	// pointersProblems.Zero(xPtr)
	// fmt.Println(*xPtr)

	// xVar := 0
	// yVar := 5

	// pointersProblems.Swap(&xVar, &yVar)

	// fmt.Println("xVar = ", xVar)
	// fmt.Println("yVar = ", yVar)
	// c := structproblems.Circle{1, 2, 3}
	// fmt.Println(structproblems.CircleArea(c))
	// fmt.Println(c.Area())

	// android1 := structproblems.Android{structproblems.Person{"android"}, "model"}

	// fmt.Println(android1.Name)
	// fmt.Println(android1.Model)

	// circle1 := structproblems.Circle{1, 2, 3}
	// ractangle1 := structproblems.Rectangle{X1: 1, X2: 2, Y1: 3, Y2: 4}

	// fmt.Println(structproblems.TotalArea(&circle1, &ractangle1))

	// for i := 0; i < 10; i++ {
	// 	go concurrencyProblems.F(i)
	// }
	// var input string
	// fmt.Scanln(&input)

	// var c chan string = make(chan string)

	// go concurrencyProblems.Pinger(c)
	// go concurrencyProblems.Ponger(c)
	// go concurrencyProblems.Printer(c)

	// var input string
	// fmt.Scanln(&input)
	// c1 := make(chan string)
	// c2 := make(chan string)

	// go func() {
	// 	for {
	// 		c1 <- "from 1"
	// 		time.Sleep(time.Second * 1)
	// 	}
	// }()
	// go func() {
	// 	for {
	// 		c2 <- "from 2"
	// 		time.Sleep(time.Microsecond * 100)
	// 	}
	// }()
	// go func() {
	// 	for {
	// 		select {
	// 		case msg1 := <-c1:
	// 			fmt.Println("Message 1", msg1)
	// 		case msg2 := <-c2:
	// 			fmt.Println("Message 2", msg2)
	// 		case <-time.After(time.Second):
	// 			fmt.Println("timeout")
	// 		default:
	// 			fmt.Println("nothing ready")
	// 		}
	// 	}
	// }()

	// var input string
	// fmt.Scanln(&input)

	c := make(chan int, 2) // Буферизированный канал с емкостью 1

	// Отправляем значение в канал
	c <- 42 // Не блокируется, так как буфер пустой
	fmt.Println("Sent 42 to channel")

	// Пытаемся отправить еще одно значение
	c <- 100                           // Блокируется, так как буфер заполнен (емкость 1)
	fmt.Println("Sent 100 to channel") // Эта строка не выполнится

	// Получаем значение
	fmt.Println(<-c) // Получаем 42, буфер освобождается
	fmt.Println(<-c)

}
