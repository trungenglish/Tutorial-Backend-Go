package controller

import (
	"net/http"
	"strconv"
	"tutorial/model"
	"tutorial/service/db"
	"tutorial/service/logger"

	"github.com/gin-gonic/gin"
)

func createMovies(c *gin.Context) {
	var movie model.Movies
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.CreateMovie(&movie); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movie)
}

func getMoviesById(c *gin.Context) {
	ctx := c.Request.Context()

	idParam, _ := strconv.Atoi(c.Param("id"))

	logger.Info(ctx).
		Uint("movie_id", uint(idParam)).
		Msg("Getting movie by ID")

	movie, err := db.GetMovieById(ctx, uint(idParam))
	if err != nil {
		logger.Error(ctx).
			Err(err).
			Uint("movie_id", uint(idParam)).
			Msg("Movie not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	logger.Info(ctx).
		Uint("movie_id", movie.ID).
		Msg("Movie retrieved successfully")
	c.JSON(http.StatusOK, movie)
}

func searchMovies(c *gin.Context) {
	q := c.Query("q")
	year, _ := strconv.Atoi(c.Query("year"))

	movies, err := db.SearchMovies(q, year)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movies)
}

func getMoviesOffset(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

	movies, err := db.GetMoviesOffset(page, size)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movies)
}

func GetMoviesCursor(c *gin.Context) {
	cursor, _ := strconv.Atoi(c.DefaultQuery("cursor", "0"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
	movies, err := db.GetMoviesCursor(uint(cursor), size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movies)
}
