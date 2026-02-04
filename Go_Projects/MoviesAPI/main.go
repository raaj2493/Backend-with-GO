package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

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

var movies []Movies

func main() {
	rand.Seed(time.Now().UnixNano())

	r := gin.Default()

	r.GET("/movies", getMovies)
	r.GET("/movies/:id", getMovieByID)
	r.POST("/movies", createMovie)
	r.PUT("/movies/:id", updateMovie)
	r.DELETE("/movies/:id", deleteMovie)

	r.Run(":8080")
}

/* ================= GET ALL MOVIES ================= */

func getMovies(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"count":  len(movies),
		"movies": movies,
	})
}

/* ================= GET MOVIE BY ID ================= */

func getMovieByID(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid movie id",
		})
		return
	}

	for _, movie := range movies {
		if movie.ID == id {
			c.JSON(http.StatusOK, movie)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "movie not found",
	})
}

/* ================= CREATE MOVIE ================= */

func createMovie(c *gin.Context) {
	var input Movies

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	if input.Title == "" || input.Year == 0 || input.Director == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "title, year and director are required",
		})
		return
	}

	newMovie := Movies{
		ID:     rand.Intn(1000000),
		Title:  input.Title,
		Year:   input.Year,
		Rating: input.Rating,
		Director: &Director{
			Name: input.Director.Name,
		},
	}

	movies = append(movies, newMovie)

	c.JSON(http.StatusCreated, newMovie)
}

/* ================= UPDATE MOVIE ================= */

func updateMovie(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid movie id",
		})
		return
	}

	var input Movies
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	for i, movie := range movies {
		if movie.ID == id {
			movies[i].Title = input.Title
			movies[i].Year = input.Year
			movies[i].Rating = input.Rating
			movies[i].Director = input.Director

			c.JSON(http.StatusOK, movies[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "movie not found",
	})
}

/* ================= DELETE MOVIE ================= */

func deleteMovie(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid movie id",
		})
		return
	}

	for i, movie := range movies {
		if movie.ID == id {
			movies = append(movies[:i], movies[i+1:]...)

			c.JSON(http.StatusOK, gin.H{
				"message": "movie deleted successfully",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "movie not found",
	})
}
