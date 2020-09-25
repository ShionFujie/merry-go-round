package main

import (
	"fmt"
)


func main() {
	names := make(chan string)
	greetings := make(chan string)

	go func() {
		for _, x := range []string{"Eric", "Sheryl", "Peter"} {
			names <- x
		}
		close(names)
	}()

	go func() {
		for x := range names {
			greetings <- fmt.Sprintf("Hello, %s!!!", x)
		}
		close(greetings)
	}()

	for x := range greetings {
		fmt.Println(x)
	}
}