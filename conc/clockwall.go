package main

import (
	"strings"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {
		address := strings.SplitN(arg, "=", 2)[1]
		go connect(address)
	}
	for { }
}

func connect(address string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
