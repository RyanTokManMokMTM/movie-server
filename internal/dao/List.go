package dao

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/internal/models"
)

func (d *DAO) CreateNewList(ctx context.Context, ListTitle string) (*models.List, error) {
	newList := &models.List{
		ListTitle: ListTitle,
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

func (d *DAO) GetOneList(ctx context.Context, listID uint) (*models.List, error) {
	list := &models.List{
		ListId: listID,
	}

	if err := list.FindOneList(ctx, d.engine); err != nil {
		return nil, err
	}

	return list, nil
}

func (d *DAO) GetUserLists(ctx context.Context, userID uint) ([]*models.List, error) {
	list := &models.List{
		UserId: userID,
	}

	resp, err := list.FindAllList(ctx, d.engine)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *DAO) FindMovieFromList(ctx context.Context, movieID, listID, userID uint) ([]*models.List, error) {
	list := &models.List{
		UserId: userID,
	}

	resp, err := list.FindAllList(ctx, d.engine)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
