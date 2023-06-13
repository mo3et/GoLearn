package main

import(
    "net/http"

    "github.com/gin-gonic/gin"
)

/*
# Design API endpoints
You’ll build an API that provides access to a store selling vintage recordings on vinyl.
So you’ll need to provide endpoints through which a client can get and add albums for users.

When developing an API, you typically begin by designing the endpoints.
Your API’s users will have more success if the endpoints are easy to understand.

Here are the endpoints you’ll create in this tutorial.

/albums

	- GET – Get a list of all albums, returned as JSON.
	- POST – Add a new album from request data sent as JSON.

/albums/:id

	- GET – Get an album by its ID, returning the album data as JSON.

Next, you’ll create a folder for your code.
*/

// # Create the data

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// # Write a handler to return all items

// getAlbums reponds with the list of all albums as JSON.
func getAlbums(c *gin.Context){
    c.IndetedJSON(http.StatusOK,albums)
}

func main(){
    router:=gin.Default()
    router.GET("/albums",getAlbums)

    router.Run("localhost:8080")
}


// # Write a handler to return a specific item

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

    // Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
