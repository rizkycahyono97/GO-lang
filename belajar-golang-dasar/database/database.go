package database

var connection string

// seperti constructor, akan dijalankan pertama kali
func init() {
	connection = "MySql"
}

// seperti public function
func GetDatabase() string {
	return connection
}