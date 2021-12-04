package main

import "fmt"

func main() {
	// Declaring Variables
	var myStr string = "Hello"
	var myInt int = 100
	var myFloat float64 = 45.12
	fmt.Println(myStr, myInt, myFloat)

	// Declaring Variable more
	var myString string = "Hello"
	var myfloat  float64 = 32.E-2
	var myInteger	int16	= 2E+2
	var myarray = [2]string{"Hello", "World"}
	var myslice []int64;
	fmt.Println(myString,myfloat,myInteger,myarray,myslice)

	//================================


	// Multiple Declarations
	var (
		employeeId int = 5
		firstName, lastName string = "Satoshi", "Nakamoto"
	)
	fmt.Println(employeeId, firstName, lastName)


	//================================


	// Short variable declaration syntax
	name := "Rajeev Singh"
	age, salary, isProgrammer := 35, 50000.0, true

	fmt.Println(name, age, salary, isProgrammer)

	// Short variable declearation synctax more
	myFname := "Kevin Nguyen"
	myage, mysalary, myIsProgrammer := 38, 3000000000, true

	fmt.Println(myFname, myage, mysalary, myIsProgrammer)
}
