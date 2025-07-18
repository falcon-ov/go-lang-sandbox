package functions

// asdasd
func Average(xs []float64) float64 {
	total := 0.0
	if len(xs) == 0 {
		return 0.0
	}
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func Add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func AddFloats(args ...float64) float64 {
	total := 0.0
	for _, v := range args {
		total += v
	}
	return total
}

func AddStrings(args ...string) string {
	total := ""
	for _, v := range args {
		total += (" " + v)
	}
	return total
}
func MakeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func Factorial(x uint) uint {
	if x == 0 {
		return 1
	}

	return x * Factorial(x-1)
}

/*
Функция sum принимает срез чисел и складывает их вместе.
Как бы выглядела сигнатура этой функции?
*/

func Sum(x []int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

/*
Напишите функцию, которая принимает число,
делит его пополам и возвращает true в случае,
если исходное число чётное, и false, если нечетное.
Например, half(1) должна вернуть (0, false),
в то время как half(2) вернет (1, true).
*/
func Half(num int) (int, bool) {
	if num%2 == 0 {
		return 1, true
	} else {
		return 0, false
	}
}

// Напишите функцию с переменным числом параметров,
// которая находит наибольшее число в списке.

func MaxNum(args ...int) int {
	maxNum := 0
	for _, v := range args {
		if maxNum < v {
			maxNum = v
		}
	}
	return maxNum
}

/*
Используя в качестве примера функцию makeEvenGenerator
напишите makeOddGenerator, генерирующую нечётные числа.
*/

func MakeOddGenerator() func() int {
	i := 1
	return func() (ret int) {
		ret = i
		i += 2
		return ret
	}
}

/*
Последовательность чисел Фибоначчи определяется как fib(0) = 0, fib(1) = 1,
 fib(n) = fib(n-1) + fib(n-2). Напишите рекурсивную функцию, находящую fib(n).
*/

func Fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}

// Что такое отложенный вызов, паника и восстановление?
// Как восстановить функцию после паники?
/*
	Отложенный вызов - это опреатор defer, который откладывает вызов
функции до конца выполнения родительской функции, при это даже будет
вызвана ошибка функция будет выполнена перед завершением работы программы.
	Паника - это ошибка вызваная на этапе выполнения программы
	Восстановление - recover() позволяет перехватить панику, и
выполниться опредленные действия, при этом recover() нужно использовать вместе с
defer, т.к panic прерывает работу программы => recover() не будет выполнен
*/

/*
Написать хороший набор тестов не всегда легко,
но даже сам процесс их написания, зачастую,
может выявить много проблем для первой реализации функции.
Например, что произойдет с нашей функцией Average,
если ей передать пустой список ([]float64{})?
Как нужно изменить функцию, чтобы она возвращала 0 в таких случаях?

Напишите серию тестов для функций Min и Max из предыдущей главы.
*/
