//Задача 5. Ресторан

package main
import (
	"fmt"
)
func main() {
	var weekDay string
	var guestCount int
	var checkSumm float64

	fmt.Println("число гостей: ")
	fmt.Scan(&guestCount)
	fmt.Println("сумму чека: ")
	fmt.Scan(&checkSumm)
	fmt.Println("день недели: ")
	fmt.Scan(&weekDay)

	if guestCount > 5 {
		checkSumm = checkSumm + checkSumm * 0.10
		if weekDay == "monday" {
			checkSumm = checkSumm - checkSumm * 0.10
			fmt.Println("сумма к оплате - 10% по понедельникам скидка 10%", checkSumm)
		} else if weekDay == "friday" {
			if checkSumm >= 10000 {
				checkSumm = checkSumm - checkSumm * 0.15
				fmt.Println("сумма к оплате + 10% - 15% сумма чека превышает 10 000 руб., включается дополнительная скидка в размере 5%", checkSumm)
			}
		}

	} else if guestCount <= 5 {
		fmt.Println("сумма к оплате normal", checkSumm)
		if weekDay == "monday" {
			checkSumm = checkSumm - checkSumm * 0.10
			fmt.Println("сумма к оплате - 10% for monday", checkSumm)
		} else if weekDay == "friday" {
			 if checkSumm > 10000 {
				checkSumm = checkSumm - checkSumm * 0.15
				fmt.Println("сумма к оплате - 15% for friday", checkSumm)
			}
		}
	}
}