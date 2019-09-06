package api

import (
	"fmt"
	"log"

	"github.com/gorilla/mux"
	"github.com/pschlump/MiscLib"
	"github.com/pschlump/godebug"
	"github.com/pschlump/socketio"
)

func applySocketConnection(r *mux.Router) {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	server.On("connection", func(so socketio.Socket) {
		fmt.Println("Socket connection")
		so.Join("test")
		so.On("ping", func() string {
			fmt.Println("PONG")
			so.BroadcastTo("test", "pong", "Hello World")
			return "PONG"
		})
		so.On("disconnect", func() {
			fmt.Printf("%suser disconnect%s, %s\n", MiscLib.ColorYellow, MiscLib.ColorReset, godebug.LF())
		})
	})

	server.On("error", func(so socketio.Socket, err error) {
		fmt.Printf("Error: %s, %s\n", err, godebug.LF())
	})
	r.Handle("/socket.io/", server)
}
