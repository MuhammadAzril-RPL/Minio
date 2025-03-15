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
	downloadedFilePath := "downloaded_" + filePath

	fmt.Println("\n=== Upload File ===")
	content := []byte("Ini adalah contoh file txt yang akan diupload ke MinIO")
	err = os.WriteFile(filePath, content, 0644)
	if err != nil {
		log.Fatalf("Error membuat file: %v\n", err)
	}

	info, err := minioClient.FPutObject(context.Background(), bucketName, objectName, filePath,
		minio.PutObjectOptions{ContentType: "text/plain"})
	if err != nil {
		log.Fatalf("Error upload file: %v\n", err)
	} else {
		fmt.Printf("Berhasil upload %s dengan ukuran %d bytes\n", objectName, info.Size)
	}

	fmt.Println("\n=== Download File ===")
	err = minioClient.FGetObject(context.Background(), bucketName, objectName, downloadedFilePath,
		minio.GetObjectOptions{})
	if err != nil {
		log.Fatalf("Error download file: %v\n", err)
	} else {
		fmt.Printf("Berhasil download file ke %s\n", downloadedFilePath)
	}
	fmt.Println("\n=== Daftar File dalam Bucket ===")
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
	presignedURL, err := minioClient.PresignedGetObject(context.Background(), bucketName, objectName, time.Hour*24, nil)
	if err != nil {
		log.Fatalf("Error membuat presigned URL: %v\n", err)
	} else {
		fmt.Printf("Presigned URL (berlaku 24 jam): %s\n", presignedURL.String())
	}

	publicURL := fmt.Sprintf("http://%s/%s/%s", endpoint, bucketName, objectName)
	fmt.Printf("Public URL: %s\n", publicURL)

	var pilihan string
	fmt.Print("\nApakah Anda ingin menghapus file lokal? (y/n): ")
	fmt.Scanln(&pilihan)

	if pilihan == "y" || pilihan == "Y" {
		os.Remove(filePath)
		os.Remove(downloadedFilePath)
		fmt.Println("File telah dihapus.")
	} else {
		fmt.Println("File tidak dihapus.")
	}

	fmt.Println("\nProgram selesai")
}
