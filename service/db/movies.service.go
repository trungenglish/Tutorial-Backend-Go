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
