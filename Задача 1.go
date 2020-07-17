//Задача 1. Баллы ЕГЭ

package main

import "fmt"

func main() {
	var entranceScore int
	fmt.Println("Введите резултат ЕГЭ по трём экзаменам: ")
	fmt.Scan(&entranceScore)

	if entranceScore >= 275 {
		fmt.Println("Поздравляем вы поступили в Университет!")
	} else {
		fmt.Println("Беда! Пора собиратся в Армию!!!")
	}
}