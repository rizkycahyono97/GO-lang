package golangdatabase

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	// connection
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/golang-database?parseTime=true")	// parseDate = untuk konversi date/time dari bit array

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)	//  jumlah maksimum koneksi idle (tidak aktif) dalam connection pool.
	db.SetMaxOpenConns(100)	// maximal koneksi aktif
	db.SetConnMaxIdleTime(5 * time.Minute)	// Menentukan berapa lama koneksi idle bisa bertahan sebelum ditutup.
	db.SetConnMaxLifetime(60 * time.Minute)	// berapa lama total waktu hidup (lifetime) sebuah koneksi sebelum ditutup dan dibuat ulang.

	return db
}