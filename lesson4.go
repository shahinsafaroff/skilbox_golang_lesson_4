//package main
//
//import (
//"fmt"
//)
//
//func main() {
//	var productCost, deliveryPrice, discountAmount int
//	fmt.Scan(&productCost)
//	fmt.Scan(&deliveryPrice)
//	fmt.Scan(&discountAmount)
//
//	if productCost >= 10000 {
//		fmt.Println("Purchase is bigger than 10000 usd so you have discount on delivery in half of the actual price")
//		deliveryPrice /= 2
//		if productCost % 2 == 0 {
//			fmt.Println("Buyer gets a gift")
//		}
//	}
//
//	summary:= productCost + deliveryPrice - discountAmount
//
//
//	fmt.Println("Итого: ", summary)
//
//}

//package main

//func main() {
//	fmt.Println("Barbery Calculator")
//	fmt.Println("Enter the quantity of men in city")
//	var mansCount, barbersCount, mansPerBarber, mansPerBarberPerMonth int
//	fmt.Scan(&mansCount)
//
//	fmt.Println("How many Barbers we could hire?")
//	fmt.Scan(&barbersCount)
//
//	mansPerBarber = 8
//	mansPerBarberPerMonth = mansPerBarber * 30
//
//	requiredBarbersCount := mansCount / mansPerBarberPerMonth
//
//	if mansCount % mansPerBarberPerMonth != 0 {
//		requiredBarbersCount ++
//	}
//	fmt.Println("Needed quantity of Barbers", requiredBarbersCount)
//
//	fmt.Println(requiredBarbersCount, "Barbers can shave up to", mansPerBarberPerMonth * requiredBarbersCount, "men per month")
//
//
//
//	if requiredBarbersCount > barbersCount {
//		fmt.Println("More Barbers are needed")
//
//	} else if requiredBarbersCount == barbersCount {
//		fmt.Println("Barbers are enough for this city")
//	} else {
//		fmt.Println("Barbers are hired more than needed!!!")
//	}
//}

//package main
//import ("fmt")
//
//func main() {
//	fmt.Println("Enter your income: ")
//	var profit, tax float64
//	fmt.Scan(&profit)
//
//	if profit >= 50000 {
//		tax = profit * 30 / 100
//		fmt.Println("Proportion of tax is equal: ", tax)
//	} else if profit >= 10000 {
//		tax = profit * 20 / 100
//		fmt.Println("Proportion of tax is equal: ", tax)
//	} else if profit > 0 {
//		tax = profit * 13 / 100
//		fmt.Println("Proportion of tax is equal: ", tax)
//	} else if profit == 0 {
//		fmt.Println("Really???")
//	} else {
//		fmt.Println("You've mistyped")
//	}
//}

//package main

//func main() {
//	var first_digit, second_digit float64
//	fmt.Println("Enter your first digit: ")
//	fmt.Scan(&first_digit)
//	fmt.Println("Enter your second digit: ")
//	fmt.Scan(&second_digit)
//
//	if first_digit > second_digit {
//		fmt.Println("First Digit is greater than Second Digit")
//	} else if first_digit < second_digit {
//		fmt.Println("Second Digit is greater than First Digit")
//	} else {
//		fmt.Println("Digits are equal!!!")
//	}
//}

//package main
//import ("fmt")
//
//func main() {
//	var personnelCount, chief_salary, admin_salary, waiter_salary int
//	fmt.Println("Enter Chief salary: ")
//	fmt.Scan(&chief_salary)
//	fmt.Println("Enter Admin Salary")
//	fmt.Scan(&admin_salary)
//	fmt.Println("Enter Waiter Salary")
//	fmt.Scan(&waiter_salary)
//	fmt.Println("Personnel Count")
//	fmt.Scan(&personnelCount)
//	fmt.Println("difference between Chief and Waiter Salary", chief_salary-waiter_salary)
//	fmt.Println("Average Salary", (chief_salary+admin_salary+waiter_salary)/personnelCount)
//
//}

