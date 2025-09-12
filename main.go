package main

import (
	"tutorial/config"
	"tutorial/controller"
	"tutorial/service/db"

	"github.com/gin-gonic/gin"
)

func main() {
	//config
	cfg := config.InitConfig()

	//connect to database
	db.ConnectDB(cfg)

	//route
	r := gin.Default()
	controller.SetupRouter(r)

	r.Run(":" + cfg.Port)
}
