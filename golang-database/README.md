# Panduan Belajar Database di Golang

## 📌 Pendahuluan

Panduan ini mencakup dasar-dasar penggunaan database dalam Golang menggunakan package `database/sql`.

## 📚 Materi yang Dibahas

### 1️⃣ **Pengenalan Package Database**

Memahami penggunaan package `database/sql` untuk mengelola koneksi dan eksekusi perintah SQL dalam Golang.

### 2️⃣ **Menambah Database Driver**

Menggunakan driver MySQL atau PostgreSQL agar Golang dapat berkomunikasi dengan database.

### 3️⃣ **Membuka Koneksi ke Database**

Cara menghubungkan aplikasi Golang dengan database menggunakan `sql.Open()`.

### 4️⃣ **Database Pooling**

Konfigurasi koneksi agar lebih efisien dengan pengaturan pool koneksi seperti `SetMaxOpenConns()`, `SetMaxIdleConns()`.

### 5️⃣ **Eksekusi Perintah SQL**

Cara menjalankan perintah SQL seperti `INSERT`, `UPDATE`, `DELETE` menggunakan `Exec()`.

### 6️⃣ **Query SQL**

Mengambil data dari database menggunakan `Query()` dan `QueryRow()`.

### 7️⃣ **Tipe Data Column**

Memahami bagaimana Golang menangani berbagai tipe data dalam database.

### 8️⃣ **SQL Injection**

Menjelaskan risiko SQL Injection dan bagaimana mengamankan query dalam Golang.

### 9️⃣ **SQL Dengan Parameter**

Menggunakan parameter dalam query untuk mencegah SQL Injection.

### 🔟 **Auto Increment**

Mengelola kolom auto-increment dalam tabel database.

### 1️⃣1️⃣ **Prepare Statement**

Menggunakan `Prepare()` untuk mengoptimalkan eksekusi query berulang kali.

### 1️⃣2️⃣ **Database Transaction**

Menggunakan transaksi database (`BEGIN`, `COMMIT`, `ROLLBACK`) untuk memastikan integritas data.

### 1️⃣3️⃣ **Repository Pattern**

Menerapkan pola repository untuk memisahkan logika akses database dari logika bisnis.

## 🔗 Referensi

* *   [Golang Database Documentation](https://pkg.go.dev/database/sql)
* *   [MySQL Golang Driver](https://github.com/go-sql-driver/mysql)

* * *

Selamat belajar! 🚀