package main

import (
	"fmt"
	"log"
	database "main/Database"
	model "main/Model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func loadDatabase() {
	database.Init()
	database.Database.AutoMigrate(&model.Book{})
}

func main() {
	loadEnv()
	loadDatabase()
	fmt.Println("hellow world !!!")
	fmt.Println("your application is working")
	router := gin.Default()
	router.GET("/books", getBooks)
	// router.GET("books/:id", bookById)
	router.POST("/books", createBook)
	// router.PATCH("/checkout", checkoutBook)
	// router.PATCH("/return", returnBook)
	router.Run("localhost:8080")
}

func getBooks(c *gin.Context) {
	books := []model.Book{}
	db := database.Database
	if err := db.Find(&books).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Books not found"})
	} else {
		c.IndentedJSON(http.StatusOK, books)
	}
}

// func bookById(c *gin.Context) {
// 	id := c.Param("id")
// 	book, err := getBookById(id)

// 	if err != nil {
// 		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
// 		return
// 	}
// 	c.IndentedJSON(http.StatusOK, book)
// }

// func getBookById(id string) (*book, error) {
// 	for i, book := range books {
// 		if book.ID == id {
// 			return &books[i], nil
// 		}
// 	}
// 	return nil, errors.New("book not found")
// }

func createBook(c *gin.Context) {
	newBook := model.Book{}
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	db := database.Database
	db.Create(&newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

// func checkoutBook(c *gin.Context) {
// 	id, ok := c.GetQuery("id")

// 	if !ok {
// 		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
// 		return
// 	}

// 	book, err := getBookById(id)

// 	if err != nil {
// 		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
// 		return
// 	}

// 	if book.Quantity <= 0 {
// 		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not available."})
// 		return
// 	}

// 	book.Quantity -= 1
// 	c.IndentedJSON(http.StatusOK, book)
// }

// func returnBook(c *gin.Context) {
// 	id, ok := c.GetQuery("id")

// 	if !ok {
// 		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
// 		return
// 	}

// 	book, err := getBookById(id)

// 	if err != nil {
// 		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
// 		return
// 	}

// 	book.Quantity += 1
// 	c.IndentedJSON(http.StatusOK, book)
// }

// func main() {
// fmt.Println("hellow world !!!")
// fmt.Println("your application is working")
// router := gin.Default()
// router.GET("/books", getBooks)
// router.GET("books/:id", bookById)
// router.POST("/books", createBook)
// router.PATCH("/checkout", checkoutBook)
// router.PATCH("/return", returnBook)
// router.Run("localhost:8080")
// }
