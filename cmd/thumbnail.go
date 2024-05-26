package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/disintegration/imaging"
	"github.com/minio/minio-go/v7"
	"github.com/nats-io/nats.go"

	"github.com/fsmardani/go-for-example/config"
)

func init() {
	config.MinioConnection()
	config.InitNats()

}

// func createThumbnail(natsConn *nats.Conn) {
func createThumbnail() {
    _, err := config.JST.Subscribe("IMAGES.uploaded", func(m *nats.Msg) {
        fileName := string(m.Data)
        ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
        defer cancel()
        

        minioClient, err:= config.MinioConnection()
        obj, err := minioClient.GetObject(ctx, "images", fileName, minio.GetObjectOptions{})
        if err != nil {
            log.Println(err)
            return
        }
        defer obj.Close()

        img, err := imaging.Decode(obj)
        if err != nil {
            log.Println(err)
            return
        }

        thumbnail := imaging.Thumbnail(img, 100, 100, imaging.Lanczos)
        thumbFileName := strings.Split(fileName, ".")[0] + "_thumbnail."+strings.Split(fileName, ".")[1]

        thumbFile, err := os.Create(thumbFileName)
        if err != nil {
            log.Println(err)
            return
        }
        defer thumbFile.Close()

        err = imaging.Encode(thumbFile, thumbnail, imaging.JPEG)
        if err != nil {
            log.Println(err)
            return
        }

        thumbFile.Seek(0, 0)
        _, err = minioClient.PutObject(ctx, "images", thumbFileName, thumbFile, -1, minio.PutObjectOptions{ContentType: "image/jpeg"})
        if err != nil {
            log.Println(err)
        }
		os.Remove(thumbFileName)
    })
	// 
    if err != nil {
		fmt.Println("reached cons")
        log.Fatalln(err)
    }
}
