package main

import (
	"io/ioutil"
	"log"
	"os"
	"fmt"
)

func main() {
	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	// defer f.Close()

	// var buf []byte
	// if _, err := io.ReadFull(f, buf); err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Printf("%s\n", b)
}