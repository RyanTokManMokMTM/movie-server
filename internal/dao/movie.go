package dao

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/internal/models"
)

func (d *DAO) FindOneMovieDetail(ctx context.Context, movieID uint) (*models.MovieInfo, error) {
	movies := &models.MovieInfo{
		MovieId: movieID,
	}

	if err := movies.FindOneMovieWithGenres(ctx, d.engine); err != nil {
		return nil, err
	}
	return movies, nil
}

func (d *DAO) FindOneMovie(ctx context.Context, movieID uint) (*models.MovieInfo, error) {
	movies := &models.MovieInfo{
		MovieId: movieID,
	}

	if err := movies.FindOneMovie(ctx, d.engine); err != nil {
		return nil, err
	}
	return movies, nil
}
