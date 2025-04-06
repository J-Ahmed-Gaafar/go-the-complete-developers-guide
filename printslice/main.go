package main

import "fmt"

func main() {
	//ns := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//for i := 0; i < len(ns); i++ {
	//	if ns[i]%2 == 0 {
	//		fmt.Printf("The number %v is Even.\n", ns[i])
	//	} else {
	//		fmt.Printf("The number %v is Odd.\n", ns[i])
	//	}
	//}

	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, number := range numbers {
		if number%2 == 0 {
			fmt.Println(number, "is even")
		} else {
			fmt.Println(number, "is odd")
		}
	}
}
