# 📘 Golang HttpRouter Tutorial

Repositori ini berisi materi pembelajaran tentang penggunaan **HttpRouter** di bahasa pemrograman Go (Golang). HttpRouter adalah salah satu router tercepat dan ringan yang cocok digunakan untuk membangun aplikasi web dengan performa tinggi.

* * *

## Materi

### ✅ Pendahuluan

Pengenalan awal mengenai HttpRouter, alasan penggunaannya, serta bagaimana router ini berbeda dengan `net/http` standar.

### 🌐 Pengenalan HttpRouter

Menjelaskan dasar-dasar penggunaan `github.com/julienschmidt/httprouter`, termasuk struktur dan fitur-fiturnya yang efisien.

### 🧭 Router

Cara membuat router baru, mengatur route, dan menambahkan handler ke dalam aplikasi menggunakan HttpRouter.

### 📦 Params

Menangani parameter dinamis pada URL, seperti `/user/:id`, dan cara mengambil nilai parameter menggunakan `httprouter.Params`.

### 🧩 Router Pattern

Membahas pola rute (route pattern) yang didukung oleh HttpRouter, termasuk wildcard (`*`) dan parameter (`:`).

### 🗂️ Serve File

Cara melayani file statis seperti gambar, CSS, dan JavaScript dari direktori publik menggunakan router.

### 💥 Panic Handler

Penanganan error tidak terduga (panic) dengan `PanicHandler`, untuk mencegah aplikasi crash dan memberikan feedback yang aman.

### 🚫 Not Found Handler

Custom handler untuk menangani rute yang tidak ditemukan (404 Not Found), agar user experience lebih baik.

### ⛔ Method Not Allowed Handler

Menangani permintaan dengan metode HTTP yang tidak diizinkan pada suatu rute, seperti `POST` pada route yang hanya menerima `GET`.

### 🔄 Middleware

Penerapan middleware di HttpRouter untuk logging, autentikasi, dan fitur lainnya yang dapat membungkus route handler.

* * *

## 🔧 Teknologi

*    Golang    

*    HttpRouter (`github.com/julienschmidt/httprouter`)    
 
*    Unit testing dengan `net/http/httptest` dan `stretchr/testify`

* * *

## 🧪 Testing

Semua handler dan middleware diuji menggunakan pendekatan unit test untuk memastikan kestabilan dan validitas output.