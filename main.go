package main

import (
	"tutorial/config"
	"tutorial/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	//config
	cfg := config.InitConfig()

	//route
	r := gin.Default()
	controller.SetupRouter(r)

	r.Run(":" + cfg.Port)
}
