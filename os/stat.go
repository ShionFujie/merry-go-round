package main

import (
	"fmt"
	"os"
)

func main() {
	for _, pathname := range os.Args[1:] {
		fileinfo, err := os.Stat(pathname)
		if err != nil {
			fmt.Printf("%v", err)
			os.Exit(1)
		}
		fmt.Printf("%v", fileinfo)
	}
}