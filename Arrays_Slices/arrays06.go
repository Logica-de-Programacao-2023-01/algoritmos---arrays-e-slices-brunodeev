package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var number int

	fmt.Print("\nConfira se um número está no array: ")
	fmt.Scan(&number)

	for i := 0; i < len(array); i++ {
		if array[i] == number {
			fmt.Println("\nEste número pertece ao array!")
		}
	}
	fmt.Println("\nEste número não pertence ao array!")

}
