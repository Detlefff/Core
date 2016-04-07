package jabber

import (
	"crypto/tls"
	"fmt"
	"log"

	"github.com/Detlefff/go-xmpp/xmpp"
)

//Connection holds the connection to the server
var Connection *xmpp.Client

//Connect connects to a XMPP-Server
func Connect(user, password string, debug bool) (*xmpp.Client, chan xmpp.Status, error) {
	var error error

	tlsconf := tls.Config{InsecureSkipVerify: true}
	statusChan := make(chan xmpp.Status)
	jid := xmpp.JID(user)

	if debug {
		go func() {
			for status := range statusChan {
				log.Printf("connection status %d", status)
			}
		}()
	}

	Connection, error = xmpp.NewClient(&jid, password, &tlsconf, nil, xmpp.Presence{}, statusChan)
	if error != nil {
		return nil, statusChan, error
	}

	return Connection, statusChan, error
}

func SendRoutine(messageChannel chan xmpp.Message) {
	for message := range messageChannel {
		//Send message
		fmt.Println(message)
	}
}
