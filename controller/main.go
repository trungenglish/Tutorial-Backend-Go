package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

	"tutorial/service/metrics"
)

func SetupRouter(r *gin.Engine) {
	//middleware
	r.Use(metrics.PrometheusMiddleware())
	r.Use(otelgin.Middleware("tutorial-service"))
	//r.Use(logger.TraceLoggingMiddleware())

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
