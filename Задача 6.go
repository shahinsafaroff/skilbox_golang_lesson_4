//Задача 6. Студенты* (сложная задача)
package main

import "fmt"

func main() {
	fmt.Println("Число Студентиков: ")
	var N, K, studentListNumber, studentsPerGroup, studentsGroupNumber int
	fmt.Scan(&N)

	fmt.Println("Число Групп: ")
	fmt.Scan(&K)

	fmt.Println("Порядковый номер Студентика")
	fmt.Scan(&studentListNumber)
	studentsPerGroup = N / K - N % K
	studentsGroupNumber = studentListNumber - studentsPerGroup * K -

	if studentListNumber % studentsPerGroup == 0 {
		if N % K == 0 {
		}
	}
	fmt.Println("Группа Студентика", studentsGroupNumber)
} else {
fmt.Println("Группа Студентика", studentsGroupNumber + 1)
}
}
