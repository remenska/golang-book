package main

import "fmt"

func average(xs []float64) float64 {
	// panic("Not implemented")
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))

}

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))
	// x, y := f()
	fmt.Println(add(1, 2, 3))
	ys := []int{1, 2, 3}
	fmt.Println(add(ys...))

	x := 0
	increment := func() int {
		x++
		return x
	}
	// A function like this together with the non-local variables it references is known as a closure.
	// In this case increment and the variable x form the closure.

	fmt.Println(increment())
	fmt.Println(increment())
	// closures
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	// recursion
	fmt.Println(factorial(3))

	// max
	fmt.Println("max=", greatest(1, 5, 7, 3, 4, 6, 9, 3))

	// panic & recover
	defer func() {
		str := recover()
		fmt.Println("PANIC!", str)
	}()
	// panic("PANIC!")
	result := []int{1, 2, 3}
	fmt.Println(result[5])
}

func f() (int, int) {
	return 5, 6
}

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func greatest(args ...float64) float64 {
	max := args[0]
	for _, val := range args {
		if max < val {
			max = val
		}
	}
	return max
}

// Each time it's called it adds 2 to the local i variable which – unlike normal local variables – persists between calls.
func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}
