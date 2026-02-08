package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func ParseBody(c *gin.Context, x interface{}) error {
	return c.ShouldBindJSON(x)
}