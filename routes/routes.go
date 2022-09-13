package routes

import (
	"mini-project/book"
	"mini-project/buy"
	"mini-project/handler"
	"mini-project/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	buyRepository := buy.NewRepository(db)
	buyService := buy.NewService(buyRepository)
	buyHandler := handler.NewBuyHandler(buyService)

	v1 := router.Group("/v1")
	v1.POST("/books", bookHandler.CreateBook)
	v1.GET("/books", bookHandler.GetBooks)
	v1.GET("/books/:id", bookHandler.GetBook)
	v1.PUT("/books/:id", bookHandler.UpdateBook)
	v1.DELETE("/books/:id", bookHandler.DeleteBook)

	v1.POST("/users", userHandler.CreateUser)
	v1.GET("/users", userHandler.GetUsers)
	v1.GET("/users/:id", userHandler.GetUser)
	v1.PUT("/users/:id", userHandler.UpdateUser)
	v1.DELETE("/users/:id", userHandler.DeleteUser)

	v1.POST("/buys", buyHandler.CreateBuy)
	v1.GET("/buys", buyHandler.GetBuys)
	v1.GET("/buys/:id", buyHandler.GetBuy)
	v1.PUT("/buys/:id", buyHandler.UpdateBuy)
	v1.DELETE("/buys/:id", buyHandler.DeleteBuy)
	v1.POST("/pembelians", buyHandler.GetPembelian)

	return router
}
