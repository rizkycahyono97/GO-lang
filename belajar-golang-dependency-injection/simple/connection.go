package simple

import "fmt"

//==============================//
// clean up func
//==============================//

type Connection struct {
	*File
}

func (conn *Connection) Close() {
	fmt.Println("Closing Connection ", conn.Name)
}

// provider with clean up func
func NewConnection(file *File) (*Connection, func()) {
	conn := &Connection{File: file}

	return conn, func() {
		conn.Close()
	}
}
