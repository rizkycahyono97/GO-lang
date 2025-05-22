# Golang Database Migrations

Repository ini menggunakan **`golang-migrate`**, sebuah alat untuk mengelola migrasi database dalam proyek Go. Berikut adalah panduan untuk membuat, menjalankan, dan mengelola migrasi database.

---

## **Prasyarat**

1. **Instalasi `golang-migrate`**
   Pastikan `golang-migrate` sudah terinstal di sistem Anda. Jika belum, instal dengan perintah berikut:
   ```bash
   go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
   ```
   Atau unduh binary langsung dari [GitHub Releases](https://github.com/golang-migrate/migrate/releases).

2. **Database MySQL**
   Pastikan database MySQL sudah berjalan dan dapat diakses. Konfigurasi koneksi database digunakan dalam perintah migrasi.

3. **Direktori Migrasi**
   Direktori `db/migrations` digunakan untuk menyimpan file migrasi. Pastikan direktori ini ada atau buat dengan perintah:
   ```bash
   mkdir -p db/migrations
   ```

---

## **Command Dasar**

### **1. Membuat File Migrasi Baru**
Untuk membuat file migrasi baru, gunakan perintah berikut:
```bash
migrate create -ext sql -dir db/migrations <nama_migrasi>
```

- **`-ext sql`**: Menentukan ekstensi file migrasi (`.sql`).
- **`-dir db/migrations`**: Menentukan lokasi penyimpanan file migrasi.
- **`<nama_migrasi>`**: Nama deskriptif untuk migrasi (misalnya, `create_table_users`).

Contoh:
```bash
migrate create -ext sql -dir db/migrations create_table_users
```

Ini akan membuat dua file baru di direktori `db/migrations`:
- `YYYYMMDDHHMMSS_create_table_users.up.sql` (untuk migrasi naik)
- `YYYYMMDDHHMMSS_create_table_users.down.sql` (untuk rollback)

---

### **2. Menjalankan Migrasi**
Untuk menjalankan semua migrasi yang belum diterapkan, gunakan perintah berikut:
```bash
migrate -database "<koneksi_database>" -path db/migrations up
```

- **`<koneksi_database>`**: URL koneksi ke database. Contoh untuk MySQL:
  ```
  mysql://root:root@tcp(localhost:3306)/golang_restful_api_migrations
  ```

Contoh:
```bash
migrate -database "mysql://root:root@tcp(localhost:3306)/golang_restful_api_migrations" -path db/migrations up
```

---

### **3. Rollback Migrasi**
Untuk rollback migrasi, gunakan perintah berikut:
```bash
migrate -database "<koneksi_database>" -path db/migrations down
```

Ini akan menjalankan file `.down.sql` dari migrasi terakhir.

Contoh:
```bash
migrate -database "mysql://root:root@tcp(localhost:3306)/golang_restful_api_migrations" -path db/migrations down
```

---

### **4. Menjalankan/Rollback Berdasarkan Jumlah Langkah**
Anda dapat menjalankan atau rollback migrasi berdasarkan jumlah langkah tertentu.

- **Menjalankan migrasi sebanyak `n` langkah:**
  ```bash
  migrate -database "<koneksi_database>" -path db/migrations up <jumlah_langkah>
  ```

- **Rollback migrasi sebanyak `n` langkah:**
  ```bash
  migrate -database "<koneksi_database>" -path db/migrations down <jumlah_langkah>
  ```

Contoh:
```bash
migrate -database "mysql://root:root@tcp(localhost:3306)/golang_restful_api_migrations" -path db/migrations up 2
```

---

### **5. Memeriksa Versi Migrasi**
Untuk memeriksa versi migrasi saat ini, gunakan perintah berikut:
```bash
migrate -database "<koneksi_database>" -path db/migrations version
```

Contoh:
```bash
migrate -database "mysql://root:root@tcp(localhost:3306)/golang_restful_api_migrations" -path db/migrations version
```

---

### **6. Menangani Dirty State**
Jika migrasi gagal dan database dalam keadaan "dirty", Anda perlu memperbaiki status migrasi secara manual menggunakan perintah `force`.

- **Memaksa ke versi tertentu:**
  ```bash
  migrate -database "<koneksi_database>" -path db/migrations force <versi_target>
  ```

Contoh:
```bash
migrate -database "mysql://root:root@tcp(localhost:3306)/golang_restful_api_migrations" -path db/migrations force 20250522122711
```

Setelah itu, jalankan migrasi kembali.

---

## **Struktur File Migrasi**

File migrasi terdiri dari dua bagian:
1. **`up.sql`**: Digunakan untuk menerapkan perubahan ke database (misalnya, membuat tabel baru).
2. **`down.sql`**: Digunakan untuk rollback perubahan (misalnya, menghapus tabel).

Contoh:

**`20231001000000_create_table_users.up.sql`:**
```sql
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

**`20231001000000_create_table_users.down.sql`:**
```sql
DROP TABLE users;
```

---

## **Tips Penggunaan**

1. **Backup Database**
    - Selalu backup database sebelum menjalankan migrasi, terutama di lingkungan produksi.

2. **Uji Migrasi Lokal**
    - Uji migrasi di lingkungan lokal sebelum menerapkannya ke server produksi.

3. **Gunakan Transaksi**
    - Pastikan migrasi Anda menggunakan transaksi jika memungkinkan untuk memastikan rollback otomatis jika terjadi kesalahan.

4. **Nama Migrasi Deskriptif**
    - Gunakan nama migrasi yang jelas dan deskriptif untuk memudahkan pemahaman.

---

## **Referensi**

- Dokumentasi Resmi: [https://github.com/golang-migrate/migrate](https://github.com/golang-migrate/migrate)
- Driver Database: Pastikan driver untuk database Anda (misalnya, MySQL) telah diinstal.

---

Dengan panduan ini, Anda dapat mengelola migrasi database dengan mudah menggunakan `golang-migrate`. Jika ada pertanyaan lebih lanjut, silakan ajukan! 😊