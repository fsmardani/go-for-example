package config

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

var NatsConn *nats.Conn
var JetSTRM nats.JetStreamContext

func InitNats() {
    var err error
    NatsConn, err = nats.Connect("nats:4222")
    if err != nil {
        fmt.Println("fatal nats")
        log.Fatalln(err)
    }
    JetSTRM, err = NatsConn.JetStream()
    if err != nil {
        fmt.Println("fatal js")
        log.Fatalln(err)

    }
}