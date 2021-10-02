/*
Последовательность состоит из натуральных чисел и завершается числом 0. Определите количество элементов этой последовательности,
которые равны её наибольшему элементу.
*/

package main

import "fmt"

var (
	listNum   uint
	maxNum uint = 0
	count  uint
)

func main() {
	for {
		fmt.Scan(&listNum)
		if listNum >= maxNum {
			if listNum == maxNum {
				count++
			} else if listNum > maxNum {
				maxNum = listNum
				count = 1
			}
		} else if listNum == 0 {
			break
		}
	}
	fmt.Print(count)
}
