// Найдите первое число от 1 до n включительно, кратное c, но НЕ кратное d.

package main

import "fmt"

var (
	n, c, d int
	num int = 1
)

func main() {
	fmt.Scan(&n, &c, &d)
	for num < n+1 {
		if num % c == 0 && num % d != 0 {
			fmt.Print(num)
			break
		} else {num++}
	}
}
