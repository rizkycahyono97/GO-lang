# Golang Embed

Golang menyediakan package `embed` yang memungkinkan kita untuk menyertakan file statis langsung ke dalam binary aplikasi. Fitur ini sangat berguna untuk mengemas file konfigurasi, template, gambar, atau asset lainnya tanpa perlu dependensi eksternal.

## Fitur Utama

* *   **Embed File ke String**: Memasukkan file teks langsung ke dalam variabel string.
*     
* *   **Embed File ke Byte**: Digunakan untuk file biner seperti gambar atau video.
*     
* *   **Embed Multiple Files**: Mengelola beberapa file sekaligus dalam satu variabel.
*     
* *   **Path Matcher**: Memungkinkan seleksi file berdasarkan pola tertentu.
*     
* *   **Hasil Embed di Compile**: Semua file yang di-embed akan disertakan dalam binary, sehingga tidak diperlukan file tambahan saat runtime.
*     

## Kesimpulan

Dengan package `embed`, pengelolaan aset dalam aplikasi menjadi lebih sederhana dan mandiri, karena semua file dapat dikemas langsung dalam satu binary tanpa dependensi tambahan.