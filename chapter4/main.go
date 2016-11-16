package main
import "fmt"
import "math"

var something_outside string = "Hello from here"
var (
	a = 5
	b = 10
	c = 15
	d = "sosgdgha"
)
func main() {
	var x string = "Hello World"
	fmt.Println(x)

	var y string
	y = "Hello again!"
	fmt.Println(y)

	var z string
	z = "first "
	fmt.Println(z)
	z = z + "second"
	fmt.Println(z)
	fmt.Println(z == y)

	xx := "Hello World"
	fmt.Println(xx)
	// zz := 5 
	// fmt.Println(zz)
	fmt.Println(something_outside)
	f()

	const some_constant string = "Hello agian from constant"
	fmt.Println(some_constant)
	fmt.Println(math.Pi)
	// ./main.go:30: cannot assign to some_constant
	// some_constant = "something else
	fmt.Println(d)
	g()
}

func f() {
	fmt.Println(something_outside)
}

func g() {
	fmt.Println("Enter a number: ")
	var some_input float64
	fmt.Scanf("%f", &some_input)
	output := some_input * 2
	fmt.Println(output)
}