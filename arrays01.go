package main

import "fmt"

func main() {
	numbers := []int{3, 5, 7}

	var sum int

	for i := 0; i < len(numbers); i++ {

		sum += numbers[i]
	}

	fmt.Println("\nA soma dos números é: ", sum)
}
