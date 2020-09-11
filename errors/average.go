package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more floats.")
		os.Exit(1)
	}
	sum, n := .0, 0
	for _, arg := range os.Args[1:] {
		a, err := strconv.ParseFloat(arg, 64)
		if err == nil {
			sum += a
			n += 1
		}
	}
	if n > 0 {
		fmt.Println("Average:", sum / float64(n))
	} else {
		fmt.Println("None of the arguments is a float!")
	}
}
