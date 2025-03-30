package golangdatabase

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)	

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {

	// connection
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/golang-database")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	// gunakan db
}