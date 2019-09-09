package messaging

import (
	"fmt"
	"log"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var rxHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("[RECIEVED] TOPIC: %s\n", msg.Topic())
	fmt.Printf("[RECIEVED] MSG: %s\n", msg.Payload())
	go consumeMessage(msg.Payload(), msg.Topic())
}

var DownstreamChannel chan channelMessage

// Start starts the mqtt device interface listening
func Start(ClientID string, host string, username string, password string) {
	//mqtt.DEBUG = log.New(os.Stdout, "", 0)
	//mqtt.ERROR = log.New(os.Stdout, "", 0)
	opts := mqtt.NewClientOptions().AddBroker(host).SetUsername(username).SetPassword(password).SetClientID(ClientID)
	opts.SetKeepAlive(5 * time.Second)
	//opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(2 * time.Second)
	//opts.OnConnectionLost(connLostHandler)
	opts.AutoReconnect = true

	opts.OnConnect = func(c mqtt.Client) {
		fmt.Printf("Client connected, subscribing to required topics\n")

		//Subscribe here, otherwise after connection lost,
		//you may not receive any message
		if token := c.Subscribe("fixtures/+/up", 0, rxHandler); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
			os.Exit(1)
		}
	}

	c := mqtt.NewClient(opts)
	time.Sleep(2 * time.Second)
	fails := 0
	for {
		if token := c.Connect(); token.Wait() && token.Error() != nil {
			log.Println("[ ERROR ] Failed to connect to MQTT... ", token.Error().Error(), " trying again in 5 sec...")
			time.Sleep(10 * time.Second)

			//if fails > 5 {
			//	panic(token.Error().Error())
			//}
			fails++
		} else {
			break
		}
	}

	log.Println("Connected to MQTT broker")
	DownstreamChannel = make(chan channelMessage, 1)
	for {
		msg := <-DownstreamChannel
		t := c.Publish("fixtures/"+msg.Device+"/dwn", 0, false, msg.Payload)
		if t.Error() != nil {
			log.Printf("Error publishing - %s", t.Error().Error())
		}
	}
}
