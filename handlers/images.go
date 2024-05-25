package handlers

import (
	"context"
	// "fmt"
	// "fmt"
	// "time"
	// "io"
	// "bytes"
	"log"

	"github.com/fsmardani/go-for-example/config"
    // "github.com/nats-io/nats.go"

	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
)

func init() {
	config.InitNats()
	config.MinioConnection()

}

func UploadFile(c *fiber.Ctx) error {
    ctx := context.Background()
    bucketName := "images"
    file, err := c.FormFile("fileUpload")

    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": true,
            "msg":   err.Error(),
        })
    }

    // Get Buffer from file
    buffer, err := file.Open()

    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": true,
            "msg":   err.Error(),
        })
    }
    defer buffer.Close()

     // Create minio connection.
    minioClient := config.MinioClient
    // if err != nil {
    //             // Return status 500 and minio connection error.
    //     return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
    //         "error": true,
    //         "msg":   err.Error(),
    //     })
    // }

    objectName := file.Filename
    fileBuffer := buffer
    contentType := file.Header["Content-Type"][0]
    fileSize := file.Size

    // Upload the zip file with PutObject
    info, err := minioClient.PutObject(ctx, bucketName, objectName, fileBuffer, fileSize, minio.PutObjectOptions{ContentType: contentType})

    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": true,
            "msg":   err.Error(),
        })
    }

    log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)

    return c.JSON(fiber.Map{
        "error": false,
        "msg":   nil,
        "info":  info,
    })
}