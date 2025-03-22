# Golang Standard Library

Proyek ini berisi contoh penggunaan berbagai pustaka standar di Golang. Setiap file berisi implementasi kode terkait dengan pustaka tertentu, yang berguna untuk memahami konsep dasar dalam pengembangan aplikasi dengan Go.

## struktur folder

```
.
├── base64.go          # Encoding dan decoding Base64
├── bufio_read.go      # Membaca input dengan bufio
├── bufio_write.go     # Menulis output dengan bufio
├── csv_reader.go      # Membaca file CSV
├── csv_writer.go      # Menulis file CSV
├── duration.go        # Operasi durasi dan waktu
├── errors.go          # Manajemen error di Go
├── file.go            # Operasi file dasar
├── filepath.go        # Manipulasi path file
├── flag.go            # Parsing flag dari command-line
├── fmt.go             # Format string dan output
├── go.mod             # File modul Go
├── list.go            # Struktur data list dari container/list
├── math.go            # Operasi matematika
├── os.go              # Interaksi dengan sistem operasi
├── path.go            # Manipulasi path
├── README.md          # Dokumentasi proyek ini
├── reflect.go         # Reflection dalam Go
├── regexp.go          # Ekspresi reguler
├── ring.go            # Struktur data circular list dari container/ring
├── sample.log         # Contoh file log
├── slices.go          # Operasi slice
├── sort.go            # Sorting data
├── strconv.go         # Konversi string dan angka
├── strings.go         # Manipulasi string
└── time.go            # Operasi tanggal dan waktu
```

## Cara Penggunaan

Untuk menjalankan salah satu contoh, gunakan perintah berikut:

```sh
 go run nama_file.go
```

Contoh:

```sh
 go run base64.go
```