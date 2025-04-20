package simple

import "fmt"

//==============================//
// clean up func
//==============================//

type File struct {
	Name string
}

func (f *File) Close() {
	fmt.Println("Closing File ", f.Name)
}

// provider with clean up func
func NewFile(name string) (*File, func()) {
	file := &File{Name: name}

	return file, func() {
		file.Close()
	}
}
