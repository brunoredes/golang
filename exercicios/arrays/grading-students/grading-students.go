package main

import (
	"fmt"
)

// import "math"

func main() {
	gradingStudents([]int32{4, 73, 67, 38, 33})
}

func gradingStudents(grades []int32) []int32 {
	// var notas []int32
	var result []int32

	for _, grade := range grades {
		roundedNumber := roundNumber(grade)
		result = append(result, roundedNumber)
	}
	fmt.Println(result)
	return result
}

func roundNumber(num int32) int32 {
	if num < 38 {
		return num
	} else if 5-(num%5) < 3 {
		return (num/5 + 1) * 5
	} else {
		return num
	}
}
