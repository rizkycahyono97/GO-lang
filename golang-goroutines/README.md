# Golang Goroutines

Repository ini berisi berbagai contoh implementasi **Goroutine** di Golang. Setiap file dalam repository ini merupakan pengujian (`_test.go`) yang menunjukkan berbagai teknik concurrency dalam Golang.

## Struktur Direktori

```
.
├── atomic_test.go           # Contoh penggunaan atomic untuk operasi thread-safe
├── channel_test.go          # Implementasi channel untuk komunikasi antar goroutine
├── cond_test.go             # Penggunaan sync.Cond untuk sinkronisasi goroutine
├── gomaxprocs_test.go       # Pengaturan GOMAXPROCS untuk mengatur jumlah thread
├── go.mod                   # File module Go
├── goroutine_test.go        # Dasar penggunaan goroutine di Golang
├── map_test.go              # Penggunaan sync.Map untuk concurrent map
├── mutex_test.go            # Implementasi sync.Mutex untuk mencegah race condition
├── once_test.go             # Penggunaan sync.Once untuk inisialisasi satu kali
├── pool_test.go             # Pemanfaatan sync.Pool untuk object pooling
├── race_condition_test.go   # Simulasi race condition dalam concurrency
├── ticker_test.go           # Penggunaan time.Ticker untuk eksekusi berkala
├── timer_test.go            # Penggunaan time.Timer untuk eksekusi terjadwal
└── waitgroup_test.go        # Penggunaan sync.WaitGroup untuk menunggu goroutine selesai
```

## Penjelasan Singkat

* *   **Goroutine** digunakan untuk menjalankan fungsi secara asynchronous.
*     
* *   **Channel** membantu komunikasi antar goroutine.
*     
* *   **Sync Package** menyediakan berbagai mekanisme sinkronisasi seperti `Mutex`, `WaitGroup`, `Cond`, dan `Once`.
*     
* *   **Atomic Operations** digunakan untuk operasi thread-safe tanpa lock.
*     
* *   **Ticker & Timer** memungkinkan eksekusi fungsi secara periodik atau terjadwal.
*     
* *   **Race Condition Handling** menunjukkan bagaimana menghindari race condition dengan teknik concurrency yang benar.
*     

## Cara Menjalankan

Pastikan kamu sudah menginstal **Go** dan berada dalam direktori proyek ini, lalu jalankan perintah berikut untuk menjalankan semua test:

```sh
go test ./...
```

Atau jalankan test spesifik:

```sh
go test -v -run TestNamaFungsi
```

## Kontribusi

Silakan buat **pull request** atau **issue** jika ingin menambahkan contoh atau perbaikan dalam repository ini.

* * *

📌 **Catatan**: Repository ini bertujuan sebagai referensi untuk memahami konsep **concurrency** dan **parallelism** dalam Golang.