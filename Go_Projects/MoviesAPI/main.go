package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

}


