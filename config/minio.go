package config

import (
	"context"
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	// "github.com/spf13/viper"
	// "github.com/fsmardani/go-for-example/main"
	// "github.com/fsmardani/go-for-example/models"
	// "github.com/gofiber/storage/minio"
)

// var MinioClient *minio.Client


func MinioConnection()(*minio.Client, error){
    ctx := context.Background()
    endpoint := "minio:9000"
    accessKeyID := "elq1BRA1F0MjqEYJTBRJ"
	// viper.GetString("MINIO_ACCESS_KEY_ID")
    secretAccessKey := "WVdUY7SUQZs7DEQUh8MbNvcD4HGVmBKlpQqaNEvc"
	// viper.GetString("MINIO_SECRET_ACCESS_KEY")
    useSSL := false
	fmt.Println(accessKeyID,secretAccessKey)
    // Initialize minio client object.
    minioClient, errInit := minio.New(endpoint, &minio.Options{
        Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
        Secure: useSSL,
    })
    if errInit != nil {
		fmt.Println("minio")
        log.Fatalln(errInit)
    }

    // Make a new bucket called dev-minio.
    bucketName := "images"
    location := "us-east-1"

    err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
    if err != nil {
        // Check to see if we already own this bucket (which happens if you run this twice)
        exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
        if errBucketExists == nil && exists {
            log.Printf("We already own %s\n", bucketName)
        } else {
            log.Fatalln(err)
        }
    } else {
        log.Printf("Successfully created %s\n", bucketName)
    }
    return minioClient, errInit

    // info, err := MinioClient.FPutObject(ctx, bucketName,"images.jpeg", "images.jpeg", minio.PutObjectOptions{ContentType: "image/jpeg"})

    // if err != nil {
    //     log.Printf("Successfully not uploaded of size %d\n", info.Size)

    // }

    // log.Printf("Successfully uploaded of size %d\n", info.Size)


}