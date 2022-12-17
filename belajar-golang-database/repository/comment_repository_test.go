package repository

import (
	belajargolangdatabase "belajar-golang-database"
	"belajar-golang-database/entity"
	"context"
	"fmt"
	"testing"
)

func TestCommentRespositoryInsert(t *testing.T) {
	db := belajargolangdatabase.GetConnection()
	ctx := context.Background()

	repository := NewCommentRepository(db)
	comment := entity.Comment{
		Email:   "rizki@repository.com",
		Comment: "Hello Dari repository",
	}
	result, err := repository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestCommentRespositoryFindById(t *testing.T) {
	ctx := context.Background()
	repository := NewCommentRepository(belajargolangdatabase.GetConnection())
	comment, err := repository.FindById(ctx, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(comment)
}

func TestCommentRespositoryFindAll(t *testing.T) {
	ctx := context.Background()
	comments, err := NewCommentRepository(belajargolangdatabase.GetConnection()).FindAll(ctx)
	if err != nil {
		panic(err)
	}
	for _, comment := range comments {
		fmt.Println(comment)
	}
}
