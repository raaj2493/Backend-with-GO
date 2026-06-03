package herosapi

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"database/sql"
)

func main (){
     r := gin.Default();

	 r.GET("/heros", getHeros);
	 r.GET("/heros/:id", getHerosbyId);
	 r.Run(":8080");
}

func getHeros (c *gin.Context){
	

}

func getHerosbyId (c *gin.Context){

}