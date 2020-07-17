//Задача 7. Прогрессивный налог* (сложная задача)

package main

import "fmt"

func main() {
	fmt.Println("Enter your income: ")
	var profit, tax float64
	fmt.Scan(&profit)



	if profit > 50000 {
		tax = (profit - 50000) * 30 / 100 + ((50000 - 10000) * 20 / 100) + (10000 * 13 /100)
		fmt.Println("Proportion of tax is equal: ", tax)
	} else if profit > 10000 && profit < 50000 {
		tax = (((profit - 10000) * 20 / 100) + (10000 * 13 /100))
		fmt.Println("Proportion of tax is equal: ", tax)
	} else if profit > 0 && profit < 10000 {
		tax = profit * 13 / 100
		fmt.Println("Proportion of tax is equal: ", tax)
	} else if profit == 0 {
		fmt.Println("Really???")
	} else {
		fmt.Println("You've mistyped")
	}
}