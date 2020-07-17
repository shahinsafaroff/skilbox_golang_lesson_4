//Задача 2. Три числа

package main

import "fmt"

func main() {
	var firstDigit, secondDigit, thirdDigit int
	fmt.Println("Введите первую число: ")
	fmt.Scan(&firstDigit)
	fmt.Println("Введите вторую число: ")
	fmt.Scan(&secondDigit)
	fmt.Println("Введите третью число: ")
	fmt.Scan(&thirdDigit)

	if firstDigit > 5 {
		fmt.Println("Первуя Число болще чем 5")
	} else if secondDigit > 5 {
		fmt.Println("Вторая Число болще чем 5")
	} else if thirdDigit > 5 {
		fmt.Println("Третья Число болще чем 5")
	} else {
		fmt.Println("Ни один из цифр не больще чем 5!!!")
	}
}