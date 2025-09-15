package db

import (
	"tutorial/model"
)

func CreateMovie(movie *model.Movies) error {
	return DB.Create(movie).Error
}

func GetMovieById(id uint) (*model.Movies, error) {
	var movie model.Movies
	result := DB.First(&movie, id)
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
