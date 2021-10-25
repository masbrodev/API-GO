package main

import (
	"fmt"
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	colorBlue := "\033[34m"
	colorRed := "\033[31m"
	colorReset := "\033[0m"

	dsn := "root:@tcp(127.0.0.1:3306)/go_pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(string(colorRed), "DB Connection error")
	}
	fmt.Println(string(colorBlue), "DB Connection Sucses", string(colorReset))

	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	router := gin.Default()
	v1 := router.Group("v1")
	v1.GET("/", bookHandler.RootHandler)
	v1.GET("/books", bookHandler.GetBooks)
	v1.GET("/book/:id", bookHandler.GetBook)
	v1.PUT("/book/:id", bookHandler.UpdateBook)
	v1.DELETE("/book/:id", bookHandler.DeleteBook)
	v1.POST("/post", bookHandler.CreateBook)

	router.Run(":8000")
}

//Menit : 2.53
//main
//handler
//service
//db
//mysql
