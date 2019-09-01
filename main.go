package main

import (
	"fmt"
	"log"
	"os"

	"github.com/IoTPanic/pixelpusher/internal/api"
)

func main() {
	log.Println("Launching PixelPusher... 🚀 🚀 🚀")

	go api.Start("0.0.0.0:8080")
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	log.Println("Waiting for sigterm")
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()
	<-done
	log.Println("Terminating Pixel Pusher instance gracefully")
}
