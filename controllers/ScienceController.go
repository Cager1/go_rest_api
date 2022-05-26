package controllers

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"tim_go_api/database"
	"tim_go_api/models"
)

// GetSciences GET request for all sciences in database
func GetSciences(c *gin.Context) {
	var sciences []models.Science

	rows, err := database.Db.Query("SELECT * FROM science")
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
		var science models.Science
		if err := rows.Scan(&science.ID, &science.Name); err != nil {
			if err == sql.ErrNoRows {
				_ = fmt.Errorf("bookById %v: no such book", science)
			}
			_ = fmt.Errorf("bookByAuthor %v", err)
		}
		sciences = append(sciences, science)
	}
	if err := rows.Err(); err != nil {
		_ = fmt.Errorf("bookByAuthor %v", err)
	}

	c.IndentedJSON(http.StatusOK, sciences)
}

// PostSciences POST request for posting a new science
func PostSciences(c *gin.Context) {

	var newScience models.Science

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newScience); err != nil {
		fmt.Println(err)
	}

	row, err := database.Db.Exec(
		"INSERT INTO science (name) VALUES (?)", newScience.Name)
	if err != nil {
		_ = fmt.Errorf("addScience: %v", err)
	}

	lastScienceID, _ := row.LastInsertId()
	lastScience, _ := getBook(lastScienceID)
	c.IndentedJSON(http.StatusCreated, lastScience)
}

func getScience(id int64) (models.Science, error) {
	var science models.Science

	row := database.Db.QueryRow("SELECT * FROM book WHERE id = ?", id)
	if err := row.Scan(&science.ID, &science.Name); err != nil {
		if err == sql.ErrNoRows {
			return science, fmt.Errorf("book not found")
		}
	}
	fmt.Println("book: ", science)
	return science, nil
}

// GetScience GET request for one science by specific ID
func GetScience(c *gin.Context) {
	id := c.Param("id")

	var science models.Science

	row := database.Db.QueryRow("SELECT * FROM science WHERE id = ?", id)
	if err := row.Scan(&science.ID, &science.Name); err != nil {
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Science not found"})
			return
		}
	}
	c.IndentedJSON(http.StatusOK, science)
}
