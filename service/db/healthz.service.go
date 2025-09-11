package db

import (
	"tutorial/model"
	"time"
)

func Healthz() *model.Healthz{
	return &model.Healthz{
		Status: "ok",
		Version: "v1.0.0",
		Time: time.Now().UTC().Format(time.RFC3339),
	}
}