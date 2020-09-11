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
	var min, max float64
	b := false
	for _, arg := range os.Args[1:] {
		n, err := strconv.ParseFloat(arg, 64)
		if err == nil {
			b = true
			if n < min {
				min = n
			}
			if n > max {
				max = n
			}
		}
	}
	if b {
		fmt.Println("Min:", min)
		fmt.Println("Max:", max)
	} else {
		fmt.Println("None of the arguments is a float!")
	}
}
