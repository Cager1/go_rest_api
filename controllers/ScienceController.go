package controllers

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"tim_go_api/database"
	"tim_go_api/models"
)

//TODO THIS FIRST !!!!!
// Implement so that model.Science model accepts slice of models.Book structs rather than only id's!
// This will serve if we want to change sciences on multiple books at once or when we add a new field we can add
// corresponding books to corresponding science

// GetSciences GET request for all sciences in database
// TODO Implement models.Book as collection in response
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

func getScience(id int64) (models.Science, error) {
	var science models.Science

	row := database.Db.QueryRow("SELECT * FROM science WHERE id = ?", id)
	if err := row.Scan(&science.ID, &science.Name); err != nil {
		if err == sql.ErrNoRows {
			return science, fmt.Errorf("science not found")
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

// postSciences POST request for posting a new science
func postSciences(name string) (models.Science, error) {

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	row, err := database.Db.Exec(
		"INSERT INTO science (name) VALUES (?)", name)
	if err != nil {
		return models.Science{}, err
	}
	lastScienceID, _ := row.LastInsertId()
	lastScience, _ := getScience(lastScienceID)

	return lastScience, nil
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
	lastScience, _ := getScience(lastScienceID)
	c.IndentedJSON(http.StatusCreated, lastScience)
}
