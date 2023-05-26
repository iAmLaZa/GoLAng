package models

import (
	"github.com/iAmLaZa/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db_movie *gorm.DB

type Movie struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db_movie = config.GetDB()
	db_movie.AutoMigrate(&Movie{})
}

func (m *Movie) CreateMovie() *Movie {
	db_movie.NewRecord(m)
	db_movie.Create(&m)
	return m
}

func GetAllMovies() []Movie {
	var Movies []Movie
	db_movie.Find(&Movies)
	return Movies
}

func GetMovieById(Id int64) (*Movie, *gorm.DB) {
	var getMovie Movie
	db_movie := db_movie.Where("ID=?", Id).Find(&getMovie)
	return &getMovie, db_movie
}

func DeleteMovie(ID int64) Movie {
	var movie Movie
	db_movie.Where("ID=?", ID).Delete(movie)
	return movie
}
