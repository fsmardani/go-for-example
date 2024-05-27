package main

import (
	"context"
	// "fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/disintegration/imaging"
	"github.com/minio/minio-go/v7"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"

	"github.com/fsmardani/go-for-example/config")

func init() {
	config.MinioConnection()
	config.InitNats()
}

func createThumbnail() {
	for {
		msgs, err := config.CONS.Fetch(1, jetstream.FetchMaxWait(5*time.Second))
		if err != nil {
			if err == nats.ErrTimeout {
				continue
			}
			log.Printf("Error fetching messages: %v", err)
			continue
		}

		for msg := range msgs.Messages() {
			handleMessage(msg)
		}
	}
}

func handleMessage(msg jetstream.Msg) {
	log.Printf("Received message: %s", msg.Data())
	fileName := string(msg.Data())
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	minioClient, err := config.MinioConnection()
	if err != nil {
		log.Println("Error connecting to Minio:", err)
		return
	}

	obj, err := minioClient.GetObject(ctx, "images", fileName, minio.GetObjectOptions{})
	if err != nil {
		log.Println("Error getting object from Minio:", err)
		return
	}
	defer obj.Close()

	img, err := imaging.Decode(obj)
	if err != nil {
		log.Println("Error decoding image:", err)
		return
	}

	thumbnail := imaging.Thumbnail(img, 100, 100, imaging.Lanczos)
	thumbFileName := strings.Split(fileName, ".")[0] + "_thumbnail." + strings.Split(fileName, ".")[1]

	thumbFile, err := os.Create(thumbFileName)
	if err != nil {
		log.Println("Error creating thumbnail file:", err)
		return
	}
	defer thumbFile.Close()

	err = imaging.Encode(thumbFile, thumbnail, imaging.JPEG)
	if err != nil {
		log.Println("Error encoding thumbnail:", err)
		return
	}

	thumbFile.Seek(0, 0)
	_, err = minioClient.PutObject(ctx, "images", thumbFileName, thumbFile, -1, minio.PutObjectOptions{ContentType: "image/jpeg"})
	if err != nil {
		log.Println("Error putting object to Minio:", err)
	}

	os.Remove(thumbFileName)

	// Acknowledge the message
	msg.Ack()
}

