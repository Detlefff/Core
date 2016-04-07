package main

import (
	"fmt"
	"log"

	"github.com/Detlefff/go-xmpp/xmpp"

	"./jabber"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	//Parse CLI-Flags
	kingpin.Parse()

	messageChannel := make(chan xmpp.Message)

	//Connect to the server
	//
	//If we connected successfully, start the sending routine
	go jabber.SendRoutine(messageChannel)

	//Connect to the XMPP-Server
	connection, _, error := jabber.Connect(*userName, *password, *debug)
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(connection.Jid)

	//Just wait
	for {
	}
}
