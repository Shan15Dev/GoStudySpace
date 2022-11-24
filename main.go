package main

import (
	"fmt"
)

var a = 2

func main() {
	if a == 3 {
		fmt.Println(a)
	} else {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}

	switch a {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")

	}

	array := [5]int{4, 5, 6, 7, 8}

	// Itterate through list with giving the current position
	for i, s := range array {
		fmt.Println(i, s)
	}

	array2 := [...]int{4, 5, 6, 7, 8, 234123, 412341234, 12341234, 1234, 123, 41, 23}

	for i, s := range array2 {
		fmt.Println(i, s)
	}
	// Why does that not work?
	/*
	   for s := range array2 {
	       fmt.Println(s)
	   }
	*/

	var cars = [4]string{"Switzerland", "Germany", "France"}
	fmt.Println(cars)
	fmt.Println(cars[3])

	var carsMake = make([]string, 0, 10)
	carsMake = append(carsMake, "Test")
	fmt.Println(len(carsMake))
	fmt.Println(cap(carsMake))
	fmt.Println(carsMake)

	grade := 5
	if grade > 5 {
		fmt.Println("WOW. You got over a 5")
	} else {
		fmt.Println("Ou. Well... Let's work for the 5 or higher next time.")
	}

	// Nested if
	grade = 6
	if grade >= 5 {
		fmt.Printf("Your grade is 5 or higher: %d", grade)
		fmt.Println()
		if grade > 5 {
			fmt.Printf("Your grade higher then 5:  %d", grade)
			fmt.Println()
		}
	} else {
		fmt.Printf("Your grade is lower than 5: %d", grade)
		fmt.Println()
	}

	grade = 7
	// Switch-Case switch
	// Can only be one datatype. SO I cannot put a number in the first case and a string in the second case.
	switch grade {
	case 1:
		fmt.Println("Your grade is", grade)

	case 2:
		fmt.Println("Your grade is", grade)
	case 3:
		fmt.Println("Your grade is", grade)
	case 4:
		fmt.Println("Your grade is", grade)
	case 5:
		fmt.Println("Your grade is", grade)
	case 6:
		fmt.Println("Your grade is", grade)
	default:
		fmt.Println("Your grade is over the number 6. Your input was:", grade)
	}

	// Multi-case switch

	day := 4
	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("It is a weekday")
	case 6, 7:
		fmt.Println("It is WEEKEND :)")
	default:
		fmt.Println("After Sunday, there is Monday. There is no", day, "th day")
	}

	// For loops with continue / break
	for i := 0; i < 10; i++ {
		if i == 8 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println()
	for i := 0; i < 10; i++ {
		if i == 4 {
			break
		}
		fmt.Println(i)
	}
	// Nested Loops W3Schools
	adj := [2]string{"big", "tasty"}
	fruits := [3]string{"apple", "orange", "banana"}
	for i := 0; i < len(adj); i++ {
		for j := 0; j < len(fruits); j++ {
			fmt.Println(adj[i], fruits[j])
		}
	}

	// Range loop
	for index, value := range fruits {
		fmt.Println(index, value)
	}
	fmt.Println()

	// Range loop without index
	for _, value := range fruits {
		fmt.Println(value)
	}
	fmt.Println()
	// Range loop without value
	for index, _ := range fruits {
		fmt.Println(index)
	}

	// Functions
	myMessage()

	myMessage()
	myMessage()
	myMessage()

	// Functins mit Parametern
	greeting("Shan")
	// Functions mit mehreren Parametern
	greetingWithSurname("Shan", "Dev")

	// Function with return
	fmt.Println(greetingReturn("Shan2"))

	// named return Values
	fmt.Println(greetingWithSurnameReturn("Shan2", "Dev2"))

	// mutliple return values
	fmt.Println(myFunction(5, "Hello"))

	// Save it into multiple variables
	a, b := myFunction(5, "Hello")
	fmt.Println(a, b)

	// Just save one of the variables
	_, c := myFunction(5, "Hello")
	fmt.Println(c)

	// Recursion
	// is when a function calls itself
	count(1)

	// Example
	fmt.Println()
	fmt.Println(factorial_recursion(4))
}

func factorial_recursion(i float64) (y float64) {
	if i > 0 {
		y = i * factorial_recursion(i-1)
	} else {
		y = 1
	}

	return
}

func count(i int) int {
	if i == 11 {
		return 0
	}
	fmt.Println(i)
	// cals itself. Common for mathematical and programming concepts
	return count(i + 1)
}

func myMessage() {
	fmt.Println("HELLLOOOO this is a function")
}

func greeting(name string) {
	fmt.Println("Hello", name)
}

func greetingWithSurname(firstname string, lastname string) {
	fmt.Println("Hello", firstname, lastname)
}
func greetingWithSurnameReturn(firstname string, lastname string) (result string) {
	result = "Hello " + firstname + " " + lastname
	// go will automatically return the value "result"
	return
}

func greetingReturn(name string) string {
	return "Hello " + name
}

func myFunction(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}
