package main

import "fmt"

func sum(a, b int) int {
	rez := a + b
	return rez
}

func subtraction(a, b int) int{
	rez := a - b
	return rez
}

func multiply(a, b int) int {
	rez := a * b
	return rez
}

func del(a, b int) float32 {
	arg1 := float32(a)
	arg2 := float32(b)
	rez := arg1 / arg2
	return rez
}

//func more() string {
//	var (
//		some string
//	)
//	println("что-то ещё?")
//	fmt.Scan(&some)
//
//	return some
//}

func main() {
	var (
		a int
		b int
		c string
	)

	println("дорова, что-то будем делать, видимо\nпример: 12 + 3 интер")

	fmt.Scan(&a, &c, &b)

	switch c {
	case "+":
		println(sum(a, b))
	case "-":
		println(subtraction(a, b))
	case "*":
		println(multiply(a, b))
	case "/":
		println(del(a, b))
	default:
		println("ты что абоба? ты что ввёл?")
	}

}
