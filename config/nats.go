package config

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

var NatsConn *nats.Conn
var JST nats.JetStreamContext


func init(){
    GetViperConfig()

}


func InitNats() {
    var err error
    NatsConn, err = nats.Connect("nats:4222")
    if err != nil {
        fmt.Println("fatal nats")
        log.Fatalln(err)
    }
    JST, err = NatsConn.JetStream()
    if err != nil {
        fmt.Println("fatal js")
        log.Fatalln(err)

    }
    stream, err := JST.StreamInfo(ConfigValues.GetString("stream.name"))
	if err != nil {
		log.Print(err)
	}
	if stream == nil {
		log.Printf("creating stream %q and subject %q", ConfigValues.GetString("stream.name"), ConfigValues.GetString("stream.name")+".*")
		_, err = JST.AddStream(&nats.StreamConfig{
			Name:     ConfigValues.GetString("stream.name"),
			Subjects: []string{ConfigValues.GetString("stream.name")+".*"},
		})
        if err != nil {
            fmt.Println("fatal js stream")

            log.Fatal(err)
        }	}


}