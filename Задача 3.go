//Задача 3. Банкомат

package main

import "fmt"

func main() {
	var withdraw int
	fmt.Println("максимальная сумма снятия 100000 рублей")
	fmt.Println("Введите сумму снятия: ")
	fmt.Scan(&withdraw)

	if withdraw >= 100 {
		if withdraw <= 100000 {
		}
		fmt.Println("Берите свои", withdraw, "рублей и вынимайте карту!")
	} else if withdraw < 100 {
		fmt.Println("банкомат умеет выдавать только по 100 рублей и не менще 100 рублей")
	}
}