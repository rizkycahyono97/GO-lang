package golangdatabase

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

// ================== //
// exec SQL
// ================== //
func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()	//context

	query := "INSERT INTO customer(id, name) VALUES('putra', 'putra')"	// INSERT data
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert Data")
}


// ======================= //
// query sql
// ======================= //
func TestSelectSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()	

	query := "SELECT id, name, email FROM customer"	// SELECT data
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	fmt.Println("Success Select Data")

	// menampilkan data, karena datanya masih disimpan di objeck rows
	for rows.Next() {
		var id, name, email string		// field harus sama seperti yang ada di table

		// cek error
		err := rows.Scan(&id, &name, &email)
		if err != nil {
			panic(err)
		}

		fmt.Println("ID: ", id)
		fmt.Println("Name: ", name)
		fmt.Println("Email: ", email)
	}
}


// ================== //
// tipe data sql di go
// ================== //
func TestQueryComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()	

	query := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"	// SELECT data
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	fmt.Println("Success Select Data")

	// menampilkan data, karena datanya masih disimpan di objeck rows
	for rows.Next() {
		// field / tipe data
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64 
		var createdAt time.Time
		var birthDate sql.NullTime
		var married bool

		// cek error
		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}

		fmt.Println("=========================")
		fmt.Println("ID: ", id)
		fmt.Println("Name: ", name)

		// jika null
		if email.Valid {
			fmt.Println("Email: ", email.String)
		}

		fmt.Println("Balance: ", balance)
		fmt.Println("Rating: ", rating)
		fmt.Println("Birth Date: ", birthDate)
		fmt.Println("Created At: ", createdAt)
		fmt.Println("Married: ", married)

	}
}



// ==================== //
// sql injection
// ==================== //
func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()	

	// inputan user
	username := "admin'; #"		// sql injection
	password := "salah"

	query := "SELECT username FROM user WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1" // vulnerabilty query
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login ", username)
	} else {
		fmt.Println("Gagal Login")
	}
}

// sqlInjection Safe
func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()	

	// inputan user
	username := "admin'; #"		// sql injection
	password := "salah"

	query := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1" // safe query
	rows, err := db.QueryContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login ", username)
	} else {
		fmt.Println("Gagal Login")
	}
}

func TestExecSqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()	//context

	// inputan user
	username := "rizky'; DROP TABLE user; #"
	password := "rizky"

	query := "INSERT INTO user(username, password) VALUES(?, ?)"	// safe insert data (meskipun inputan user berupa sql injection tetap aman, akan diinputankan menjadi string)
	_, err := db.ExecContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert Data")
}


// ======================== //
// auto increment 
// ======================== //
func TestAutoIncement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()	//context

	// inputan user
	email := "putra@mail.com"
	comment := "putra"

	query := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	result, err := db.ExecContext(ctx, query, email, comment)
	if err != nil {
		panic(err)
	}

	lastId, err := result.LastInsertId()	// lastInsertId jika menggunakan AUTO INCREMENT
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert Data with Id: ", lastId)
}


// ====================== //
// prepare statement = untuk eksekusi berulang dg parameter berbeda tanpa compile lagi
// ====================== //
func TestPrepareStatement(t *testing.T) {
	db :=  GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	statement, err := db.PrepareContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 0; i < 10; i++ {
		
		email := "rizky" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke: " + strconv.Itoa(i)
		
		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment Id ke: ", id)
	}
}


// ==================== //
// database tramsaction
// ==================== //
func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	query := "INSERT INTO comments(email, comment) VALUES(?, ?)"

	// do transaction
	for i := 0; i < 10; i++ {
		
		email := "cahyono" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke: " + strconv.Itoa(i)
		
		result, err := tx.ExecContext(ctx, query, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment Id ke: ", id)
	}

	// err = tx.Commit()	// untuk commit transaction
	err = tx.Rollback()		// untuk rollback transaction
	if err != nil {
		panic(err)
	}
}