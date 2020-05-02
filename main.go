package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/IoTPanic/pixelpusher/internal/api"
	"github.com/IoTPanic/pixelpusher/internal/db"
)

var helptext = `
___  _ _  _ ____ _    ____ ____ ____ ____ _  _ ____ ____ 
|__] |  \/  |___ |    |    |__/ |__| [__  |__| |___ |__/ 
|    | _/\_ |___ |___ |___ |  \ |  | ___] |  | |___ |  \
========================================================

    Help Text

  This is the help menu for pixelcrasher, we recommend 
you read the README.md file if you have access to the git
repository.
`

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
	log.Println("Launching PixelPusher...")
	// Get the standard variables
	arguments := os.Args[1:]
	if len(arguments) > 0 {
		switch arguments[0] {
		case "help":
			fmt.Println(helptext)
			os.Exit(1)
		case "dev":
			os.Setenv("DBDIR", "./db/")
			break
		}
	}
	fmt.Println(arguments)

	// Setup the database
	dbPath := os.Getenv("DBDIR")
	os.MkdirAll(dbPath, 0700) // Create db dir if nonexistant
	dbPath = dbPath + "pixelsDB.db"
	err := db.Connect(dbPath)
	if err != nil {
		panic(err)
	}

	// Start the API
	go api.Start("0.0.0.0:8080")

	// Setup code to allow for a graceful termination
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
