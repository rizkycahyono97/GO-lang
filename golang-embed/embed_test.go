package golangembed

import (
	"embed"
	// _ "embed"
	"fmt"
	"os"
	"testing"
)

// ========================= //
// embed file
// ========================= //

//go:embed version.txt
var version string

//go:embed version.txt
var version2 string

func TestEmbed(t *testing.T) {
	fmt.Println(version)
	fmt.Println(version2)
}


// ==================== //
// embed file byte
// ==================== //

//go:embed personal.png
var gambar []byte

func TestByte(t *testing.T) {
	err := os.WriteFile("personal_new.png", gambar, os.ModePerm)	// akan menyimpan gambar baru
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Copy Picture")
}


// ===================== //
// embed multiple file
// ===================== //

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

func TestMultipleFiles(t *testing.T) {
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a))

	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b))

	c, _ := files.ReadFile("files/c.txt")
	fmt.Println(string(c))
}



// ======================= //
// path matcher = memilih file berdasarkan path/directory
// ======================= //

//go:embed files/*txt
var path embed.FS

func TestPathMatcher(t *testing.T) {
	dirEntries, _ := path.ReadDir("files")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())	// print nama file
			file, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}


