package controller

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {

	//route healthz
	r.GET("/healthz", Healthz)

	//route movies
	r.POST("/movies", createMovies)
	r.GET("/movies/:id", getMoviesById)
	r.GET("movies/search", searchMovies)
}
