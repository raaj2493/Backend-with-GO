package main

import (
	"todo_api/internal/config"
	"todo_api/internal/database"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	var cfg *config.Config
	var err error
	cfg, err = config.Load()
	if err != nil {
		panic(err)
	}

	var pool *pgxpool.Pool
	pool, err = database.Connect(cfg.DATABASEURL)
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	var router *gin.Engine = gin.Default()
	router.SetTrustedProxies(nil)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "TODO API is Running",
			"Status":  "Success",
		})
	})
	router.Run(":3000")
}
