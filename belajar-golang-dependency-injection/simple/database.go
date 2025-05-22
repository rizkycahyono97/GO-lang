package simple

//======================//
//multiple binding
//======================//
//banyak parameter dengan tipe data yang sama di provider

type Database struct {
	Name string
}

// alias = harus membikin alias untuk inisialisasi parameter yang sama
type DatabaseMongoDB Database
type DatabasePostgreSQL Database

func NewDatabaseMongoDB() *DatabaseMongoDB {
	return (*DatabaseMongoDB)(&Database{Name: "MongoDB"}) //convert
}

func NewDatabasePostgreSQL() *DatabasePostgreSQL {
	return (*DatabasePostgreSQL)(&Database{Name: "PostgreSQL"})
}

type DatabaseRepository struct {
	DatabasePostgreSQL *DatabasePostgreSQL
	DatabaseMongoDB    *DatabaseMongoDB
}

func NewDatabaseRepository(databasePostgreSQL *DatabasePostgreSQL, databaseMongoDB *DatabaseMongoDB) *DatabaseRepository {
	return &DatabaseRepository{DatabasePostgreSQL: databasePostgreSQL, DatabaseMongoDB: databaseMongoDB}
}
