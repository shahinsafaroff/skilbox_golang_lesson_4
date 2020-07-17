//Задача 6. Студенты* (сложная задача)
package main

import "fmt"

func main() {
	fmt.Println("Число Студентиков: ")
	var N, K, studentListNumber, studentsGroupNumber int
	fmt.Scan(&N)

	fmt.Println("Число Групп: ")
	fmt.Scan(&K)

	fmt.Println("Порядковый номер Студентика")
	fmt.Scan(&studentListNumber)
	studentsPerGroup := N/K

	if N%K != 0 {
		studentsGroupNumber =  1
		fmt.Println("Группа Студентика", studentsGroupNumber)
	} else if N%K == 0 {
		studentsGroupNumber = studentListNumber / studentsPerGroup + 1
		fmt.Println("Группа Студентика", studentsGroupNumber)
		}
	}
