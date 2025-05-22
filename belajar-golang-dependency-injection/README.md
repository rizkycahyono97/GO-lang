* * *

# 📦 Belajar Dependency Injection dengan Google Wire

Dokumentasi ini berisi ringkasan materi pembelajaran mengenai **Dependency Injection (DI)** dengan menggunakan library **Google Wire** dalam pengembangan aplikasi Go, mulai dari pengantar konsep hingga penerapannya dalam RESTful API. Ini adalah lanjutan dari **belajar golang-restful-api**, repo ini merubah **Dependency Injection** yang sebelumnya Manual menjadi menggunakan **google wire**.

## 🧠 Daftar Materi

### 1\. Pendahuluan

Pengenalan umum tentang pentingnya manajemen dependensi dalam pemrograman modern, serta manfaat Dependency Injection dalam menjaga arsitektur kode yang bersih dan mudah diuji.

### 2\. Pengenalan Dependency Injection

Penjelasan konsep dasar dari Dependency Injection, termasuk jenis-jenisnya dan bagaimana prinsip ini membantu dalam memisahkan logika program dan dependensinya.

### 3\. Library Dependency Injection

Pengenalan terhadap berbagai library Dependency Injection, dan fokus pada pemanfaatan **Google Wire** sebagai tool utama.

### 4\. Membuat Project

Langkah-langkah awal dalam membuat proyek Go untuk mendemonstrasikan penggunaan Google Wire.

### 5\. Menginstall Google Wire

Panduan instalasi library Google Wire ke dalam project Go menggunakan `go get` dan setup awal.

### 6\. Provider

Penjelasan tentang fungsi provider dalam Google Wire untuk mengelola dan menyuntikkan dependensi secara otomatis.

### 7\. Injector

Mengenal injector yang digunakan untuk menginisialisasi dan menggabungkan berbagai provider menjadi satu struktur yang utuh.

### 8\. Dependency Injection

Implementasi nyata teknik Dependency Injection dalam berbagai skenario aplikasi menggunakan Google Wire.

### 9\. Error

Pembahasan penanganan error dalam proses dependency injection dan bagaimana Google Wire mengelola kesalahan.

### 10\. Injector Parameter

Cara memberikan parameter pada injector agar lebih fleksibel dan sesuai kebutuhan pada saat runtime.

### 11\. Multiple Binding

Menangani situasi di mana terdapat lebih dari satu implementasi dari interface yang sama, dan bagaimana menyelesaikannya dengan multiple binding.

### 12\. Provider Set

Membuat set dari beberapa provider untuk mempermudah pengelompokan dan penggabungan dependensi.

### 13\. Binding Interface

Menghubungkan interface dengan implementasi konkret melalui konfigurasi di Google Wire.

### 14\. Struct Provider

Menyediakan dependensi berupa struct menggunakan Wire, serta praktik terbaik dalam penggunaannya.

### 15\. Binding Values

Teknik untuk menyuntikkan nilai tetap (seperti konfigurasi atau konstanta) sebagai dependensi.

### 16\. Struct Field Provider

Cara menyediakan dependensi langsung ke field dalam struct secara otomatis.

### 17\. Cleanup Function

Mengelola resource seperti koneksi database atau file dengan cleanup function dalam Google Wire.

### 18\. Dependency Injection di RESTful API

Contoh implementasi Dependency Injection dalam konteks RESTful API menggunakan framework atau library yang relevan dalam ekosistem Go.

* * *

> 🛠️ Dokumentasi ini cocok untuk kamu yang ingin memahami konsep Dependency Injection dan bagaimana menerapkannya dengan Google Wire di proyek Go yang skalabel dan maintainable.
