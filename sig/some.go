package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handleSignal(sig os.Signal) {
	fmt.Println("We're handling ", sig)
}

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGINFO)
	go func() {
		for {
			sig := <-sigs
			switch sig {
			case os.Interrupt:
				fmt.Println("We're ignoring ", sig)
			case syscall.SIGINFO:
				handleSignal(sig)
				return
			}
		}
	}()
	for {
		fmt.Printf(".")
		time.Sleep(20 * time.Second)
	}
}
