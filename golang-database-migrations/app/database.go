package app

import (
	"database/sql"
	"golang-restful-api/helper"
	"time"
)

func NewDb() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/golang_restful_api_migrations")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db

	// CATATAN GOLANG DATABASE MIGRATIONS

	// untuk create miigration
	// migrate create -ext sql -dir db/migrations create_table_first
	// migrate create -ext sql -dir db/migrations create_table_second
	// migrate create -ext sql -dir db/migrations create_table_third

	// untuk menjalankan migration
	// migrate -database "mysql://root:root@tcp(localhost:3306)/golang_restful_api_migrations" -path db/migrations up

	// untuk rollback
	// migrate -database "mysql://root:root@tcp(localhost:3306)/golang_restful_api_migrations" -path db/migrations down

	// untuk menjalankan dan rollback berdasarkan jumlah angka
	// migrate -database "mysql://root:root@tcp(localhost:3306)/golang_restful_api_migrations" -path db/migrations up 2
	// migrate -database "mysql://root:root@tcp(localhost:3306)/golang_restful_api_migrations" -path db/migrations down 3

	// version migrate
	// migrate -database "mysql://root:root@tcp(localhost:3306)/golang_restful_api_migrations" -path db/migrations version

	// dirty state (jika query mysql salah)
	// force untuk rollback version
	// migrate -database "mysql://root:root@tcp(localhost:3306)/golang_restful_api_migrations" -path db/migrations force 20250522122711
}
