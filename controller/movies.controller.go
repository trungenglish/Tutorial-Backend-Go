package controller

import (
	"net/http"
	"strconv"
	"tutorial/model"
	"tutorial/service/db"

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
	idParam, _ := strconv.Atoi(c.Param("id"))
	movie, err := db.GetMovieById(uint(idParam))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}
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
