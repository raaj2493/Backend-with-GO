package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Movies struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	Year     int       `json:"year"`
	Rating   float64   `json:"rating"`
	Director *Director `json:"director"`
}

type Director struct {
	Name string `json:"name"`
}

var movies[] Movies

func main() {
    moviesapi := gin.Default()

    moviesapi.GET("/movies", getMovies)
	moviesapi.GET("/movies/:id" , getMoviesbyId)
	moviesapi.POST("/movies" , postMovies)
	moviesapi.PUT("/movies/:id" , updateMovies)
	moviesapi.DELETE("/movies/:id", deleteMovies)

}

func getMovies (c *gin.Context){
   c.JSON(http.StatusOK , gin.H{
	    "count" : len(movies),
		"movies" : movies,
   })
}

func getMoviesbyId (c *gin.Context){
     id := c.Param("id")

	 if id == "" {
		c.JSON(http.StatusBadRequest , gin.H{
			"error" : "Movie id is required",
		})
		return

	 }

	 for _, movie := range movies {
		 if movie.ID == id {
			c.JSON(http.StatusOK, movie)
		 }
	 }

	 c.JSON(http.StatusNotFound, gin.H{
		"error": "movie not found",
	})
}

func postMovies (c *gin.Context) {
     var input movie 

	 if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest , gin.H{
			"message" : "invalid request body",
		})
		return
	 }

	 if input.Title == "" || input.Isbn == "" || input.Director == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "title, isbn and director are required",
		})
		return
	}

	newmovie := Movies{
		ID:    strconv.Itoa(rand.Intn(1000000)),
		Title: input.Title,
		Year: input.Year,
		Rating: input.Rating,
		Director: &Director{
			Name: input.Director.name,
		},
	}

	movies = append(movies, newmovie)
	c.JSON(http.StatusCreated, newmovie)


}

func updateMovies (c *gin.Context)  {
	
}

func deleteMovies (c *gin.Context)  {
	id = c.Param("id")

	for i , movies := range movies {
         if movies.ID == id {
			movies = append(movies[:i] , movies[i+1]...)

			c.JSON(http.StatusOK , gin.H{
				"message": "movie deleted successfully",
			})
			return
		 }

		 c.JSON(http.StatusNotFound, gin.H{
		"error": "movie not found",
	})
	}
}
