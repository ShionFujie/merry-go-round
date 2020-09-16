package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handleSignal(sig os.Signal) {
	fmt.Println("We're handling", sig)
}

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs)
	go func() {
		for {
			sig := <-sigs
			switch sig {
			case os.Interrupt:
				handleSignal(sig)
			case syscall.SIGUSR2:
				fmt.Println("We're witnessing", sig)
			case syscall.SIGTERM:
				handleSignal(sig)
				os.Exit(0)
			default:
				fmt.Println("We're pretending not to see", sig)
			}
		}
	}()
	for {
		fmt.Printf(".")
		time.Sleep(30 * time.Second)
	}
}
