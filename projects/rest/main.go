// Testing API Endpoint with Gin

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Use Struct Tags without spaces!!!
// Indicate the name of the field when serialized into JSON
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Album 1", Artist: "Artist1", Price: 1.},
	{ID: "2", Title: "Album 2", Artist: "Artist2", Price: 2.},
	{ID: "3", Title: "Album 3", Artist: "Artist3", Price: 3.},
}

func appendAlbum(newAlbum album) {
	albums = append(albums, newAlbum)
}

// Route Handler, Gin Context contains informations about requests
func getAllAlbums(context *gin.Context) {
	// JSON is compact version of IndentedJSON, there is also AsciiJSON
	context.JSON(http.StatusOK, albums)
}

func addNewAlbum(context *gin.Context) {
	var newAlbum album

	// Bind the request JSON to variable newAlbum by using BindJSON
	if err := context.BindJSON(&newAlbum); err != nil {
		return
	}

	appendAlbum(newAlbum)
	context.JSON(http.StatusCreated, newAlbum)
}

func getAlbum(context *gin.Context) {
	// Get the parameters in the URL with the key specified on the router
	albumId := context.Param("id")

	for _, album := range albums {
		if album.ID == albumId {
			context.JSON(http.StatusOK, album)
			return
		}
	}
	// gin.H is basically map[string]interface{}, can be used to create JSON on the fly
	context.JSON(http.StatusNotFound, gin.H{"message": "Album Not Found"})
}

// The server
func main() {
	// Creating Gin router and engine with HTTP Server attached
	router := gin.Default()
	// Assigning the route with the handler function
	router.GET("/albums", getAllAlbums)
	router.POST("/albums", addNewAlbum)

	// URL Parameters handling
	router.GET("/albums/:id", getAlbum)

	router.Run("localhost:3000")
}
