package main

import "fmt"

func main() {
	var firstNumber int = 4
	var secondNumber *int = &firstNumber

	fmt.Println("First number (value) :", firstNumber)
	fmt.Println("First number (memory address) :", &firstNumber)

	fmt.Println("Secound number (value) :", *secondNumber)
	fmt.Println("Secound number (memory address) :", secondNumber)

	//ganti value
	var firstPerson string = "Jhon"
	var secondPerson *string = &firstPerson

	fmt.Println("First person (value) :", firstPerson)
	fmt.Println("First person (memory address) :", &firstPerson)

	fmt.Println("Secound person (value) :", *secondPerson)
	fmt.Println("Secound person (memory address) :", secondPerson)

	*secondPerson = "doe"
	fmt.Println("First person (value) :", firstPerson)
	fmt.Println("First person (memory address) :", &firstPerson)

	fmt.Println("Secound person (value) :", *secondPerson)
	fmt.Println("Secound person (memory address) :", secondPerson)

	//sebagai paramter
	a := 1
	changeValue(&a)
	fmt.Println("Hasil :", a)
}

func changeValue(number *int) {
	*number = 123
}
