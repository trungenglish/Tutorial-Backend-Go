package main

import (
	"tutorial/controller"
	"tutorial/config"
	"github.com/gin-gonic/gin"
)

func main() {
	//config
	cfg := config.InitConfig()

	//route
	r := gin.Default()
	controller.SetupRouter(r)

	r.Run("project run on port: " + cfg.Port)
}