package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

var books = []book{
	{ID: "1", Title: "A Day", Author: "Vasant Korada", Price: 98.87},
	{ID: "2", Title: "Book 2", Author: "Savi Sharma", Price: 198.00},
	{ID: "3", Title: "Book 3", Author: "Chetan Bhagat", Price: 200},
}

func getAllBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func postBook(c *gin.Context) {
	var newBook book

	c.BindJSON(&newBook)

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func getABook(c *gin.Context) {
	var id = c.Param("id")

	for _, book := range books {
		if book.ID == id {
			c.IndentedJSON(http.StatusOK, book)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

func main() {
	router := gin.Default()

	router.GET("/books", getAllBooks)
	router.POST("/books", postBook)
	router.GET("/books/:id", getABook)

	router.Run()
}
