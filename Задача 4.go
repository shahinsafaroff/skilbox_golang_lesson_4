//Задача 4. Три числа: еще попытка

package main

import "fmt"

func main() {
	var a, b, c, count int
	fmt.Println("Введите первую число: ")
	fmt.Scan(&a)
	fmt.Println("Введите вторую число: ")
	fmt.Scan(&b)
	fmt.Println("Введите третью число: ")
	fmt.Scan(&c)

	if a >= 5 {
		count++
		if b >= 5 {
			count++
		}
		if c >= 5 {
			count++
		}
		fmt.Println(" количество чисел, которые больше 5 -->", count, "числа")

	} else if b >= 5 {
		count++
		if c >= 5 {
			count++
		}
		fmt.Println(" количество чисел, которые больше 5 -->", count, "чисел")
	} else if c >= 5 {
		count++
		fmt.Println(" количество чисел, которые больше 5 -->", count, "число")
	} else if a >= 5 {
		count++
		if c >= 5 {
			count++
		}
		fmt.Println(" количество чисел, которые больше 5 -->", count, "число")
	} else {
		count = 0
		fmt.Println(" количество чисел, которые больше 5 -->", count, "число")
	}
}