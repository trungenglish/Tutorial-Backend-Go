package main

import (
	"log"
	"tutorial/config"
	"tutorial/controller"
	"tutorial/service/db"
	"tutorial/service/db/seed"

	"github.com/gin-gonic/gin"
)

func main() {
	//config
	cfg := config.InitConfig()

	//connect to database
	db.ConnectDB(cfg)

	if err := seed.SeedMovies(db.DB); err != nil {
		log.Fatalf("❌ Seed failed: %v", err)
	}

	//route
	r := gin.Default()
	controller.SetupRouter(r)

	r.Run(":" + cfg.Port)
}
