package test

import (
	"fmt"
	"net/http/httptest"
	"os"
	"testing"
	"tutorial/config"
	"tutorial/controller"
	"tutorial/service/cache"
	"tutorial/service/db"

	"github.com/gavv/httpexpect/v2"
)

func TestMovieAPI(t *testing.T) {
	os.Chdir("..")
	cfg := config.InitConfig()
	db.ConnectDB(cfg)

	//logger.InitLogger()
	//
	//metrics.InitMetrics()

	cache.InitCache(cfg.MemcachedAddr)

	handler := controller.SetupRouter()

	server := httptest.NewServer(handler)
	defer server.Close()

	e := httpexpect.Default(t, server.URL)

	// POST /movies
	movie := map[string]interface{}{
		"title": "Inception",
		"year":  2010,
	}
	res := e.POST("/movies").
		WithJSON(movie).
		Expect().
		Status(201).
		JSON().Object()
	res.HasValue("title", "Inception")
	id := int(res.Value("id").Number().Raw())

	// GET /movies/{id}
	e.GET("/movies/"+fmt.Sprint(id)).
		Expect().
		Status(200).
		JSON().Object().HasValue("id", float64(id))

	// GET /movies/search?q=Inception
	e.GET("/movies/search").
		WithQuery("q", "Inception").
		Expect().
		Status(200).
		JSON().Array().Length().Ge(1)
}
