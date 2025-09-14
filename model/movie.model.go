package model

import "time"

type Movies struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `json:"title"`
	Genre     string    `json:"genre"`
	Year      int       `json:"year"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
}
