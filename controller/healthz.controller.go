package controller

import (
	"net/http"
	"tutorial/service/db"

	"github.com/gin-gonic/gin"
)

func Healthz(c *gin.Context) {
	result := db.Healthz()
	c.JSON(http.StatusOK, result)
}
