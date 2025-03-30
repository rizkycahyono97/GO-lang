package repository

import (
	"context"
	"fmt"
	golangdatabase "golang-database" // alias ke package golang-database
	"golang-database/entity"
	"testing"
)

// test for repository insert
func TestCommentRepository(t *testing.T) {
	commentRepository := NewCommentRepository(golangdatabase.GetConnection())	// koneksi database

	ctx := context.Background()
	comment := entity. Comment {
		Email: "repository2@gmail.com",
		Comment: "repository Comment 2",
	}

	result, err := commentRepository.Insert(ctx, comment) 
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}



// test for repository FindById
func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(golangdatabase.GetConnection())	// koneksi database

	result, err := commentRepository.FindById(context.Background(), 65)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}


// test for repository FindAll
func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(golangdatabase.GetConnection())	// koneksi database

	comments, err := commentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}
	
	for _, comment := range comments {
		fmt.Println(comment)
	}
}