package test

import (
	"fmt"
	"testing"

	"github.com/ydhnwb/golang_heroku/repo"
)

func TestPreload(t *testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	repo := repo.NewProductRepo(db)

	product, err := repo.FindOneProductByID("1")

	if err != nil {
		panic(err)
	}

	fmt.Println(product.User)
}
