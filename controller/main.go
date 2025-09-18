package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"tutorial/service/metrics"
)

func SetupRouter(r *gin.Engine) {
	//middleware
	r.Use(metrics.PrometheusMiddleware())

	//route healthz
	r.GET("/healthz", Healthz)

	//route movies
	r.POST("/movies", createMovies)
	r.GET("/movies/:id", getMoviesById)
	r.GET("/movies/search", searchMovies)
	//r.GET("/movies", getMoviesOffset)
	r.GET("/movies", GetMoviesCursor)

	//metrics
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

}
