package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Movie struct {
	ID string `json:"id"`;
	Title string `json:"title"`;
	Director string `json:"director"`
	Year     int    `json:"year"`;
}

var movie = []Movie{
	{
		ID:       "1",
		Title:    "Interstellar",
		Director: "Christopher Nolan",
		Year:     2014,
	},
	{
		ID:       "2",
		Title:    "Inception",
		Director: "Christopher Nolan",
		Year:     2010,
	},
}

func main (){

	router := gin.Default();
	router.Run(":8080");
	router.GET("/movies" , getMovies);
	router.GET("/movies/id", getMoviesbyId);
}

func getMovies(c *gin.Context){
     c.JSON(http.StatusOK , movie)
}

func getMoviesbyId(c *gin.Context){
	id := c.Param("id");

	for _ , movie := range movie{
		if movie.ID == id {
			c.JSON(http.StatusOK , movie )
			return 
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"error": "Movie not found",
	})
}


func createMovies(c *gin.Context){
	var newMovie Movie;
	err := c.BindJSON(&newMovie);
	if err != nil {
		c.JSON(http.StatusBadRequest , gin.H{
			"error": err.Error(),	
		})
	}
	movie = append(movie, newMovie)

	c.JSON(http.StatusCreated , gin.H{
		"message": "Movie created successfully",
		"movie":   "newMovie",
	})
}


func updateMovies(c *gin.Context){
	var updateMovie Movie;
	 err := c.BindJSON(&updateMovie);
	 id := c.Params("id");

	 if err != nil {
		c.JSON(http.StatusBadRequest , gin.H{
			"error": "Invalid JSON",
		})
	 }

	 for index , movie := range movie{
		if movie.ID == id {
              updateMovie.ID = id ;
			  movie[index] = updateMovie

			  c.JSON(http.StatusOK , gin.H{
					"message": "Movie updated successfully",
				"movie":   updateMovie,
			  })

		}
		c.JSON(http.StatusNotFound, gin.H{
		"error": "Movie not found",
	})

	 }
}


func deleteMovie(c *gin.Context) {

	id := c.Param("id")

	for index, movie := range movie {

		if movie.ID == id {

			movie = append(
				movie[:index],
				movie[index+1:]...,
			)

			c.JSON(http.StatusOK, gin.H{
				"message": "Movie deleted successfully",
			})

			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "Movie not found",
	})
}

