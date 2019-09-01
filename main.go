package main

import (
	"log"

	"github.com/IoTPanic/pixelpusher/internal/api"
)

func main() {
	log.Println("Launching PixelPusher... 🚀 🚀 🚀")

	api.Start("0.0.0.0:8080")
}
