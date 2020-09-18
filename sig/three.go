package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handleSignal(signal os.Signal) {
	fmt.Println("Handling", signal)
}

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGINFO, syscall.SIGTERM)
	go func() {
		for {
			sig := <-sigs
			switch sig {
			case os.Interrupt:
				fmt.Println("Ignoring", sig)
			case syscall.SIGINFO:
				handleSignal(sig)
			case syscall.SIGTERM:
				fmt.Println("Terminating...")
				os.Exit(0)
			}
		}
	}()
	for {
		fmt.Printf(".")
		time.Sleep(20 * time.Second)
	}
}
