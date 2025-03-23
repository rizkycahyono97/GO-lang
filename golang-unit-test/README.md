# Panduan Belajar Golang Testing

Dokumen ini berisi ringkasan materi yang dipelajari dalam video tentang **Golang Testing**. Materi ini mencakup dasar-dasar pengujian di Golang, mulai dari unit test hingga benchmarking.

## Daftar Materi

### 1\. Pendahuluan

* *   Pengenalan tentang pentingnya testing dalam pengembangan perangkat lunak.
* *   Manfaat melakukan testing sejak tahap awal pengembangan.

### 2\. Pengenalan Software Testing

* *   Konsep dasar software testing.
* *   Jenis-jenis testing seperti unit test, integration test, dan functional test.

### 3\. Pengenalan Testing Package

* *   Mengenal package `testing` bawaan Golang.
* *   Struktur dasar dalam melakukan pengujian di Golang.

### 4\. Membuat Unit Test

* *   Cara membuat unit test di Golang.
* *   Menulis fungsi pengujian dengan `testing.T`.

### 5\. Menggagalkan Test

* *   Teknik untuk membuat test gagal.
* *   Menggunakan `t.Fail()`, `t.FailNow()`, `t.Fatal()`, dan `t.Errorf()`.

### 6\. Assertion

* *   Menggunakan **assertion** dalam pengujian.
* *   Pengenalan package `testify/assert` untuk memudahkan pengecekan hasil.

### 7\. Skip Test

* *   Bagaimana cara melewati test dalam kondisi tertentu.
* *   Penggunaan `t.Skip()` dan `t.SkipNow()`.

### 8\. Before dan After Test

* *   Menjalankan kode sebelum (`setup`) dan sesudah (`teardown`) pengujian.
* *   Menggunakan `TestMain()` untuk setup dan cleanup test.

### 9\. Sub Test

* *   Cara membuat **sub test** dalam satu fungsi pengujian.
* *   Menggunakan `t.Run()` untuk menjalankan sub test.

### 10\. Table Test

* *   Teknik **table-driven test** untuk menguji banyak skenario dalam satu fungsi.
* *   Memanfaatkan array atau slice untuk menyimpan data uji.

### 11\. Mock 

* *   Pengenalan konsep **mocking** dalam testing.
* *   Menggunakan package `testify/mock` untuk membuat mock object.

### 12\. Benchmark 

* *   Pengenalan benchmark testing di Golang.
* *   Cara mengukur performa fungsi dengan `testing.B`.

### 13\. Membuat Benchmark 

* *   Cara membuat fungsi benchmark di Golang.
* *   Penggunaan `b.N` untuk mengulangi eksekusi hingga mendapatkan hasil optimal.

### 14\. Sub Benchmark 

* *   Membuat sub benchmark untuk menguji variasi kode dalam satu fungsi.
* *   Menggunakan `b.Run()` untuk sub benchmark.

### 15\. Table Benchmark 

* *   Teknik **table-driven benchmark** untuk menguji berbagai skenario dalam satu fungsi.
* *   Memanfaatkan array atau slice dalam pengujian performa.

* * *

## Cara Menjalankan Testing di Golang

1.  Pastikan Anda berada dalam direktori proyek yang berisi kode dan test.
2.  Jalankan perintah berikut untuk menjalankan semua unit test:
    
    ```sh
    go test ./...
    ```
    
3.  Jalankan perintah berikut untuk menjalankan benchmark:
    
    ```sh
    go test -bench .
    ```
    

* * *

Dokumen ini memberikan gambaran singkat tentang konsep-konsep utama dalam **Golang Testing**. Untuk pemahaman lebih lanjut, pastikan untuk mencoba setiap konsep dalam proyek nyata! 🚀

## Doc
https://pkg.go.dev/testing