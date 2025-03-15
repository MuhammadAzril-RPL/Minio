# MinIO Golang Client

## Deskripsi
Program ini dibuat dalam bahasa **Golang** untuk berinteraksi dengan MinIO. Program ini dapat melakukan beberapa operasi dasar seperti:

- Koneksi ke MinIO
- Upload file ke bucket
- Download file dari bucket
- Menampilkan daftar file dalam bucket
- Menghapus file dari bucket
- Membuat URL presigned
- Membuat URL public

## Persyaratan
Sebelum menjalankan program, pastikan:
- **Golang** sudah terinstall di sistem Anda.
- **MinIO Server** telah berjalan.
- Kredensial akses MinIO tersedia.
- Library `minio-go` sudah diinstall.

```sh
 go get github.com/minio/minio-go/v7
```

## Instalasi dan Konfigurasi
1. Clone repositori atau salin kode ini ke dalam file `main.go`.
2. Pastikan MinIO Server berjalan dan sesuaikan kredensial berikut:

```go
endpoint := "203.194.113.6:9000"
accessKeyID := "MidFtK0wfiZ6AUjDfZbz"
secretAccessKey := "KxkgFNq196ok2AKq9U5h2naOUq0Akpi8HyjA4RO3"
bucketName := "smk-telkom"
```

## Alur Program

1. **Inisialisasi koneksi ke MinIO**  
   - Menggunakan `minio.New()` dengan `accessKeyID` dan `secretAccessKey`.

2. **Membuat file contoh untuk diupload**  
   - File `test.txt` dibuat menggunakan `os.WriteFile()`.

3. **Mengunggah file ke MinIO**  
   - File diunggah menggunakan `FPutObject()` dengan path `public/test.txt`.

4. **Mengunduh file dari MinIO**  
   - File diunduh dengan `FGetObject()` dan disimpan sebagai `downloaded_test.txt`.

5. **Menampilkan daftar file dalam bucket**  
   - Menggunakan `ListObjects()` untuk menampilkan semua file dalam bucket.

6. **Membuat URL presigned**  
   - Menggunakan `PresignedGetObject()` untuk menghasilkan URL yang berlaku selama 24 jam.

7. **Membuat URL public**  
   - Dibentuk manual dengan format `http://endpoint/bucket/object`.

8. **Menghapus file sementara**  
   - File lokal (`test.txt` dan `downloaded_test.txt`) dihapus dengan `os.Remove()`.

## Menjalankan Program
Setelah semuanya dikonfigurasi, jalankan perintah berikut:
```sh
go run main.go
```

## Output Contoh
```sh
=== Upload File ===
Berhasil upload public/test.txt dengan ukuran 30 bytes

=== Download File ===
Berhasil download file ke downloaded_test.txt

=== Daftar File dalam Bucket ===
- public/test.txt (Size: 30 bytes)

=== Generate URL ===
Presigned URL (berlaku 24 jam): http://203.194.113.6:9000/smk-telkom/public/test.txt?X-Amz-Signature=...
Public URL: http://203.194.113.6:9000/smk-telkom/public/test.txt

Program selesai
```





