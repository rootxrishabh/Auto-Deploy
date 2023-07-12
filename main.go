// This is the main.go file
package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// This is the Book struct
type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

// This is a slice of type Book which is a struct
var books = []Book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

// This is the main function with all the routes
func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.POST("/books", createBooks)
	router.PATCH("/checkout", checkoutBook)
	router.PATCH("/return", returnBook)
	router.Run("0.0.0.0:8080")
}

// The functions below simply retries all the books
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// the 2 functions below helps us retrive a book by its id number
func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func getBookById(id string) (*Book, error) {
	for i, book_name := range books {
		if book_name.ID == id {
			return &books[i], nil
		}
	}
	x := errors.New("Book not found")
	return nil, x
}

// the function below helps us create a new entry ok a book.
func createBooks(c *gin.Context) {
	var newBook Book
	newData := c.BindJSON(&newBook)
	checkForErrors(newData)
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

// This function lends the book on request, manages the quantity
func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if ok == false {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "missing query parameter"})
	}

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ID NOT FOUND"})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not available"})
	} else {
		book.Quantity -= 1
		c.IndentedJSON(http.StatusOK, book)
	}

}

// This function takes the book on request, manages the quantity
func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if ok == false {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "missing query parameter"})
	}

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ID NOT FOUND"})
		return
	}
		book.Quantity += 1
		c.IndentedJSON(http.StatusOK, book)

}

// This is a standard error handling function
func checkForErrors(err error) {
	if err != nil {
		return
	}
}
