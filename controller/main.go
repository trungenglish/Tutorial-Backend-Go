package controller

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {

	//route healthz
	r.GET("/healthz", Healthz)

}
