package main

import "fmt"

func main() { 
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}
	another_way()
	yet_another()
	switch_statement()
	test_me()
	problem()
}

func another_way() { 
	for i := 1; i <=10; i++ {
		fmt.Println(i)
	}
}

func yet_another() { 
	for i := 1; i <= 10; i++ {
		if i % 2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}

	}
}

func switch_statement() {
	for i := 0; i <= 10; i++ {
		switch i {
			case 0: fmt.Println("zero")
			case 1: fmt.Println("one")
			case 3: fmt.Println("three")
			default: fmt.Println("got tired of counting...")
		}
	}
}

func test_me() {
	i := 10
	if i > 10 {
		fmt.Println("Big")
	} else {
		fmt.Println("Small")
	}
}

func problem() {

	for i:= 1; i <= 100; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println(i, "FizzBuzz")
		} else if i % 3 == 0 {
			fmt.Println(i, "Fizz")
		} else if i % 5 == 0{
			fmt.Println(i, "Buzz")
		}
	}

}




