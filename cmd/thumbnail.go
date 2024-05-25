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

	"github.com/fsmardani/go-for-example/config"
)

func init() {
	config.InitNats()
	config.MinioConnection()

}

func createThumbnail(natsConn *nats.Conn) {
    _, err := natsConn.Subscribe("images.uploaded", func(m *nats.Msg) {
        fileName := string(m.Data)
        ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
        defer cancel()

        obj, err := config.MinioClient.GetObject(ctx, "images", fileName, minio.GetObjectOptions{})
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
        thumbFileName := strings.TrimSuffix(fileName, ".jpg") + "_thumbnail.jpg"

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
        _, err = config.MinioClient.PutObject(ctx, "images", thumbFileName, thumbFile, -1, minio.PutObjectOptions{ContentType: "image/jpeg"})
        if err != nil {
            log.Println(err)
        }
    })
    if err != nil {
        log.Fatalln(err)
    }
}
