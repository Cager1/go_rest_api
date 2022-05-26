package controllers

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"tim_go_api/database"
	"tim_go_api/models"
)

// GetBooks GET request for all books in database
func GetBooks(c *gin.Context) {
	var books []models.Book

	rows, err := database.Db.Query("SELECT * FROM book")
	if err != nil {
		_ = fmt.Errorf("%v", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(rows)
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author); err != nil {
			if err == sql.ErrNoRows {
				_ = fmt.Errorf("bookById %v: no such book", book)
			}
			_ = fmt.Errorf("bookByAuthor %v", err)
		}
		books = append(books, book)
	}
	if err := rows.Err(); err != nil {
		_ = fmt.Errorf("bookByAuthor %v", err)
	}

	c.IndentedJSON(http.StatusOK, books)
}

// PostBook POST request for posting a new book
func PostBook(c *gin.Context) {

	var newBook models.Book

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newBook); err != nil {
		fmt.Println(err)
	}

	row, err := database.Db.Exec(
		"INSERT INTO book (title, author) VALUES (?, ?)", newBook.Title, newBook.Author)
	if err != nil {
		_ = fmt.Errorf("addBook: %v", err)
	}

	lastBookId, _ := row.LastInsertId()
	lastBook, _ := getBook(lastBookId)

	sciences := newBook.Sciences
	fmt.Println(sciences)
	for _, science := range sciences {
		fmt.Println(science)
		row, err := database.Db.Exec(
			"INSERT INTO book_sciences (book_id, science_id) VALUES (?, ?)", lastBook.ID, science)
		if err != nil {
			_ = fmt.Errorf("addBookScience: %v", err)
		}
		fmt.Println(row)
	}
	c.IndentedJSON(http.StatusCreated, lastBook)
}

func getBook(id int64) (models.Book, error) {
	var book models.Book

	row := database.Db.QueryRow("SELECT * FROM book WHERE id = ?", id)
	if err := row.Scan(&book.ID, &book.Title, &book.Author); err != nil {
		if err == sql.ErrNoRows {
			return book, fmt.Errorf("book not found")
		}
	}
	fmt.Println("book: ", book)
	return book, nil
}

// GetBook GET request for one book by specific ID
func GetBook(c *gin.Context) {
	id := c.Param("id")

	var book models.Book

	row := database.Db.QueryRow("SELECT * FROM book WHERE id = ?", id)
	if err := row.Scan(&book.ID, &book.Title, &book.Author); err != nil {
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
			return
		}
	}

	println(row)
	c.IndentedJSON(http.StatusOK, book)
}
