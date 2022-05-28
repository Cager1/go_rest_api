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
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	//Close all rows at the end of a function
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, err)
			return
		}
	}(rows)

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {

		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author); err != nil {
			if err == sql.ErrNoRows {
				c.IndentedJSON(http.StatusNotFound, err)
			}
			c.IndentedJSON(http.StatusNotFound, err)
		}

		// Get all science_ids for this specific book from book_sciences table
		rows, err := database.Db.Query(
			"SELECT science_id FROM book_sciences WHERE book_id = ?", book.ID)
		if err != nil {
			_ = fmt.Errorf("%v", err)
		}

		// Assign returned sciences to book
		for rows.Next() {
			var scienceID int
			var science models.Science

			if err := rows.Scan(&scienceID); err != nil {
				if err == sql.ErrNoRows {
					fmt.Println(err)
				}
				fmt.Println(err)
			}

			science, err = getScience(int64(scienceID))
			if err != nil {
				fmt.Println(err)
				return
			}
			book.Sciences = append(book.Sciences, science)
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
	// newBook.
	if err := c.BindJSON(&newBook); err != nil {
		fmt.Println(err)
	}

	row, err := database.Db.Exec(
		"INSERT INTO book (title, author) VALUES (?, ?)", newBook.Title, newBook.Author)
	if err != nil {

		er := fmt.Errorf("PostBook: %w", err)
		fmt.Println(er)
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	lastBookId, _ := row.LastInsertId()
	lastBook, _ := getBook(lastBookId)

	sciences := newBook.Sciences
	// Iterate thought all given sciences
	for _, science := range sciences {
		newScience := science
		fmt.Println(newScience)
		ns := &newScience
		// If science.Name is not empty try to post new Science
		if science.Name != "" {
			// If science.Name is contained in Science database table don't add to table

			row := database.Db.QueryRow("SELECT * FROM science WHERE name = ?", science.Name)
			if err := row.Scan(&science.ID, &science.Name); err != nil {
				if err == sql.ErrNoRows {
					fmt.Println("No such science")
					*ns, err = postSciences(science.Name)
					if err != nil {
						fmt.Println(err)
						continue
					}
				}
			} else {
				*ns, err = getScience(int64(science.ID))
			}
		} else {
			*ns, err = getScience(int64(science.ID))
		}

		fmt.Println(*ns)
		// When above is done, insert existing/newly created science as foreign constraint in book_sciences table
		_, err := database.Db.Exec(
			"INSERT INTO book_sciences (book_id, science_id) VALUES (?, ?)", lastBook.ID, ns.ID)
		if err != nil {
			fmt.Printf("BookController PostBook:addBookScience: %v\n", err)
			continue
		}

		lastBook.Sciences = append(lastBook.Sciences, *ns)
	}
	c.IndentedJSON(http.StatusCreated, lastBook)
}

// GetBook GET request for one book by specific ID
// TODO Get models.Science for the book
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

// TODO Get models.Science as well
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

//TODO Implement option for deleting models.Book, also its connection with models.Science
