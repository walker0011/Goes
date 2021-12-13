package main

import (
	"fmt"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

// func main() {
// 	f("direct")       //  calling a function
// 	go f("goroutine") // invoke the function in a goroutine, this will execute concurrently with the calling one

// 	go func(msg string) { // start a goroutine for an anonymous function call
// 		fmt.Println(msg)
// 	}("going")

// 	time.Sleep(time.Second)
// 	fmt.Println("done")
// }
