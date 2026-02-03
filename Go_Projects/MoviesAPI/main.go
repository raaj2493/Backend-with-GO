package moviesapi

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
    router := gin.Default()

	
}
