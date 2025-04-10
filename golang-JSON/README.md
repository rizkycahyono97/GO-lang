# 📘 Golang JSON Package

Repositori ini berisi materi dan contoh penggunaan package `encoding/json` di Golang, untuk kebutuhan encoding dan decoding data JSON. Cocok bagi kamu yang ingin memahami dasar hingga teknik lanjutan pemrosesan JSON di Go.

## 📚 Materi

### Pendahuluan

Penjelasan mengenai tujuan dan pentingnya memahami JSON dalam pengembangan aplikasi, khususnya untuk komunikasi API atau penyimpanan data.

### Pengenalan Package json

Mengenal package `encoding/json`, fungsi-fungsi utamanya, dan struktur data yang kompatibel.

### Encode JSON

Belajar mengubah data struct ke format JSON menggunakan `json.Marshal` dan `json.NewEncoder`.

### JSON Object

Cara membentuk dan meng-encode object JSON dari struct, termasuk nested struct.

### Decode JSON

Belajar membaca data JSON dan mengubahnya menjadi struct dengan `json.Unmarshal` dan `json.NewDecoder`.

### JSON Array

Cara menangani JSON array dan decode ke dalam slice of struct.

### JSON Tag

Penggunaan tag `json` dalam struct untuk mengatur nama field, pengabaian field (`-`), dan lainnya.

### Map

Penggunaan map sebagai alternatif struct untuk struktur JSON dinamis atau fleksibel.

### Streaming Decoder

Teknik membaca data JSON besar atau terus-menerus (streaming) menggunakan `json.NewDecoder`.

### Streaming Encoder

Teknik menulis/encode data JSON besar secara streaming ke output seperti file atau koneksi HTTP.

### Materi Selanjutnya

Penutupan materi dan gambaran tentang kelanjutan pembelajaran, seperti manipulasi data kompleks atau integrasi dengan API.