package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/ping/:id", Pong)
		v1.GET("/ping-query/", PongQuery)
	}
	return r
}

func Pong(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong....ping",
		"metadata": gin.H{
			"id": id,
		},
	})
}
func PongQuery(c *gin.Context) {
	name := c.Query("name")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong query....ping",
		"metadata": gin.H{
			"name": name,
		},
	})
}
