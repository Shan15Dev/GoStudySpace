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

	grade = 6
	if grade >= 5 {
        fmt.Printf("Your grade is 5 or higher: %d", grade)
		if grade > 5 {
            fmt.Printf("Your grade higher then 5:  %d", grade)
		}
	} else {
        fmt.Printf("Your grade is lower than 5: %d", grade)
	}
}
