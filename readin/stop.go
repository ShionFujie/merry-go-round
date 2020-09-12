package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	f := os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() && scanner.Text() != "STOP" {
		_, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err == nil {
			fmt.Println("Go on...")
		} else {
			io.WriteString(os.Stderr, "Please, enter an integer!!!\n")
		}
	}
}
