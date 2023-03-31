package main

import "fmt"

func main() {
	array := [4]float64{1, 2, 3, 4}

	sum := 1.0

	for i := 0; i < len(array); i++ {
		sum *= array[i]
	}
	fmt.Println(sum)
}
