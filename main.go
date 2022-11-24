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
s
    array := [5]int{4, 5, 6, 7, 8}

    // Itterate through list
    for s := range array {
        fmt.Println(s)
    }

    // Itterate through list with giving the current position
    for i, s := range array {
        fmt.Println(i, s)
    }
}
