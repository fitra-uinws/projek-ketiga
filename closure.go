package main

import (
	"fmt"
	"strings"
)

func main() {
	var eventNumbers = func(number ...int) []int {
		var result []int

		for _, v := range number {
			if v%2 == 0 {
				result = append(result, v)
			}
		}
		return result

	}

	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(eventNumbers(numbers...))

	var isPalindrom = func(str string) bool {
		var temp string = ""

		for i := len(str) - 1; i >= 0; i-- {
			temp += string(byte(str[i]))
		}
		return temp == str
	}("katak")

	fmt.Println("merupakan palindrom :", isPalindrom)

	fmt.Println("========Closure AS Return Func")
	studentList := []string{"Ariell", "Nanda"}

	find := findStudent(studentList)

	fmt.Println(find("Nandas"))
	fmt.Println("========Closure AS Callback=========")
	fmt.Println("total odd numb :", findOddNumbers([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, func(numb int) bool { return numb%2 != 0 }))

}

func findOddNumbers(numbers []int, callback func(int) bool) int {
	var totalOddNumbers int
	for _, v := range numbers {
		if callback(v) {
			totalOddNumbers++
		}
	}

	return totalOddNumbers
}

func findStudent(students []string) func(string) string {
	return func(s string) string {
		var student string
		var position int

		for i, v := range students {
			if strings.ToLower(v) == strings.ToLower(s) {
				student = v
				position = i

			}
		}

		if student == "" {
			return fmt.Sprintf("%s doesn't exists", s)
		} else {
			return fmt.Sprintf("we found %s at position %d", s, position+1)
		}
	}
}
