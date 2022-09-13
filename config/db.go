package config

import (
	"log"
	"mini-project/book"
	"mini-project/buy"
	"mini-project/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	dsn := "host=localhost user=postgres password=SemangatDev dbname=mini-project port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection error")
	}
	// fmt.Println("Data connection succeed")
	db.AutoMigrate(&book.Book{}, &user.User{}, &buy.Buy{})

	return db
}
