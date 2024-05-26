package config

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

)

// var MinioClient *minio.Client

func init(){
	GetViperConfig()

}

func MinioConnection()(*minio.Client, error){
    ctx := context.Background()
    endpoint := "minio:9000"
    accessKeyID := ConfigValues.GetString("minio.access_key_id")
    secretAccessKey := ConfigValues.GetString("minio.secret_access_key")
    useSSL := false

	minioClient, errInit := minio.New(endpoint, &minio.Options{
        Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
        Secure: useSSL,
    })
    if errInit != nil {
        log.Fatalln(errInit)
    }

    bucketName := "images"
    location := "us-east-1"

    err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
    if err != nil {
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
}