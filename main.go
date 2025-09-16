package main

import (
	"tutorial/config"
	"tutorial/controller"
	"tutorial/service/cache"
	"tutorial/service/db"

	"github.com/gin-gonic/gin"
)

func main() {
	//config
	cfg := config.InitConfig()

	//connect to database
	db.ConnectDB(cfg)

	// Kết nối Memcached
	cache.InitCache(cfg.MemcachedAddr)

	//if err := seed.SeedMovies(db.DB); err != nil {
	//	log.Fatalf("❌ Seed failed: %v", err)
	//}

	//route
	r := gin.Default()
	controller.SetupRouter(r)

	r.Run(":" + cfg.Port)
}
