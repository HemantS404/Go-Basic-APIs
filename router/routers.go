package router

import (
	"example/mongo-go/helper"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	r := gin.Default()
	r.GET("/movies", helper.GetAllMovies)
	r.POST("/movies", helper.InsertOneMovie)
	r.PATCH("/movies/:id", helper.UpdateByID)
	r.DELETE("/movies/:id", helper.DeleteByID)
	r.DELETE("/movies", helper.DeleteAllMovie)
	return r
}
