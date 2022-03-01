package test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/ydhnwb/golang_heroku/entity"
	"github.com/ydhnwb/golang_heroku/repo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func openDBConn() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s port=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		"localhost", "handiism", "mrongoz", "5432", "go_heroku_test",
	)

	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(entity.User{}, entity.Product{})
	return db
}

func closeDBConn(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.Close()
}

func TestFaker(*testing.T) {
	db := openDBConn()
	defer closeDBConn(db)

	userRepo := repo.NewUserRepo(db)
	productRepo := repo.NewProductRepo(db)

	for i := 0; i < 100; i++ {
		user := entity.User{}

		if err := faker.FakeData(&user); err != nil {
			panic(err)
		}

		userRepo.InsertUser(user)
	}

	for i := 0; i < 100; i++ {
		product := entity.Product{}

		if err := faker.FakeData(&product); err != nil {
			panic(err)
		}

		product.Price = rand.Uint64() + uint64(10000)
		product.UserID = rand.Int63n(100) + 1

		productRepo.InsertProduct(product)
	}
}
