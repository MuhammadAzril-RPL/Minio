package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	
	endpoint := "203.194.113.6:9000"
	accessKeyID := "MidFtK0wfiZ6AUjDfZbz"
	secretAccessKey := "KxkgFNq196ok2AKq9U5h2naOUq0Akpi8HyjA4RO3"
	useSSL := false

	
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	bucketName := "smk-telkom"
	objectName := "public/test.txt" 
	filePath := "./test.txt"

	fmt.Println("\n=== Upload File ===")
	content := []byte("Ini adalah contoh file txt yang akan diupload ke MinIO")
	err = os.WriteFile(filePath, content, 0644)
	if err != nil {
		log.Printf("Error membuat file: %v\n", err)
	}

	info, err := minioClient.FPutObject(context.Background(), bucketName, objectName, filePath,
		minio.PutObjectOptions{ContentType: "text/plain"})
	if err != nil {
		log.Printf("Error upload file: %v\n", err)
	} else {
		fmt.Printf("Berhasil upload %s dengan ukuran %d bytes\n", objectName, info.Size)
	}

	fmt.Println("\n=== Download File ===")
	// Download file
	err = minioClient.FGetObject(context.Background(), bucketName, objectName, "downloaded_"+filePath,
		minio.GetObjectOptions{})
	if err != nil {
		log.Printf("Error download file: %v\n", err)
	} else {
		fmt.Printf("Berhasil download file ke downloaded_%s\n", filePath)
	}

	fmt.Println("\n=== Daftar File dalam Bucket ===")
	// List files
	objectCh := minioClient.ListObjects(context.Background(), bucketName, minio.ListObjectsOptions{
		Recursive: true,
	})
	for object := range objectCh {
		if object.Err != nil {
			log.Printf("Error: %v\n", object.Err)
			continue
		}
		fmt.Printf("- %s (Size: %d bytes)\n", object.Key, object.Size)
	}

	fmt.Println("\n=== Generate URL ===")
	// Generate presigned URL
	presignedURL, err := minioClient.PresignedGetObject(context.Background(), bucketName, objectName, time.Hour*24, nil)
	if err != nil {
		log.Printf("Error membuat presigned URL: %v\n", err)
	} else {
		fmt.Printf("Presigned URL (berlaku 24 jam): %s\n", presignedURL.String())
	}

	publicURL := fmt.Sprintf("http://%s/%s/%s", endpoint, bucketName, objectName)
	fmt.Printf("Public URL: %s\n", publicURL)
	
	os.Remove(filePath)
	os.Remove("downloaded_" + filePath)

	fmt.Println("\nProgram selesai")
}
