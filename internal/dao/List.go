package dao

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/internal/models"
)

func (d *DAO) CreateNewList(ctx context.Context, ListTitle, ListIntro string, userID uint) (*models.List, error) {
	newList := &models.List{
		ListTitle: ListTitle,
		ListIntro: ListIntro,
		UserId:    userID,
	}

	if err := newList.CreateNewList(ctx, d.engine); err != nil {
		return nil, err
	}
	return newList, nil
}

func (d *DAO) UpdateList(ctx context.Context, list *models.List) error {
	return list.UpdateList(ctx, d.engine)
}

func (d *DAO) DeleteList(ctx context.Context, listID, userID uint) error {
	list := &models.List{
		UserId: userID,
		ListId: listID,
	}
	return list.DeleteList(ctx, d.engine)
}

func (d *DAO) FindOneList(ctx context.Context, listID uint) (*models.List, error) {
	list := &models.List{
		ListId: listID,
	}

	if err := list.FindOneList(ctx, d.engine); err != nil {
		return nil, err
	}
	return list, nil
}

func (d *DAO) FindOneUserList(ctx context.Context, listID, userID uint) (*models.List, error) {
	list := &models.List{
		ListId: listID,
		UserId: userID,
	}

	if err := list.FindOneList(ctx, d.engine); err != nil {
		return nil, err
	}
	return list, nil
}

func (d *DAO) FindUserLists(ctx context.Context, userID uint, limit, pageOffset int) ([]*models.List, int64, error) {
	list := &models.List{
		UserId: userID,
	}

	return list.FindAllList(ctx, d.engine, limit, pageOffset)
}

func (d *DAO) CountCollectedMovie(ctx context.Context, userID uint) (int64, error) {
	list := &models.ListMovie{}
	return list.CountMovieCollectedByUser(ctx, d.engine, userID)
}

////TO Check movie is in the list already
func (d *DAO) FindOneMovieFromList(ctx context.Context, movieID, listID, userID uint) (*models.MovieInfo, error) {
	list := &models.List{
		UserId: userID,
		ListId: listID,
	}
	MovieInfo := &models.MovieInfo{Id: movieID}

	if err := list.FindOneMovieFromList(ctx, d.engine, MovieInfo); err != nil {
		return nil, err
	}

	return MovieInfo, nil
}

func (d *DAO) InsertMovieToList(ctx context.Context, movieID, listID, userID uint) error {
	list := &models.List{
		UserId: userID,
		ListId: listID,
	}

	MovieInfo := models.MovieInfo{Id: movieID}

	if err := list.InsertMovieToList(ctx, d.engine, &MovieInfo); err != nil {
		return err
	}

	return nil
}

func (d *DAO) RemoveMovieFromList(ctx context.Context, movieID, listID, userID uint) error {
	list := &models.List{
		UserId: userID,
		ListId: listID,
	}

	MovieInfo := models.MovieInfo{Id: movieID}

	if err := list.RemoveMovieFromList(ctx, d.engine, &MovieInfo); err != nil {
		return err
	}

	return nil
}

func (d *DAO) RemoveMoviesFromList(ctx context.Context, movieIds []uint, listID, userID uint) error {
	list := &models.List{
		UserId: userID,
		ListId: listID,
	}

	return list.RemoveMoviesFromList(ctx, d.engine, movieIds)
}
