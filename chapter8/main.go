package main

import "fmt"

// func zero(x int) {
// 	x = 0
// }

func zero(xPtr *int) {
	*xPtr = 0
	fmt.Println("xPtr", *xPtr)
}

func square(x *float64) {
	*x = *x * *x
}

func swap(x *int, y *int) {
	temp := new(int)
	*temp = *x
	*x = *y
	*y = *temp
}

func main() {
	x := 5
	zero(&x)
	fmt.Println(x)
	y := new(int)
	zero(y)
	fmt.Println(*y)
	z := 1.5
	square(&z)
	fmt.Println(z)
	a := 22
	b := 33
	fmt.Println("Before swap: ", a, b)
	swap(&a, &b)
	fmt.Println("After swap: ", a, b)
}
