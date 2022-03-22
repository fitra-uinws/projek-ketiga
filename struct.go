package main

import "fmt"

func main() {

	// var employee Employee

	// employee.name = "joni"
	// employee.age = 10
	// employee.division = "Marketing"

	// fmt.Println(employee.name)
	// fmt.Println(employee.age)
	// fmt.Println(employee.division)

	// //init struct
	// employee1 := Employee{}
	// employee1.name = "joni"
	// employee1.age = 10
	// employee1.division = "Marketing"

	// employee2 := Employee{name: "jono", age: 9, division: "Finance"}

	// fmt.Printf("employee1 %+v \n", employee1)
	// fmt.Printf("employee2 %+v \n", employee2)

	// //pointer struct
	// employee3 := Employee{name: "toni", age: 11, division: "Developer"}
	// var employee4 *Employee = &employee3

	// fmt.Println("Employee 3 :", employee3)
	// fmt.Println("Employee 4 :", employee4)
	// fmt.Println("################################")
	// employee4.name = "Ananda"
	// fmt.Println("Employee 3 :", employee3)
	// fmt.Println("Employee 4 :", employee4)

	employee1 := Employee{}

	employee1.person.name = "Airell"
	employee1.person.age = 12
	employee1.division = "Curiculum Developer"

	fmt.Printf("%+v", employee1)

	//anonym struct

	employee2 := struct {
		person   Person
		division string
	}{}

	employee2.division = "Curiculum dev"
	employee2.person.age = 15
	employee2.person.name = "Ananda"

	fmt.Printf("Employee 2 : %+v \n", employee2)
	employee3 := struct {
		person   Person
		division string
	}{
		person:   Person{name: "ani", age: 17},
		division: "Sales",
	}
	fmt.Printf("Employee 3 : %+v \n", employee3)

	//slice and strunc

	people := []Person{
		{name: "Tono", age: 10},
		{name: "Toni", age: 10},
	}

	for _, v := range people {
		fmt.Printf("%+v \n", v)
	}

}

type Person struct {
	name string
	age  int
}
type Employee struct {
	person   Person
	division string
}
