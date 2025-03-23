package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// table benhmark
func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name string
		request string
	}{
		{
			name: "Rizky",
			request: "Rizky",
		},
		{
			name: "Cahyono",
			request: "Cahyono",
		},
		{
			name: "RizkyCahyonoPutra",
			request: "Rizky Cahyono Putra",
		},{
			name: "DwiAgus",
			request: "Dwi Agus",
		},
	}

	for _, bencmark := range benchmarks {
		b.Run(bencmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(bencmark.request)
			}
		})
	}
}

// sub bencmark
func BenchmarkSub(b *testing.B) {
	b.Run("Rizky", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Rizky")
		}
	})

	b.Run("Cahyono", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Cahyono")
		}
	})
}

// bencmark = test kecepatan
func BenchmarkHelloWorld(t *testing.B) {
	for i := 0; i < t.N; i++ {	// akan looping terus dari GOnya
		HelloWorld("Rizky")
	}
}

func BenchmarkHelloWorldCahyono(t *testing.B) {
	for i := 0; i < t.N; i++ {
		HelloWorld("Cahyono")
	}
}

// Main test
func TestMain(m *testing.M) {
	// BEFORE 
	fmt.Println("Before Testing")

	m.Run()

	// AFTER
	fmt.Println("After Testing")
}

func TestSubTest(t *testing.T) {
	t.Run("Rizky", func(t *testing.T) { 	// "Rizky" -> nama subtest
		result := HelloWorld("Rizky")
		assert.Equal(t, "Hello Rizky", result, "Result must be 'Hello Rizky'")	
	})

	t.Run("Cahyono", func(t *testing.T) {
		result := HelloWorld("Cahyono")
		assert.Equal(t, "Hello Cahyono", result, "Result must be 'Hello Cahyono'")	
	})
}

func TestHelloWorldRizky(t *testing.T) {
	result := HelloWorld("Rizky")
	if result != "Hello Rizky" {

		// error
		// t.Fail()	// akan tetap eksekusi kode setelahnya
		t.Error("Result must be 'Hello Rizky'")		// akan tetap eksekusi kode setelahnya
	}

	fmt.Println("TestHelloWorldRizky Done!")
}

func TestSkip(t *testing.T) {

	// akan skip jika OSnya linux
	if runtime.GOOS == "linux" {
		t.Skip("Can't Test on Linux")
	}

	result := HelloWorld("Rizky")
	require.Equal(t, "Hello Rizky", result, "Result must be 'Hello Rizky'")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Rizky")
	require.Equal(t, "Hello Rizky", result, "Result must be 'Hello Rizky'")	// require = seperti FailNow()

	fmt.Println("TestHelloWorldRequire Done!")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Rizky")
	assert.Equal(t, "Hello Rizky", result, "Result must be 'Hello Rizky'")	// assert = seperti Fail()

	fmt.Println("TestHelloWorldAssert Done!")
}

func TestHelloWorldCahyono(t *testing.T) {
	result := HelloWorld("Rizky")
	if result != "Hello Rizky" {

		// error
		// t.FailNow()		// tidak akan eksekusi kode program setelahnya
		t.Fatal("Result must be 'Hello Cahyono'")  // tidak akan eksekusi kode program setelahnya
	}

	fmt.Println("TestHelloWorldCahyono Done!")
}

// table test
func TestTableHelloWorld(t *testing.T) {

	// slice
	tests := []struct {
		name string
		request string
		expected string
	}{
		{
			name: "Rizky",
			request: "Rizky",
			expected: "Hello Rizky",
		},
		{
			name: "Cahyono",
			request: "Cahyono",
			expected: "Hello Cahyono",
		},
		{
			name: "Putra",
			request: "Putra",
			expected: "Hello Putra",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}