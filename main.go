package main

import (
	"context"
	"tutorial/config"
	"tutorial/controller"
	"tutorial/service/cache"
	"tutorial/service/db"
	"tutorial/service/logger"
	"tutorial/service/metrics"
	"tutorial/service/tracing"

	"github.com/gin-gonic/gin"
)

func main() {
	//config
	cfg := config.InitConfig()

	// init logger
	logger.InitLogger()

	// init metrics
	metrics.InitMetrics()

	// init tracing
	cleanup := tracing.InitTracer(context.Background())
	defer cleanup()

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
