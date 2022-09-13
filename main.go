package main

import (
	"mini-project/config"
	"mini-project/routes"
)

func main() {

	db := config.ConnectDatabase()

	r := routes.SetupRouter(db)
	r.Run(":8080")

	// v1.GET("/", bookHandler.RootHandler)

	// v1.GET("/hello", bookHandler.HelloHandler)

	// v1.GET("/books/:id", bookHandler.BooksHandler)

	// v1.GET("/query/", bookHandler.QueryHandler)

	// bookRequest := book.BooksRequest{
	// 	Title: "Only You",
	// 	Price: "25000",
	// }
	// bookService.Create(bookRequest)

	// book := book.Book{
	// 	Title:       "3 Lagkah Maju",
	// 	Price:       185000,
	// 	Rating:      5,
	// 	Description: "kiat menjadi sukses sejak muda",
	// }
	// bookRepository.Craete(book)

	// books, err := bookRepository.FindAll()

	// for _, book := range books {
	// 	fmt.Println("Title: ", book.Title)
	// }
	// book, err := bookRepository.FindID(1)
	// fmt.Println("Title: ", book.Title)

	//CRUD
	//INSERT
	// book := book.Book{}
	// book.Title = "3 Lagkah Maju"
	// book.Price = 185000
	// book.Rating = 5
	// book.Description = "kiat menjadi sukses sejak muda"

	// err = db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println("==========")
	// 	fmt.Println("!!Error!!!")
	// 	fmt.Println("==========")
	// }
	// var book book.Book

	//READ
	// var books []book.Book

	// err = db.Debug().Find(&books).Error
	// if err != nil {
	// 	fmt.Println("==========")
	// 	fmt.Println("!!Error!!!")
	// 	fmt.Println("==========")
	// }
	// for _, b := range books {
	// 	fmt.Println("Title:", b.Title)
	// 	fmt.Println("book object %v", b)
	// }

	//UPDATE``
	// var book book.Book
	// err = db.Debug().Where("id=?", 1).First(&book).Error
	// if err != nil {
	// 	fmt.Println("==========")
	// 	fmt.Println("!!Error!!!")
	// 	fmt.Println("==========")
	// }
	// book.Title = "The Last Word Ed 3"
	// err = db.Save(&book).Error
	// if err != nil {
	// 	fmt.Println("==========")
	// 	fmt.Println("!!Error!!!")
	// 	fmt.Println("==========")
	// }

	//DELETE
	// var book book.Book
	// err = db.Debug().Where("id=?", 2).First(&book).Error
	// if err != nil {
	// 	fmt.Println("==========")
	// 	fmt.Println("!!Error!!!")
	// 	fmt.Println("==========")
	// }
	// err = db.Delete(&book).Error
	// if err != nil {
	// 	fmt.Println("==========")
	// 	fmt.Println("!!Error!!!")
	// 	fmt.Println("==========")
	// }

}
