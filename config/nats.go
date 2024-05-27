package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

var NatsConn *nats.Conn
var JST jetstream.JetStream
var CONS jetstream.Consumer

func init() {
	GetViperConfig()

}

func InitNats() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var err error
	NatsConn, err = nats.Connect("nats:4222")
	if err != nil {
		fmt.Println("fatal nats")
		log.Fatalln(err)
	}
	JST, err = jetstream.New(NatsConn)
	if err != nil {
		fmt.Println("fatal js")
		log.Fatalln(err)

	}

	log.Printf("creating stream %q and subject %q", ConfigValues.GetString("stream.name"), ConfigValues.GetString("stream.name")+".>")
	_, err = JST.CreateOrUpdateStream(ctx, jetstream.StreamConfig{
		Name:     ConfigValues.GetString("stream.name"),
		Subjects: []string{ConfigValues.GetString("stream.name") + ".>"},
	})
	if err != nil {
		fmt.Println("fatal js stream")

		log.Fatal(err)
	}

	stream, err := JST.Stream(ctx, ConfigValues.GetString("stream.name"))
    // log.Println(stream.Info(ctx))
    if err != nil {
		fmt.Println("stream err")

		log.Fatal(err)
	}

    CONS, err = stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
        AckPolicy: jetstream.AckExplicitPolicy,
        Name: "thumbnail_processor",
        // Durable: "thumbnail_processor",
    })
    if err != nil {
		fmt.Println("cosumer  err")

		log.Fatal(err)
	}
}
