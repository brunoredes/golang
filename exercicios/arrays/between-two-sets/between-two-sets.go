package main

import (
	"fmt"
)

func main() {
	total := getTotalX([]int32{1, 2}, []int32{72, 48})
	fmt.Println(total)
}

func getTotalX(a []int32, b []int32) int {
	higherNumberOnA, _ := MinMax(a)
	_, lowerNumberOnB := MinMax(b)
	totalInterval := interval(higherNumberOnA, lowerNumberOnB)
	var dividable []int32
	var result []int32

	// Verificação se o número do intervalo é divisível por todos os elementos de a
	for i := 0; i < len(totalInterval); i++ {
		isDivisible := true
		for _, num := range a {
			if totalInterval[i]%num != 0 {
				isDivisible = false
				break
			}
		}
		if isDivisible {
			dividable = append(dividable, totalInterval[i])
		}
	}

	// Verificação se os números divisíveis por a também dividem todos os elementos de b
	for j := 0; j < len(dividable); j++ {
		isFactor := true
		for k := 0; k < len(b); k++ {
			if b[k]%dividable[j] != 0 {
				isFactor = false
				break
			}
		}
		if isFactor {
			result = addToSlice(result, dividable[j])
		}
	}

	return len(result)
}

func interval(x int32, y int32) []int32 {
	var result []int32
	var i = x

	for i <= y {
		result = append(result, i)
		i++
	}
	return result
}

func contains(slice []int32, num int32) bool {
	for _, v := range slice {
		if v == num {
			return true
		}
	}
	return false
}

// Function to add a number to the slice
func addToSlice(slice []int32, num int32) []int32 {
	if !contains(slice, num) {
		slice = append(slice, num)
	}
	return slice
}

func MinMax(array []int32) (int32, int32) {
	var max int32 = array[0]
	var min int32 = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}
