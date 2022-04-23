package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// album represents data about a record album.
type Attraction struct {
	Id         string `db:"id" json:"id"`
	Name       string `db:"name" json:"name"`
	Detail     string `db:"detail" json:"detail"`
	Coverimage string `db:"coverimage" json:"coverimage"`
}

// albums slice to seed record album data.
// var albums = []album{
// 	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
// 	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
// 	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// }

var db *sql.DB

func main() {
	var err error
	//you can change your database: username, password, localhost, dbname
	// db, err = sql.Open("mysql", "username:password@tcp(localhost:XXXX)/dbname")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/attractions", getAttractions)

	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAttractions(c *gin.Context) {
	// var (
	// 	id         int
	// 	name       string
	// 	detail     string
	// 	coverimage string
	// )
	var attractions []Attraction

	rows, err := db.Query("select id, name, detail, coverimage from attractions")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var a Attraction
		err := rows.Scan(&a.Id, &a.Name, &a.Detail, &a.Coverimage)
		if err != nil {
			log.Fatal(err)
		}
		attractions = append(attractions, a)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, attractions)
}
