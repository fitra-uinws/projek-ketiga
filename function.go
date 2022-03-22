package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	//greet("hacktivate", 8)
	greet("hacktivate", "semarang")

	//function with return
	names := []string{"hacktivate", "Jalan sultan iskandar muda"}
	printMsg := greeting("hei", names)

	fmt.Println(printMsg)

	var diameter float64 = 15

	area, circumference := calculatePredifine(diameter)

	fmt.Println("Area :", area)
	fmt.Println("circumference :", circumference)

	studentlist := print("hacktiv8", "hactivkids", "kode")
	fmt.Println("studentlist :", studentlist)

	numberlist := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("total :", sum(numberlist...))

	profile("Ariel", "pasta", "Ayam Geprek", "ikan roa", "Sate")
}

func profile(name string, favfood ...string) {
	mergeFavFood := strings.Join(favfood, ",")
	fmt.Println("Hello there!!! I'm", name)
	fmt.Println("I really love to eat", mergeFavFood)
}

func sum(numbers ...int) int {
	total := 0
	for _, v := range numbers {
		total += v
	}
	return total
}
func print(names ...string) []map[string]string {
	var result []map[string]string
	for i, v := range names {
		key := fmt.Sprintf("student%d", i+1)
		temp := map[string]string{
			key: v,
		}
		result = append(result, temp)
	}
	return result
}

func calculate(d float64) (float64, float64) {
	//menghitung luas
	var area float64 = math.Pi * math.Pow(d/2, 2)

	//menghitung keliling
	var circumference float64 = math.Pi * d

	return area, circumference
}

func calculatePredifine(d float64) (area float64, circumference float64) {
	//menghitung luas
	area = math.Pi * math.Pow(d/2, 2)

	//menghitung keliling
	circumference = math.Pi * d

	return
}

func greet(name, address string) {
	// /fmt.Printf("Hello there! my name is %s and im %d year old \n", name, age)

	fmt.Println("Hello there my name is", name)
	fmt.Println("i live live in", address)
}

func greeting(msg string, names []string) string {
	joinStr := strings.Join(names, " ")
	result := fmt.Sprintf("%s %s", msg, joinStr)

	return result
}
