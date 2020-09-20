package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var port = flag.Int("port", 8000, "port number")

func main() {
	flag.Parse()
	address := fmt.Sprintf("localhost:%d", *port)
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
