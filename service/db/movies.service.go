package db

import (
	"encoding/json"
	"fmt"
	"time"
	"tutorial/model"
	"tutorial/service/cache"
	"tutorial/service/logger"
	"tutorial/service/metrics"

	"github.com/bradfitz/gomemcache/memcache"
	"gorm.io/gorm"
)

func CreateMovie(movie *model.Movies) error {
	start := time.Now()
	result := DB.Create(movie)
	duration := time.Since(start).Seconds()

	metrics.DBQueryDuration.Observe(duration)

	if result.Error != nil {
		logger.Log.Error().
			Err(result.Error).
			Str("operation", "CreateMovie").
			Msg("❌ Failed to create movie")
		return result.Error
	}

	// Invalidate cache (xóa cache list movies hoặc liên quan)
	cache.Client.Delete(fmt.Sprintf("movie:%d", movie.ID))
	logger.Log.Info().
		Uint("movie_id", movie.ID).
		Msg("🗑️ Cache invalidated for movie")

	return nil
}

func GetMovieById(id uint) (*model.Movies, error) {
	cacheKey := fmt.Sprintf("movie:%d", id)

	// 1. Kiểm tra cache
	item, err := cache.Client.Get(cacheKey)

	if err == nil {
		if string(item.Value) == "null" {
			logger.Log.Info().
				Uint("movie_id", id).
				Msg("✅ Cache hit (negative)")
			metrics.CacheHitTotal.Inc()
			return nil, gorm.ErrRecordNotFound
		}

		logger.Log.Info().
			Uint("movie_id", id).
			Msg("✅ Cache hit")
		metrics.CacheHitTotal.Inc()

		var movie model.Movies
		json.Unmarshal(item.Value, &movie)
		return &movie, nil
	}
	if err == memcache.ErrCacheMiss {
		logger.Log.Warn().
			Uint("movie_id", id).
			Msg("❌ Cache miss")
		metrics.CacheMissTotal.Inc()
	} else {
		logger.Log.Error().
			Err(err).
			Str("cache_key", cacheKey).
			Msg("⚠️ Cache error")
	}

	// 2. Nếu cache miss → query DB
	start := time.Now()
	var movie model.Movies
	result := DB.First(&movie, id)
	duration := time.Since(start).Seconds()
	metrics.DBQueryDuration.Observe(duration)

	if result.Error != nil {
		// Negative cache (not found)
		cache.Client.Set(&memcache.Item{
			Key:        cacheKey,
			Value:      []byte("null"),
			Expiration: 30, // TTL 30s
		})
		logger.Log.Warn().
			Uint("movie_id", id).
			Msg("Movie not found, cached as null")
		return nil, result.Error
	}

	// 3. Lưu vào cache (TTL 5 phút)
	data, _ := json.Marshal(movie)
	cache.Client.Set(&memcache.Item{
		Key:        cacheKey,
		Value:      data,
		Expiration: 300, // TTL 5 phút
	})
	logger.Log.Info().
		Uint("movie_id", movie.ID).
		Msg("Movie cached for 5 minutes")

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

	start := time.Now()
	err := query.Find(&movies).Error
	duration := time.Since(start).Seconds()

	metrics.DBQueryDuration.Observe(duration)

	if err != nil {
		logger.Log.Error().
			Err(err).
			Str("query", q).
			Int("year", year).
			Msg("❌ Failed to search movies")
	} else {
		logger.Log.Info().
			Int("count", len(movies)).
			Str("query", q).
			Int("year", year).
			Float64("duration_sec", duration).
			Msg("✅ Movies search executed")
	}

	return movies, err
}

func GetMoviesOffset(page, pageSize int) ([]model.Movies, error) {
	var movies []model.Movies
	offset := (page - 1) * pageSize

	start := time.Now()
	result := DB.Limit(pageSize).Offset(offset).Find(&movies)
	duration := time.Since(start).Seconds()
	metrics.DBQueryDuration.Observe(duration)

	if result.Error != nil {
		logger.Log.Error().
			Err(result.Error).
			Int("page", page).
			Int("pageSize", pageSize).
			Msg("❌ Failed to get movies (offset)")
	} else {
		logger.Log.Info().
			Int("page", page).
			Int("pageSize", pageSize).
			Int("count", len(movies)).
			Float64("duration_sec", duration).
			Msg("✅ Movies retrieved with offset pagination")
	}

	return movies, result.Error
}

func GetMoviesCursor(cursorID uint, pageSize int) ([]model.Movies, error) {
	var movies []model.Movies

	query := DB.Limit(pageSize).Order("id ASC")

	if cursorID > 0 {
		query = query.Where("id > ?", cursorID)
	}

	start := time.Now()
	result := query.Find(&movies)
	duration := time.Since(start).Seconds()
	metrics.DBQueryDuration.Observe(duration)

	if result.Error != nil {
		logger.Log.Error().
			Err(result.Error).
			Uint("cursorID", cursorID).
			Int("pageSize", pageSize).
			Msg("❌ Failed to get movies (cursor)")
	} else {
		logger.Log.Info().
			Uint("cursorID", cursorID).
			Int("pageSize", pageSize).
			Int("count", len(movies)).
			Float64("duration_sec", duration).
			Msg("✅ Movies retrieved with cursor pagination")
	}

	return movies, result.Error
}
