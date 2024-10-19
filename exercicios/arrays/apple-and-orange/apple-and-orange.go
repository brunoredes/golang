package main

import "fmt"

func main() {
	apple, orange := countApplesAndOranges(7, 11, 5, 15, []int32{-2, 2, 1}, []int32{5, -6})
	fmt.Println(apple, "\n", orange)
}

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) (int, int) {
	// Write your code here
	appleCount := 0
	orangeCount := 0

	for _, apple := range apples {
		finalPosition := a + apple
		if finalPosition >= s && finalPosition <= t {
			appleCount++
		}
	}

	for _, orange := range oranges {
		finalPosition := b + orange // Posição final da laranja
		if finalPosition >= s && finalPosition <= t {
			orangeCount++
		}
	}

	return appleCount, orangeCount
}
