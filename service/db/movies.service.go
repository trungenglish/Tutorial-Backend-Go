package db

import (
	"encoding/json"
	"fmt"
	"log"
	"tutorial/model"
	"tutorial/service/cache"

	"github.com/bradfitz/gomemcache/memcache"
	"gorm.io/gorm"
)

func CreateMovie(movie *model.Movies) error {
	result := DB.Create(movie)
	if result.Error != nil {
		return result.Error
	}

	// Invalidate cache (xóa cache list movies hoặc liên quan)
	cache.Client.Delete(fmt.Sprintf("movie:%d", movie.ID))
	log.Println("🗑️ Cache invalidated for movie", movie.ID)

	return nil
}

func GetMovieById(id uint) (*model.Movies, error) {
	cacheKey := fmt.Sprintf("movie:%d", id)

	// 1. Kiểm tra cache
	item, err := cache.Client.Get(cacheKey)

	if err == nil {
		if string(item.Value) == "null" {
			log.Println("✅ Cache hit (negative)")
			return nil, gorm.ErrRecordNotFound
		}

		log.Println("✅ Cache hit")
		var movie model.Movies
		json.Unmarshal(item.Value, &movie)
		return &movie, nil
	}
	if err == memcache.ErrCacheMiss {
		log.Println("❌ Cache miss")
	} else {
		log.Println("⚠️ Cache error:", err)
	}

	// 2. Nếu cache miss → query DB
	var movie model.Movies
	result := DB.First(&movie, id)

	if result.Error != nil {
		// Negative cache (not found)
		cache.Client.Set(&memcache.Item{
			Key:        cacheKey,
			Value:      []byte("null"),
			Expiration: 30, // TTL 30s
		})
		return nil, result.Error
	}

	// 3. Lưu vào cache (TTL 5 phút)
	data, _ := json.Marshal(movie)
	cache.Client.Set(&memcache.Item{
		Key:        cacheKey,
		Value:      data,
		Expiration: 300, // TTL 5 phút
	})

	return &movie, result.Error
}

func SearchMovies(q string, year int) ([]model.Movies, error) {
	var movies []model.Movies
	query := DB.Model(&model.Movies{})
	if q != "" {
		query = query.Where("title ILIKE ?", "%"+q+"%")
	}
	if year != 0 {
		query = query.Where("year = ?", year)
	}
	err := query.Find(&movies).Error
	return movies, err
}

func GetMoviesOffset(page, pageSize int) ([]model.Movies, error) {
	var movies []model.Movies
	offset := (page - 1) * pageSize
	result := DB.Limit(pageSize).Offset(offset).Find(&movies)
	return movies, result.Error
}

func GetMoviesCursor(cursorID uint, pageSize int) ([]model.Movies, error) {
	var movies []model.Movies
	query := DB.Limit(pageSize).Order("id ASC")
	if cursorID > 0 {
		query = query.Where("id > ?", cursorID)
	}
	result := query.Find(&movies)
	return movies, result.Error
}
