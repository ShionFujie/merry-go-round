package main

import (
	"fmt"
	"example.com/user/hello/morestrings"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("Hello, SHION!!!"))
	fmt.Println(cmp.Diff("Hello World", "Hello SHION"))
}
