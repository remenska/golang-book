package main

import "fmt"

func main() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)
	var y [5]float64
	y[0] = 98
	y[1] = 93
	y[2] = 77
	y[3] = 82
	y[4] = 83

	var total float64 = 0
	for i := 0; i < len(y); i++ {
		total += y[i]
	}
	// fmt.Println(total / len(y))
	// .\tmp.go:19: invalid operation: total / 5 (mismatched types float64 and int)
	fmt.Println(total / float64(len(y)))

	// another way to achieve the same
	total = 0.0
	for _, value := range y {
		total += value
	}
	fmt.Println(total / float64(len(y)))

	var another_array = [3]float64{1, 2, 3}
	another_one := [3]float64{5, 6, 7}
	// another_one =another_array
	fmt.Println(another_array[2], another_one[2])

	slices()
	maps()
}

func slices() {
	var some_slice = []float64{1, 2, 3}
	// x := make([]float64, 5, 10)
	another_slice := append(some_slice, 4, 5)
	fmt.Println(some_slice)
	fmt.Println(another_slice)

	yet_another_slice := make([]float64, 2)
	copy(yet_another_slice, some_slice)
	fmt.Println(yet_another_slice)
}

func maps() {
	// var x map[string]int  // [keytype]valuetype
	// The problem with our program is that maps have to be initialized before they can be used.
	// x["key"] = 10
	x := make(map[string]int) // must be initialized before used
	x["key"] = 10
	x["another_key"] = 20
	fmt.Println(x["key"])
	fmt.Println(x)
	fmt.Println("len=", len(x))

	// another way of creating an array of ints
	int_array := make(map[int]int)
	int_array[1] = 11
	fmt.Println(int_array[1])
	delete(int_array, 1)
	fmt.Println(len(int_array))
	fmt.Println("int_array[1]=", int_array[1])
	value, ok := int_array[1] // second return value is false (ok==false)
	//Accessing an element of a map can return two values instead of just one.
	//The first value is the result of the lookup, the second tells us whether or not the lookup was successful.
	fmt.Println(value, ok)

	// Go way of doing it

	if value, ok := int_array[1]; ok {
		fmt.Println("YESS", value)
	} else {
		fmt.Println("PANIC", value)
	}

	// maps strings to maps of string->string
	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "Gas",
		},

		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
	}
	fmt.Println(elements)

	if el, ok := elements["He"]; ok {
		fmt.Println(el["name"], el["state"])
	}

	find_smallest()

}

func find_smallest() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	min_el := x[0]
	for _, el := range x {
		if el < min_el {
			min_el = el
		}
		// fmt.Println(el)
	}
	fmt.Println("Smallest number is, ", min_el)
}
