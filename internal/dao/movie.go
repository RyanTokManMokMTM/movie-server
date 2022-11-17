package dao

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/internal/models"
)

func (d *DAO) FindOneMovieDetailWithUserData(ctx context.Context, movieID, userID uint) (*models.MovieInfo, error) {
	movies := &models.MovieInfo{
		Id: movieID,
	}

	err := movies.FindOneMovieDetail(ctx, d.engine, userID)
	if err != nil {
		return nil, err
	}
	return movies, nil
}

func (d *DAO) FindOneMovieGenres(ctx context.Context, movieID uint) (*models.MovieInfo, error) {
	movies := &models.MovieInfo{
		Id: movieID,
	}

	err := movies.FindOneMovieGenres(ctx, d.engine)
	if err != nil {
		return nil, err
	}
	return movies, nil
}

func (d *DAO) FindOneMovie(ctx context.Context, movieID uint) (*models.MovieInfo, error) {
	movies := &models.MovieInfo{
		Id: movieID,
	}

	if err := movies.FindOneMovie(ctx, d.engine); err != nil {
		return nil, err
	}
	return movies, nil
}
