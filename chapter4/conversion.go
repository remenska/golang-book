package main

import "fmt"

func main() {
	fmt.Println("Enter degrees in Farenheit")
	var degrees_farenheit float32
	fmt.Scanf("%f", &degrees_farenheit)
	degrees_celsius := (degrees_farenheit - 32) * 5 / 9
	fmt.Println("Degrees in celsius:", degrees_celsius)

}
