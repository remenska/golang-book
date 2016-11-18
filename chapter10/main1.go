package main

import (
	"fmt"
	"time"
)

func mySleep(duration time.Duration) {
	_ = <-time.After(duration)
}

func main() {
	fmt.Println("asdasdsa")
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			// time.Sleep(time.Second * 2)
			mySleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "from 2"
			// time.Sleep(time.Second * 3)
			mySleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			case howmuchtime := <-time.After(time.Second):
				fmt.Println("Timeout", howmuchtime)
			}

		}
	}()

	var input string
	fmt.Scanln(&input)
}
