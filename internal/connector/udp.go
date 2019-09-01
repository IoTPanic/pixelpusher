package connector

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/IoTPanic/pixels-go"
)

type connection struct {
	ID    uint8
	addr  *net.UDPAddr
	input chan []pixels.Pixel
}

var connections []connection

func createConnection(ctx context.Context, address string, nodeID uint8, sessionID uint8) error {
	raddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return errors.New("Could not resolve address")
	}
	conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		return errors.New("Failed to open UDP connection")
	}

	defer conn.Close()

	pixelChan := make(chan []pixels.Pixel, 1)
	doneChan := make(chan error, 1)

	connections = append(connections, connection{nodeID, raddr, pixelChan})
	// Concurrent go routine to keep connection alive
	go func() {
		for p := range pixelChan {
			fmt.Println(p) // TODO add UDP send call
		}
		doneChan <- nil
	}()

	select {
	case <-ctx.Done():
		log.Println("UDP connection was closed from context")
		err = ctx.Err()
	case err = <-doneChan:
	}

	return nil
}
