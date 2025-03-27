# Golang Context

Dokumentasi ini menjelaskan tentang penggunaan `context` dalam Golang, yang berguna untuk mengontrol eksekusi goroutine dengan batasan waktu, pembatalan, dan nilai yang diwariskan.

## Struktur Materi

### 1\. Pendahuluan

Penjelasan dasar tentang apa itu `context` dan mengapa penting dalam pemrograman Golang, terutama untuk menangani goroutine.

### 2\. Pengenalan Context

Memahami konsep dasar `context`, bagaimana cara menggunakannya, dan skenario di mana context sangat berguna.

### 3\. Membuat Context

Cara membuat context menggunakan `context.Background()` dan `context.TODO()` sebagai root context.

### 4\. Parent dan Child Context

Bagaimana parent dan child context bekerja, serta bagaimana pewarisan context dalam eksekusi goroutine.

### 5\. Context With Value

Bagaimana cara menyisipkan nilai ke dalam context agar bisa diakses oleh fungsi lain.

### 6\. Context With Cancel

Menggunakan `context.WithCancel()` untuk menghentikan eksekusi goroutine dengan sinyal pembatalan.

### 7\. Context With Timeout

Menggunakan `context.WithTimeout()` untuk membatasi waktu eksekusi goroutine agar tidak berjalan lebih lama dari yang diinginkan.

### 8\. Context With Deadline

Menggunakan `context.WithDeadline()` untuk menghentikan eksekusi pada waktu tertentu yang telah ditentukan.

* * *

## Cara Menjalankan

Pastikan Anda memiliki **Golang** terinstal di sistem Anda. Kemudian jalankan perintah berikut untuk menjalankan setiap bagian:

```sh
go test ./...
```

Atau jalankan test spesifik:

```sh
go test -v -run TestNamaFungsi
```

Semoga bermanfaat! 🚀