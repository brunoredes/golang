package main

import "fmt"

func main() {
	kang := kangaroo(0, 3, 4, 2)
	kang2 := kangaroo(0, 2, 5, 3)
	fmt.Println(kang)
	fmt.Println(kang2)
}

/*
x1 starting position
v1 jump distance for kangaroo 1

x2 starting position
v2 jump distance for kangaroo 2

returns "YES" or "NO"
*/
func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {

	if v1 == v2 {
		if x1 == x2 {
			return yesOrNo(true)
		}
		return yesOrNo(false)
	}

	if (v1 > v2) && (x1-x2)%(v2-v1) == 0 {
		return yesOrNo(true)
	}
	return yesOrNo(false)

}

func yesOrNo(param bool) string {
	if param == true {
		return "YES"
	}
	return "NO"
}
