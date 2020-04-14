package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/IoTPanic/pixelpusher/internal/api"
	"github.com/IoTPanic/pixelpusher/internal/cache"
	"github.com/IoTPanic/pixelpusher/internal/db"
	"github.com/IoTPanic/pixelpusher/internal/messaging"
)

// Used to scan docker filesystem to try to find db issues; now fixed, just keeping in here
func scanf() {
	var files []string

	root := "/go/src/github.com/IoTPanic/pixelpusher/"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}
}

func main() {
	//scanf()
	log.Println("Launching PixelPusher... 🚀 🚀 🚀")

	arguments := os.Args[1:]
	if len(arguments) > 0 {
		switch arguments[0] {
		case "help":
			fmt.Println("Help text coming soon")
			os.Exit(1)
		case "dev":
			os.Setenv("DBDIR", "./db/")
			break
		}
	}

	dbPath := os.Getenv("DBDIR")


	os.MkdirAll(dbPath, 0700) // Create db dir if nonexistant
	dbPath = dbPath + "pixelsDB.db"
	err := db.Connect(dbPath)
	if err != nil {
		panic(err)
	}
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
	db.Close()
	log.Println("Terminating Pixel Pusher instance gracefully")
}
